package bigBrother

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"time"
)

type Coordinator struct {
	workers  []*Worker
	nWorkers int
	socket   *rpc.Client
	allowed  map[string]bool
}

func (c *Coordinator) SendHeartbeat(worker int, args *GetAppsArgs, reply *GetAppsReply) {
	err := c.workers[worker].connection.Call("Worker.GetApps", args, reply)
	if err != nil {
		return
	}
	fmt.Printf("app list received from port %v\n", c.workers[worker].port)
	for _, app := range reply.Applications {
		if !c.allowed[app] {
			fmt.Printf("found an app on port [%v] which isn't allowed: %v\n", c.workers[worker].port, app)
		}
	}
}

func (c *Coordinator) BroadcastHeartbeats() {
	for worker := 0; worker < c.nWorkers; worker++ {
		fmt.Printf("coordinator sending a heartbeat to port %v\n", c.workers[worker].port)
		args := &GetAppsArgs{}
		go c.SendHeartbeat(worker, args, &GetAppsReply{})
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
		workers: []*Worker{
			{port: 1234, connection: connection1},
			{port: 1235, connection: connection2}},
		socket:   connection1,
		nWorkers: 2,
		allowed:  map[string]bool{},
	}

	for true {
		coordinator.BroadcastHeartbeats()
		time.Sleep(10 * time.Second)
	}
}
