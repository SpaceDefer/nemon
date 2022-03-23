package bigBrother

import (
	"fmt"
	"net/rpc"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

func (c *Coordinator) SendDiscoveryPing(ip string) {
	Command := fmt.Sprintf("ping -c 1 -W 1 " + ip + " > /dev/null && echo true || echo false")
	output, err := exec.Command("/bin/sh", "-c", Command).Output()
	checkError(err)
	res := string(output)
	fmt.Printf("found ip %v? %v", ip, res)
	if res == "true" {
		address := fmt.Sprintf(ip + port)
		connection, err := rpc.DialHTTP("tcp", address)
		checkError(err)
		c.mu.Lock()
		c.nWorkers++
		c.workers[address] = &Worker{port: 8080, connection: connection}
		c.mu.Unlock()
	}

	//if string(output) == "true" {
	//	return true
	//}
	//return false
}

func (c *Coordinator) BroadcastDiscoveryPings() {
	ip := GetLocalIP()
	fmt.Printf("\nmy ip: %v\n", ip)
	vals := strings.Split(ip, ".")
	fmt.Printf("%v\n", vals)
	mask := vals[0] + "." + vals[1] + "." + vals[2] + "."
	var wg sync.WaitGroup
	for i := 0; i < 256; i++ {
		ip := mask + strconv.Itoa(i)
		wg.Add(1)

		// maybe we can show how fast multithreading is here in the slides?
		go func(IP string) {
			defer wg.Done()
			c.SendDiscoveryPing(ip)
		}(ip)
	}
	wg.Wait()
	//time.Sleep(10 * time.Second)
}
