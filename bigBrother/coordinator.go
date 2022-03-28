package bigBrother

import (
	"fmt"
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
	screenMu sync.Mutex
}

func (c *Coordinator) SendHeartbeat(worker *Worker, args *GetAppsArgs, reply *GetAppsReply) {
	err := worker.connection.Call("Worker.GetApps", args, reply)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	// use the i/o console exclusively
	c.screenMu.Lock()
	defer c.screenMu.Unlock()

	fmt.Printf("app list received from worker %v\n", worker.ip)
	for _, app := range reply.Applications {
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

func StartCoordinator() {
	fmt.Printf("%v started as a coordinator\n", os.Getpid())
	//connection1, err := rpc.DialHTTP("tcp", "192.168.48.62:8080")
	//checkError(err)
	//connection2, err := rpc.DialHTTP("tcp", "localhost:1235")
	//checkError(err)
	coordinator := Coordinator{
		workers: map[string]*Worker{
			//"192.168.48.62:8080": {
			//	connection: connection1,
			//	port:       8080,
			//},
		},
		//socket:   connection1,
		nWorkers: 0,
		allowed:  map[string]bool{},
	}
	coordinator.BroadcastDiscoveryPings()
	fmt.Printf("%v", coordinator.workers)
	cycle := 1
	for {
		coordinator.BroadcastHeartbeats(cycle)
		time.Sleep(heartbeatInterval)
		cycle++
	}
}
