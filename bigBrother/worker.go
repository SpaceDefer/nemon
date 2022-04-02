package bigBrother

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	pb "big_brother/protos"

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
}

// GetApps implements GetApps RPC from the generated ProtoBuf file
func (ws *workerServer) GetApps(_ context.Context, _ *pb.GetAppsRequest) (*pb.GetAppsResponse, error) {
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
				applications = append(applications, &pb.GetAppsResponse_ApplicationInfo{Name: toAppend[len(toAppend)-1], Location: str})
			}
		}
	case "windows":
		//out, err = exec.Command("powershell", "-noprofile", "Get-WmiObject").Output()
	case "linux":
		out, err = exec.Command("apt", "list", "--installed").Output()
	default:
		return nil, fmt.Errorf("unrecognised os %v", systemInfo.OS)
	}
	response := &pb.GetAppsResponse{Applications: applications}
	return response, nil
}

// StartWorker handles starting up the worker on the machine
func StartWorker() {
	InitSystemInfo()
	ip := GetLocalIP()
	workerAddr := ip + port
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
