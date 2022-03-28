package bigBrother

import (
	"fmt"
	"net/rpc"
	"os/exec"
	"strconv"
	"strings"
	"time"
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

		connection, err := rpc.DialHTTP("tcp", address)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.mu.Lock()
		c.nWorkers++
		c.workers[address] = &Worker{ip: ip, connection: connection}
		c.mu.Unlock()
	}
}

func (c *Coordinator) BroadcastDiscoveryPings() {
	ip := GetLocalIP()

	fmt.Printf("\nmy ip: %v\n", ip)
	vals := strings.Split(ip, ".")
	fmt.Printf("%v\n", vals)
	mask := vals[0] + "." + vals[1] + "." + vals[2] + "."
	//var wg sync.WaitGroup
	for i := 0; i < 256; i++ {
		ip := mask + strconv.Itoa(i)
		//wg.Add(1)
		// maybe we can show how fast multithreading is here in the slides?
		go func(IP string) {
			c.SendDiscoveryPing(IP)
			//wg.Done()
		}(ip)
	}
	time.Sleep(5 * time.Second)
	//wg.Wait()
}
