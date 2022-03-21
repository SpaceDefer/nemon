package bigBrother

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type List struct {
}

type Worker struct {
	port       int // replace by ip over wifi
	connection *rpc.Client
}

func (w *Worker) GetApps(_ *GetAppsArgs, reply *GetAppsReply) error {
	out, err := exec.Command("find", "/Applications", "-maxdepth", "3", "-iname", "*.app").Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	r := strings.Split(string(out), "\n")
	var res []ApplicationInfo
	for _, str := range r {
		if len(str) > 0 {
			toAppend := strings.Split(str, "/")
			res = append(res, ApplicationInfo{Name: toAppend[len(toAppend)-1], Location: str})
		}
	}
	//fmt.Printf("%v", res)
	reply.Applications = res
	return nil
}

func (w *Worker) Start() {
	fmt.Printf("%v started as a worker\n", os.Getpid())
	for true {
		time.Sleep(1000 * time.Millisecond)
	}
}

func StartWorker() {
	worker := new(Worker)
	err := rpc.Register(worker)
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Printf("serving from port 1234 concurrently\n")
		err = http.ListenAndServe(":1234", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		defer wg.Done()
		fmt.Printf("serving from port 1235 concurrently\n")
		err = http.ListenAndServe(":1235", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()
}
