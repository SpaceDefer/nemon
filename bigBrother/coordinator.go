package bigBrother

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	pb "big_brother/protos"
)

type Coordinator struct {
	workers  map[string]*Worker
	nWorkers int
	allowed  map[string]bool
	mu       sync.Mutex
	screenMu sync.Mutex
}

func (c *Coordinator) SendHeartbeat(worker *Worker) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := worker.client.GetApps(ctx, &pb.GetAppsRequest{})
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

func (c *Coordinator) BroadcastHeartbeats(cnt int) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("heartbeat cycle number %v\n", cnt)
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, worker := range c.workers {
		fmt.Printf("coordinator sending a heartbeat to ip %v\n", worker.ip)
		go c.SendHeartbeat(worker)
	}
}

func (c *Coordinator) Cleanup() {
	fmt.Printf("running cleanup...\n")
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, worker := range c.workers {
		err := worker.connection.Close()
		checkError(err)
	}
}

func StartCoordinator() {
	fmt.Printf("%v started as a coordinator\n", os.Getpid())
	//connection, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//checkError(err)
	coordinator := Coordinator{
		workers: map[string]*Worker{
			//"localhost:8080": {
			//	connection: connection,
			//	client:     pb.NewWorkerClient(connection),
			//	ip:         "localhost",
			//},
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
	coordinator.BroadcastDiscoveryPings()
	coordinator.mu.Lock()
	nWorkers := coordinator.nWorkers
	workers := coordinator.workers
	coordinator.mu.Unlock()
	fmt.Printf("workers found: %v\n", workers)
	if nWorkers > 0 {
		cycle := 1
		for cycle < 4 {
			coordinator.BroadcastHeartbeats(cycle)
			time.Sleep(heartbeatInterval)
			cycle++
		}
	}
}
