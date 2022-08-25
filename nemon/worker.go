package nemon

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	pb "nemon/protos"

	"github.com/martinlindhe/notify"

	"golang.org/x/crypto/chacha20poly1305"

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
	status     Status           // status of the Worker
}

func (ws *workerServer) GetSaltAndSRP(_ context.Context, _ *pb.GetSaltAndSRPRequest) (*pb.GetSaltAndSRPResponse, error) {
	Debug(dInfo, "received getsaltandsrp\n")
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
	Debug(dInfo, "enrollment info: %v\n", enrollmentInfo)

	verifier := new(big.Int)
	verifier.SetBytes(enrollmentInfo.Verifier)

	srpServerInfo = SRPServerInfo{
		Verifier: verifier,
		Group:    int(enrollmentInfo.SRPGroup),
		Salt:     enrollmentInfo.Salt,
	}

	return &pb.GetSaltAndSRPResponse{
		Salt:     enrollmentInfo.Salt,
		SRPGroup: enrollmentInfo.SRPGroup,
	}, nil
}

func (ws *workerServer) VerifyClientProof(_ context.Context, req *pb.VerifyClientProofRequest) (*pb.VerifyClientProofResponse, error) {
	Debug(dInfo, "received a request for verification\n")
	server, serverKey := srpServerInfo.Server, srpServerInfo.ServerKey

	if server == nil || serverKey == nil {
		return nil, status.Error(codes.Unauthenticated, "srpServer doesn't exist")
	}

	if !server.GoodClientProof(req.ClientProof) {
		return nil, status.Error(codes.InvalidArgument, "bad proof")
	}

	serverCryptor, _ := chacha20poly1305.NewX(serverKey)

	systemInfo.Cryptor = serverCryptor

	return &pb.VerifyClientProofResponse{}, nil
}

func (ws *workerServer) ExchangeEphemeralPublic(_ context.Context, req *pb.ExchangeEphemeralPublicRequest) (*pb.ExchangeEphemeralPublicResponse, error) {
	Debug(dInfo, "exchange ephemeral public\n")
	ABytes := req.A
	A := new(big.Int)
	A.SetBytes(ABytes)
	salt, group, verifier := srpServerInfo.Salt, srpServerInfo.Group, srpServerInfo.Verifier
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

	srpServerInfo.Server, srpServerInfo.ServerKey = server, serverKey

	serverProof, err := server.M(salt, username)

	if err != nil {
		return nil, status.Error(codes.Internal, "couldn't make serverProof")
	}

	Debug(dInfo, "%v\n%v\n", salt, serverProof)

	return &pb.ExchangeEphemeralPublicResponse{
		B:           B.Bytes(),
		ServerProof: serverProof,
	}, nil
}

func (ws *workerServer) IsEnrolled(_ context.Context, req *pb.IsEnrolledRequest) (*pb.IsEnrolledResponse, error) {
	if req.Key != systemInfo.nemonKey {
		return nil, status.Error(codes.PermissionDenied, "keys not the same, refusing connection\n")
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
	Debug(dInfo, "saving enrollment info\n")
	infoFileLocation := fmt.Sprintf(systemInfo.ConfigDir + "/enrollment_info.gob")
	infoFile, err := os.Create(infoFileLocation)
	defer infoFile.Close()
	if err != nil {
		return nil, status.Error(codes.Internal, "couldn't create gob file")
	}
	Debug(dInfo, "file created\n")
	enc := gob.NewEncoder(infoFile)
	err = enc.Encode(EnrollmentInfo{req.SRPGroup, req.Salt, req.Verifier})
	if err != nil {
		return nil, status.Error(codes.Internal, "couldn't encode")
	}

	Debug(dInfo, "saved successfully\n")
	return &pb.SaveEnrollmentInfoResponse{}, nil
}

// GetSysInfo handles the handshake and connection establishment and sends the Worker's SystemInfo if successful
func (ws *workerServer) GetSysInfo(_ context.Context, _ *pb.GetSysInfoRequest) (*pb.GetSysInfoResponse, error) {
	return &pb.GetSysInfoResponse{
		WorkerSysInfo: &pb.GetSysInfoResponse_SysInfo{
			Username: encrypt([]byte(systemInfo.username)),
			Os:       encrypt([]byte(systemInfo.OS)),
			Hostname: encrypt([]byte(systemInfo.hostname)),
		},
	}, nil
}

// GetApps implements GetApps RPC from the generated ProtoBuf file
func (ws *workerServer) GetApps(_ context.Context, _ *pb.GetAppsRequest) (*pb.GetAppsResponse, error) {
	if systemInfo.Cryptor == nil {
		return nil, status.Error(codes.Unauthenticated, "haven't authenticated yet, please authenticate")
	}
	Debug(dInfo, "got a GetApps gRPC\n")
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
				splittedName := toAppend[len(toAppend)-1]
				removedExtensionName := strings.Split(splittedName, ".")
				name := removedExtensionName[len(removedExtensionName)-2]
				applications = append(applications, &pb.GetAppsResponse_ApplicationInfo{
					Name:     encrypt([]byte(name)),
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

func (ws *workerServer) NotifyWorker(_ context.Context, req *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	if systemInfo.Cryptor == nil {
		return nil, status.Error(codes.Unauthenticated, "haven't auth yet, please auth")
	}
	Debug(dInfo, "go a notif %v", req)

	notify.Notify("nemon", "Please come to the desk", "", "")
	return &pb.NotifyResponse{}, nil
}

// DeleteApp handles the deletion of an application on the Worker
func (ws *workerServer) DeleteApp(_ context.Context, req *pb.DeleteAppsRequest) (*pb.DeleteAppsResponse, error) {
	if systemInfo.Cryptor == nil {
		return nil, status.Error(codes.Unauthenticated, "haven't authenticated yet, please authenticate")
	}
	Debug(dEnc, "delete app request %v\n", req)
	location := decrypt(req.Location)
	Debug(dInfo, "%v, %v, %v\n", req, string(location), string(decrypt(req.Name)))
	switch systemInfo.OS {
	case "darwin":
		out, err := exec.Command("sudo", "rm", "-rf", string(location)).Output()
		checkError(err)
		Debug(dInfo, "%v\n", out)
	case "windows":
		pwd, err := os.Getwd()
		_, err = exec.Command("python",
			pwd+"\\nemon\\scripts\\getAppListWindows.py",
			string(decrypt(req.GetName()))).Output()
		checkError(err)
	case "linux":
		out, err := exec.Command("apt", "remove", string(location)).Output()
		checkError(err)
		Debug(dInfo, "%v\n", out)
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

	Debug(dInfo, "my ip on the network: %v\nhostname: %v\nusername: %v\n",
		ip,
		systemInfo.hostname,
		systemInfo.username,
	)
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)
	go func() {
		<-sigCh
		Debug(dInfo, "\nworker exiting gracefully...\n")
		// worker cleanup if needed
		os.Exit(1)
	}()
	conn, err := net.Listen("tcp", workerAddr)

	checkError(err)

	grpcServer := grpc.NewServer()

	server := workerServer{}

	pb.RegisterWorkerServer(grpcServer, &server)
	Debug(dInfo, "starting gRPC server at port %v...\n", port)
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
