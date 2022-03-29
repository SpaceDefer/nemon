package bigBrother

import (
	"fmt"
	"net/http"
	"net/rpc"
	"os"
	"os/exec"
	"strings"
)

type List struct {
}

type Worker struct {
	ip         string // replace by ip:port over Wi-Fi
	connection *rpc.Client
}

func (w *Worker) GetApps(_ *GetAppsArgs, reply *GetAppsReply) error {
	var err error
	var out []byte
	var res []ApplicationInfo
	switch os.Getenv("OS") {
	case "darwin":
		out, err = exec.Command("find", "/Applications", "-maxdepth", "3", "-iname", "*.app").Output()
		checkError(err)
		r := strings.Split(string(out), "\n")
		for _, str := range r {
			if len(str) > 0 {
				toAppend := strings.Split(str, "/")
				res = append(res, ApplicationInfo{Name: toAppend[len(toAppend)-1], Location: str})
			}
		}
	case "windows":
		//out, err = exec.Command("powershell", "-noprofile", "Get-WmiObject").Output()
	case "linux":
		out, err = exec.Command("apt", "list", "--installed").Output()
	default:
		return nil
	}
	//fmt.Printf("%v", res)
	reply.Applications = res
	return nil
}

func StartWorker() {
	worker := new(Worker)
	err := rpc.Register(worker)
	checkError(err)
	fmt.Println(os.Getenv("OS"))
	//_, err = exec.Command("export", "IS_WORKER=\"true\"").Output()
	//checkError(err)
	rpc.HandleHTTP()
	fmt.Printf("serving from port %v\n", port)
	err = http.ListenAndServe(port, nil)
	checkError(err)
	fmt.Printf("worker exiting...\n")
}
