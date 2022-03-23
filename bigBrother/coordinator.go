package bigBrother

import (
	"fmt"
	"net/rpc"
	"os"
	"sync"
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

func (c *Coordinator) BroadcastHeartbeats(cnt int) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("heartbeat cycle number %v\n", cnt)
	for _, worker := range c.workers {
		fmt.Printf("coordinator sending a heartbeat to port %v\n", worker.port)
		args := &GetAppsArgs{}
		go c.SendHeartbeat(worker, args, &GetAppsReply{})
	}
}

func StartCoordinator() {
	fmt.Printf("%v started as a coordinator\n", os.Getpid())
	//connection1, err := rpc.DialHTTP("tcp", "localhost:1234")
	//checkError(err)
	//connection2, err := rpc.DialHTTP("tcp", "localhost:1235")
	//checkError(err)
	coordinator := Coordinator{
		workers: map[string]*Worker{},
		//socket:   connection1,
		nWorkers: 0,
		allowed:  map[string]bool{},
	}
	coordinator.BroadcastDiscoveryPings()
	//i := 1
	//for {
	//	coordinator.BroadcastHeartbeats(i)
	//	time.Sleep(heartbeatInterval)
	//	i++
	//}
}
