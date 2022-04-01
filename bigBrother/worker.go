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

type List struct {
}

type workerServer struct {
	pb.UnimplementedWorkerServer
}

type Worker struct {
	ip         string // replace by ip:port over Wi-Fi
	client     pb.WorkerClient
	connection *grpc.ClientConn
}

func (ws *workerServer) GetApps(_ context.Context, _ *pb.GetAppsRequest) (*pb.GetAppsResponse, error) {
	fmt.Printf("got a GetApps gRPC\n")
	var err error
	var out []byte
	var applications []*pb.GetAppsResponse_ApplicationInfo

	OS := os.Getenv("OS")

	switch OS {
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
		return nil, fmt.Errorf("unrecognised os %v", OS)
	}
	response := &pb.GetAppsResponse{Applications: applications}
	return response, nil
}

func StartWorker() {
	ip := GetLocalIP()
	workerAddr := ip + port
	fmt.Printf("my ip on the network: %v\n", ip)
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
