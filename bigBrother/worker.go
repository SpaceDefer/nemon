package bigBrother

import (
	"fmt"
	"net/http"
	"net/rpc"
	"os/exec"
	"strings"
	"sync"
)

type List struct {
}

type Worker struct {
	port       int // replace by ip:port over Wi-Fi
	connection *rpc.Client
}

func (w *Worker) GetApps(_ *GetAppsArgs, reply *GetAppsReply) error {
	out, err := exec.Command("find", "/Applications", "-maxdepth", "3", "-iname", "*.app").Output()
	checkError(err)
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

func StartWorker() {
	worker := new(Worker)
	err := rpc.Register(worker)
	checkError(err)
	rpc.HandleHTTP()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("serving from port %v concurrently\n", port)
		err = http.ListenAndServe(port, nil)
		checkError(err)
	}()
	wg.Wait()
}
