package bigBrother

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"sync"
	"time"
)

type Coordinator struct {
	workers  map[string]*Worker
	nWorkers int
	socket   *rpc.Client
	allowed  map[string]bool
	mu       sync.Mutex
}

func (c *Coordinator) SendHeartbeat(worker *Worker, args *GetAppsArgs, reply *GetAppsReply) {
	err := worker.connection.Call("Worker.GetApps", args, reply)
	if err != nil {
		return
	}

	// use the i/o console exclusively
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Printf("app list received from worker %v\n", worker.port)
	for _, app := range reply.Applications {
		if !c.allowed[app.Name] {
			fmt.Printf("found an app on port [%v] which isn't allowed: %v\n", worker.port, app)
		}
	}
}

func (c *Coordinator) BroadcastHeartbeats() {
	for _, worker := range c.workers {
		fmt.Printf("coordinator sending a heartbeat to port %v\n", worker.port)
		args := &GetAppsArgs{}
		go c.SendHeartbeat(worker, args, &GetAppsReply{})
		fmt.Printf("awaiting response from worker [%v]\n", worker.port)
	}
}

func StartCoordinator() {
	fmt.Printf("%v started as a coordinator\n", os.Getpid())
	connection1, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		return
	}
	connection2, err := rpc.DialHTTP("tcp", "localhost:1235")
	if err != nil {
		log.Fatal(err)
	}
	coordinator := Coordinator{
		workers: map[string]*Worker{
			"localhost:1234": {port: 1234, connection: connection1},
			"localhost:1235": {port: 1235, connection: connection2},
		},
		socket:   connection1,
		nWorkers: 2,
		allowed:  map[string]bool{},
	}

	for true {
		coordinator.BroadcastHeartbeats()
		time.Sleep(heartbeatInterval)
	}
}
