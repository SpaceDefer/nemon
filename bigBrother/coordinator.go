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
}

func (c *Coordinator) SendHeartbeat(worker int, args *Args, reply *Reply) {
	err := c.workers[worker].connection.Call("Worker.GetApps", args, reply)
	if err != nil {
		return
	}
	fmt.Printf("[%v]: %v received\n", c.workers[worker].port, reply.Applications)
}

func (c *Coordinator) BroadcastHeartbeats() {
	for worker := 0; worker < c.nWorkers; worker++ {
		fmt.Printf("%v\n", c.workers[worker].port)
		args := &Args{}
		go c.SendHeartbeat(worker, args, &Reply{})
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
	}
	for true {
		coordinator.BroadcastHeartbeats()
		time.Sleep(1000 * time.Millisecond)
	}
}
