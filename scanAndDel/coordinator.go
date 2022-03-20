package scanAndDel

import (
	"fmt"
	"os"
	"time"
)

func (c *Coordinator) Start() {
	for true {
		fmt.Printf("%v started as a coordinator\n", os.Getpid())
		time.Sleep(500 * time.Millisecond)
	}
}

type Coordinator struct {
	workers  []*Worker
	nWorkers int
}
