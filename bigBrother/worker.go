package bigBrother

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"sync"
	"time"
)

type Worker struct {
	port       int
	connection *rpc.Client
}

func (w *Worker) GetApps(args *Args, reply *Reply) error {
	//out, err := exec.Command("system_profiler", "-xml", "-detailLevel", "mini", "SPApplicationsDataType").Output()
	//if err != nil {
	//	log.Fatal(err)
	//	return err
	//}
	r := []string{"Tanmay"}
	reply.Applications = r
	//fmt.Printf("%v", string(out))
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
		err = http.ListenAndServe(":1234", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		defer wg.Done()
		err = http.ListenAndServe(":1235", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()
	//time.Sleep(5 * time.Second)
}
