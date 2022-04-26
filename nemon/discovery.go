package nemon

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// SendDiscoveryPing sends a single discovery ping to the give ip
func (c *Coordinator) SendDiscoveryPing(ip string) {
	Command := fmt.Sprintf("ping -c 1 -W 1 " + ip + " > /dev/null && echo true || echo false")
	var output []byte
	var err error
	switch systemInfo.OS {
	case "windows":
		output, err = exec.Command("ping", "-c", "1", "-W", "1", ip, " > nul && echo true || echo false").Output()
	default:
		output, err = exec.Command("/bin/sh", "-c", Command).Output()
	}
	checkError(err)
	res := strings.TrimSpace(string(output))
	if res == "false" {
		return
	}
	address := fmt.Sprintf(ip + port)

	Debug(dInfo, "%v found\n", ip)
	// TODO: global var if big brother installed in .rc
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		Debug(dInfo, "error occured %v\n", err.Error())
		return
	}

	response, client, err := c.Handshake(connection)

	if err != nil {
		Debug(dInfo, "handshake was unsuccessful with %v\n", ip)
		if err := connection.Close(); err != nil {
			return
		}
		fmt.Println(err.Error())
		return
	}

	Debug(dInfo, "handshake was successful with %v\n", ip)
	workerSysInfo := response.WorkerSysInfo
	c.mu.Lock()
	c.workers[ip] = &Worker{
		ip:         ip,
		connection: connection,
		client:     client,
		username:   string(decrypt(workerSysInfo.Username)),
		os:         string(decrypt(workerSysInfo.Os)),
		hostname:   string(decrypt(workerSysInfo.Hostname)),
	}
	c.nWorkers++
	c.mu.Unlock()
}

// BroadcastDiscoveryPings to available IP addresses on the network
func (c *Coordinator) BroadcastDiscoveryPings() {
	localIP := GetLocalIP()

	Debug(dInfo, "\nmy ip: %v\n", localIP)
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
