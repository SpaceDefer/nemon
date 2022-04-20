package nemon

import (
	"context"
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/gob"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	pb "nemon/protos"

	"github.com/1Password/srp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

// workerServer is  wrapper for pb.UnimplementedWorkerServer
type workerServer struct {
	pb.UnimplementedWorkerServer
}

// Worker struct is used by the
// Coordinator to connect to a worker
type Worker struct {
	ip         string           // ip address of the worker connected
	client     pb.WorkerClient  // client API exposed for the Coordinator
	connection *grpc.ClientConn // connection to the worker
	username   string           // username of the Worker
	os         string           // os of the Worker
	hostname   string           // hostname of the Worker
}

func (ws *workerServer) GetSaltAndSRP(_ context.Context, _ *pb.GetSaltAndSRPRequest) (*pb.GetSaltAndSRPResponse, error) {
	fmt.Printf("received getsaltandsrp")
	infoFile, err := os.Open(systemInfo.ConfigDir + "/enrollment_info.gob")

	if err != nil {
		return nil, status.Error(codes.Internal, "couldn't open info file")
	}

	decoder := gob.NewDecoder(infoFile)

	enrollmentInfo := EnrollmentInfo{}

	err = decoder.Decode(&enrollmentInfo)
	if err != nil {
		return nil, err
	}
	fmt.Println(enrollmentInfo)

	verifier := new(big.Int)
	verifier.SetBytes(enrollmentInfo.Verifier)

	srpServerInfo = SRPServerInfo{
		Verifier: verifier,
		Group:    int(enrollmentInfo.SRPGroup),
	}

	return &pb.GetSaltAndSRPResponse{
		Salt:     enrollmentInfo.Salt,
		SRPGroup: enrollmentInfo.SRPGroup,
	}, nil
}

func (ws *workerServer) ExchangeEphemeralPublic(_ context.Context, req *pb.ExchangeEphemeralPublicRequest) (*pb.ExchangeEphemeralPublicResponse, error) {
	fmt.Printf("exchange ephemeral public")
	ABytes := req.A
	A := new(big.Int)
	A.SetBytes(ABytes)
	group, verifier := srpServerInfo.Group, srpServerInfo.Verifier
	if verifier == nil {
		return nil, status.Error(codes.Unauthenticated, "verifier doesn't exist on the worker")
	}
	server := srp.NewSRPServer(srp.KnownGroups[group], verifier, nil)
	if server == nil {
		return nil, status.Error(codes.Internal, "couldn't set up server")
	}

	if err := server.SetOthersPublic(A); err != nil {
		return nil, status.Error(codes.Unauthenticated, "malicious A")
	}

	B := server.EphemeralPublic()
	if B == nil {
		return nil, status.Error(codes.Internal, "server couldn't make B")
	}

	serverKey, err := server.Key()

	if err != nil || serverKey == nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ExchangeEphemeralPublicResponse{
		B: B.Bytes(),
	}, nil
}

func (ws *workerServer) IsEnrolled(_ context.Context, req *pb.IsEnrolledRequest) (*pb.IsEnrolledResponse, error) {
	if req.Key != systemInfo.nemonKey {
		return nil, fmt.Errorf("keys not the same, refusing connection\n")
	}

	// check if persistently stored enrollment info
	infoFilePath := fmt.Sprintf(systemInfo.ConfigDir + "/enrollment_info.gob")
	_, err := os.Stat(infoFilePath)
	if err != nil {
		return &pb.IsEnrolledResponse{
			Enrolled: false,
		}, nil
	}
	return &pb.IsEnrolledResponse{
		Enrolled: true,
	}, nil
}

func (ws *workerServer) SaveEnrollmentInfo(_ context.Context, req *pb.SaveEnrollmentInfoRequest) (*pb.SaveEnrollmentInfoResponse, error) {
	// persist salt, verifier and SRPGroup, saving the verifier securely
	fmt.Printf("saving enrollment info\n")
	infoFileLocation := fmt.Sprintf(systemInfo.ConfigDir + "/enrollment_info.gob")
	infoFile, err := os.Create(infoFileLocation)
	defer infoFile.Close()
	if err != nil {
		return nil, status.Error(codes.Internal, "couldn't create gob file")
	}
	fmt.Printf("file created\n")
	enc := gob.NewEncoder(infoFile)
	err = enc.Encode(EnrollmentInfo{req.SRPGroup, req.Salt, req.Verifier})
	if err != nil {
		return nil, status.Error(codes.Internal, "couldn't encode")
	}

	fmt.Printf("saved successfully\n")
	return &pb.SaveEnrollmentInfoResponse{}, nil
}

// GetSysInfo handles the handshake and connection establishment and sends the Worker's SystemInfo if successful
func (ws *workerServer) GetSysInfo(_ context.Context, req *pb.GetSysInfoRequest) (*pb.GetSysInfoResponse, error) {
	if req.Key != systemInfo.nemonKey {
		return nil, fmt.Errorf("keys not the same, refusing connection\n")
	}

	var publicKey rsa.PublicKey

	publicKeyN := new(big.Int)
	var val string
	val = req.PublicKeyN
	publicKeyN.SetString(val, 10)
	publicKey.N = publicKeyN
	publicKey.E, _ = strconv.Atoi(req.PublicKeyE)
	AESKey := make([]byte, 32)

	_, err := rand.Read(AESKey)

	if err != nil {
		return nil, err
	}

	AESCipher, err := aes.NewCipher(AESKey)

	if err != nil {
		return nil, err
	}

	systemInfo.AESCipher = AESCipher
	systemInfo.AESKey = AESKey

	hash := sha512.New()
	encAESKey, err := rsa.EncryptOAEP(hash, rand.Reader, &publicKey, AESKey, nil)
	if err != nil {
		return nil, err
	}

	return &pb.GetSysInfoResponse{
		WorkerSysInfo: &pb.GetSysInfoResponse_SysInfo{
			Username: encrypt([]byte(systemInfo.username)),
			Os:       encrypt([]byte(systemInfo.OS)),
			Hostname: encrypt([]byte(systemInfo.hostname)),
		},
		AESKey: encAESKey,
	}, nil
}

// GetApps implements GetApps RPC from the generated ProtoBuf file
func (ws *workerServer) GetApps(_ context.Context, _ *pb.GetAppsRequest) (*pb.GetAppsResponse, error) {
	if systemInfo.AESKey == nil {
		return nil, status.Error(codes.Unauthenticated, "haven't authenticated yet, please authenticate")
	}
	fmt.Printf("got a GetApps gRPC\n")
	var err error
	var out []byte
	var applications []*pb.GetAppsResponse_ApplicationInfo

	switch systemInfo.OS {
	case "darwin":
		out, err = exec.Command("find", "/Applications", "-maxdepth", "3", "-iname", "*.app").Output()
		checkError(err)
		r := strings.Split(string(out), "\n")
		for _, str := range r {
			if len(str) > 0 {
				toAppend := strings.Split(str, "/")
				applications = append(applications, &pb.GetAppsResponse_ApplicationInfo{
					Name:     encrypt([]byte(toAppend[len(toAppend)-1])),
					Location: encrypt([]byte(str)),
				})
			}
		}
	case "windows":
		pwd, err := os.Getwd()
		checkError(err)
		out, err := exec.Command("python", pwd+"\\nemon\\scripts\\getAppListWindows.py").Output()
		checkError(err)
		list := string(out)
		res := strings.Split(list, "\n")
		for i := 0; i < len(res); i++ {
			str := strings.TrimSpace(res[i])
			if len(str) > 3 {
				applications = append(applications, &pb.GetAppsResponse_ApplicationInfo{
					Name:     encrypt([]byte(str)),
					Location: encrypt([]byte("/")),
				})
			}
		}
	case "linux":
		out, err = exec.Command("apt", "list", "--installed").Output()
	default:
		return nil, fmt.Errorf("unrecognised os %v", systemInfo.OS)
	}
	response := &pb.GetAppsResponse{
		Applications: applications,
		Username:     encrypt([]byte(systemInfo.username)),
	}
	return response, nil
}

// DeleteApp handles the deletion of an application on the Worker
func (ws *workerServer) DeleteApp(_ context.Context, req *pb.DeleteAppsRequest) (*pb.DeleteAppsResponse, error) {
	if systemInfo.AESKey == nil {
		return nil, status.Error(codes.Unauthenticated, "haven't authenticated yet, please authenticate")
	}
	location := decrypt(req.Location)
	fmt.Printf("%v, %v, %v\n", req, string(location), string(decrypt(req.Name)))
	switch systemInfo.OS {
	case "darwin":
		out, err := exec.Command("sudo", "rm", "-rf", string(location)).Output()
		checkError(err)
		fmt.Printf("%v\n", out)
	case "windows":
		//out, err := exec.Command("powershell", "-noprofile", "Get-WmiObject").Output()
	case "linux":
		out, err := exec.Command("apt", "remove", string(location)).Output()
		checkError(err)
		fmt.Printf("%v\n", out)
	default:
		return nil, fmt.Errorf("unrecognised os %v", systemInfo.OS)
	}
	return &pb.DeleteAppsResponse{Ok: true}, nil
}

// StartWorker handles starting up the worker on the machine
func StartWorker() {
	InitSystemInfo()
	var workerAddr, ip string
	if systemInfo.Dev {
		workerAddr = "localhost" + port
	} else {
		ip = GetLocalIP()
		workerAddr = ip + port
	}

	fmt.Printf("my ip on the network: %v\nhostname: %v\nusername: %v\n",
		ip,
		systemInfo.hostname,
		systemInfo.username,
	)
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)
	go func() {
		<-sigCh
		fmt.Printf("\nworker exiting gracefully...\n")
		// worker cleanup if needed
		os.Exit(1)
	}()
	conn, err := net.Listen("tcp", workerAddr)

	checkError(err)

	grpcServer := grpc.NewServer()

	server := workerServer{}

	pb.RegisterWorkerServer(grpcServer, &server)
	fmt.Printf("starting gRPC server at port %v...\n", port)
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
