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

func (w *Worker) GetApps(_ *GetAppsArgs, reply *GetAppsReply) error {
	var err error
	var out []byte
	var res []ApplicationInfo
	switch os.Getenv("OS") {
	case "darwin":
		out, err = exec.Command("find", "/Applications", "-maxdepth", "3", "-iname", "*.app").Output()
		checkError(err)
		r := strings.Split(string(out), "\n")
		for _, str := range r {
			if len(str) > 0 {
				toAppend := strings.Split(str, "/")
				res = append(res, ApplicationInfo{Name: toAppend[len(toAppend)-1], Location: str})
			}
		}
	case "windows":
		//out, err = exec.Command("powershell", "-noprofile", "Get-WmiObject").Output()
	case "linux":
		out, err = exec.Command("apt", "list", "--installed").Output()
	default:
		return nil
	}
	//fmt.Printf("%v", res)
	reply.Applications = res
	return nil
}

func (ws *workerServer) GetApps(ctx context.Context, _ *pb.GetAppsRequest) (*pb.GetAppsResponse, error) {
	fmt.Printf("got a GetApps gRPC with context %v", ctx)
	var applications []*pb.GetAppsResponse_ApplicationInfo
	applications = append(applications, &pb.GetAppsResponse_ApplicationInfo{Name: "tanmay", Location: "/tanmay"})
	response := &pb.GetAppsResponse{Applications: applications}
	return response, nil
}

const workerAddr = "localhost" + port

func StartWorker() {
	//worker := new(Worker)
	//err := rpc.Register(worker)
	//checkError(err)
	//fmt.Println(os.Getenv("OS"))
	////_, err = exec.Command("export", "IS_WORKER=\"true\"").Output()
	////checkError(err)
	//rpc.HandleHTTP()
	//fmt.Printf("serving from port %v\n", port)
	//err = http.ListenAndServe(port, nil)
	//checkError(err)
	//fmt.Printf("worker exiting...\n")
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
