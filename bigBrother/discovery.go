package bigBrother

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	pb "big_brother/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Coordinator) SendDiscoveryPing(ip string) {
	Command := fmt.Sprintf("ping -c 1 -W 1 " + ip + " > /dev/null && echo true || echo false")
	output, err := exec.Command("/bin/sh", "-c", Command).Output()
	checkError(err)
	res := string(output)
	//fmt.Printf("%v\n", len(res))
	if len(res) == 6 {
		return
	}
	if len(res) == 5 {
		address := fmt.Sprintf(ip + port)

		fmt.Printf("%v found\n", address)
		// TODO: global var if big brother installed in .rc
		connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("error occured %v\n", err.Error())
			return
		}
		client := pb.NewWorkerClient(connection)
		c.mu.Lock()
		c.workers[address] = &Worker{ip: ip, connection: connection, client: client}
		c.nWorkers++
		c.mu.Unlock()
	}
}

func (c *Coordinator) BroadcastDiscoveryPings() {
	localIP := GetLocalIP()

	fmt.Printf("\nmy ip: %v\n", localIP)
	vals := strings.Split(localIP, ".")
	mask := vals[0] + "." + vals[1] + "." + vals[2] + "."
	//var wg sync.WaitGroup
	for i := 0; i < 256; i++ {
		ip := mask + strconv.Itoa(i)
		//wg.Add(1)
		// maybe we can show how fast multithreading is here in the slides?
		if ip != localIP {
			go func(IP string) {
				c.SendDiscoveryPing(IP)
				//wg.Done()
			}(ip)
		}
	}
	time.Sleep(5 * time.Second)
	//wg.Wait()
}
