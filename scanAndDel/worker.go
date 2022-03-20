package scanAndDel

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type Worker struct {
}

func (w *Worker) Start() {
	fmt.Printf("%v started as a worker\n", os.Getpid())
	for true {
		out, err := exec.Command("system_profiler", "-detailLevel", "mini", "SPApplicationsDataType").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v", string(out))
		time.Sleep(1000 * time.Millisecond)
	}
}
