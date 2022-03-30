package bigBrother

import (
	"context"
	"fmt"
	"os"
	"sync"
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

func (c *Coordinator) SendHeartbeat(worker *Worker, args *GetAppsArgs, reply *GetAppsReply) {
	//err := worker.connection.Call("Worker.GetApps", args, reply)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := worker.client.GetApps(ctx, &pb.GetAppsRequest{})
	checkError(err)
	// use the i/o console exclusively
	c.screenMu.Lock()
	defer c.screenMu.Unlock()

	fmt.Printf("app list received from worker %v\n", worker.ip)
	for _, app := range response.Applications {
		if !c.allowed[app.Name] {
			fmt.Printf("found an app on ip [%v] which isn't allowed: %v\n", worker.ip, app)
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
		args := &GetAppsArgs{}
		go c.SendHeartbeat(worker, args, &GetAppsReply{})
	}
}

func (c *Coordinator) Cleanup() {
	for _, worker := range c.workers {
		err := worker.connection.Close()
		checkError(err)
	}
}

func StartCoordinator() {
	fmt.Printf("%v started as a coordinator\n", os.Getpid())
	//connection2, err := rpc.DialHTTP("tcp", "localhost:1235")
	//checkError(err)
	coordinator := Coordinator{
		workers: map[string]*Worker{
			//"192.168.48.62:8080": {
			//	connection: connection1,
			//	port:       8080,
			//},
		},
		nWorkers: 0,
		allowed:  map[string]bool{},
	}
	coordinator.BroadcastDiscoveryPings()
	fmt.Printf("workers found: %v\n", coordinator.workers)
	if coordinator.nWorkers > 0 {
		cycle := 1
		for cycle < 4 {
			coordinator.BroadcastHeartbeats(cycle)
			time.Sleep(heartbeatInterval)
			cycle++
		}
	}
	coordinator.Cleanup()
	fmt.Printf("coordinator exiting...\n")
}
