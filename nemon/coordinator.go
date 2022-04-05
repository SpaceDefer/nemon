package nemon

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	pb "nemon/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Coordinator struct implements the coordinator
type Coordinator struct {
	workers  map[string]*Worker // workers map ip addresses to Worker structs
	nWorkers int                // nWorkers gives the number of workers
	allowed  map[string]bool    // list of all allowed applications
	mu       sync.Mutex         // mu mutex to prevent data races in the Coordinator's data
	screenMu sync.Mutex         // screenMu to print on the screen exclusively
}

var deleteChan chan DeleteApplicationRequest

func ClearChannels() {
	deleteChan = make(chan DeleteApplicationRequest)
}

// ListenDeleteApplication wrapper for the server to call
func (c *Coordinator) ListenDeleteApplication() {
	for {
		var req DeleteApplicationRequest
		req, ok := <-deleteChan
		if !ok {
			fmt.Println("ch error")
			continue
		}
		fmt.Printf("received an rpc, going to delete %v on %v\n", req.ApplicationName, req.WorkerIp)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		worker := c.workers[req.WorkerIp]
		if worker == nil {
			fmt.Printf("didn't find that specific ip\n")
		}
		fmt.Printf("worker found %v\n", worker)
		response, err := worker.client.DeleteApp(ctx, &pb.DeleteAppsRequest{Name: req.ApplicationName, Key: systemInfo.nemonKey})
		if err != nil {
			log.Fatalf(err.Error())
			return
		}
		fmt.Printf("%v\n", response.Ok)
	}
}

// SendHeartbeat to a single Worker
func (c *Coordinator) SendHeartbeat(worker *Worker) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := worker.client.GetApps(ctx, &pb.GetAppsRequest{Key: systemInfo.nemonKey})
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}
	// use the i/o console exclusively
	c.screenMu.Lock()
	defer c.screenMu.Unlock()

	fmt.Printf("app list received from worker %v\n", worker.ip)
	for _, app := range response.Applications {
		if !c.allowed[app.GetName()] {
			fmt.Printf("found an app on ip [%v] which isn't allowed: %v\n", worker.ip, app.GetName())
		}
	}
}

// BroadcastHeartbeats to multiple workers, cycle signifies the heartbeat cycle
func (c *Coordinator) BroadcastHeartbeats(cycle int) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("heartbeat cycle number %v\n", cycle)
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, worker := range c.workers {
		fmt.Printf("coordinator sending a heartbeat to ip %v\n", worker.ip)
		go c.SendHeartbeat(worker)
	}
}

// Cleanup closes the connections in the event of a shutdown
func (c *Coordinator) Cleanup() {
	fmt.Printf("running cleanup...\n")
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, worker := range c.workers {
		err := worker.connection.Close()
		checkError(err)
	}
}

// StartCoordinator starts up a Coordinator process
func StartCoordinator() {
	InitSystemInfo()
	StartServer()
	ClearChannels()
	fmt.Printf("%v started as a coordinator\n", os.Getpid())
	connection, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	checkError(err)
	coordinator := Coordinator{
		workers: map[string]*Worker{
			"localhost": {
				connection: connection,
				client:     pb.NewWorkerClient(connection),
				ip:         "localhost",
			},
		},
		nWorkers: 0,
		allowed:  map[string]bool{},
	}
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)
	go func() {
		<-sigCh
		fmt.Printf("\ncoordinator exiting gracefully...\n")
		coordinator.Cleanup()
		os.Exit(1)
	}()
	//
	//setupRoutes()
	//log.Fatal(http.ListenAndServe(":4000", nil))

	//coordinator.BroadcastDiscoveryPings()
	coordinator.mu.Lock()
	nWorkers := coordinator.nWorkers
	workers := coordinator.workers
	coordinator.mu.Unlock()
	fmt.Printf("number of workers found: %v\nworkers: %v\n", nWorkers, workers)
	go coordinator.ListenDeleteApplication()
	if nWorkers > -1 {
		cycle := 1
		for cycle < 20 {
			coordinator.BroadcastHeartbeats(cycle)
			time.Sleep(heartbeatInterval)
			cycle++
		}
	}
}
