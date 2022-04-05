package nemon

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

// heartbeatInterval is the duration the Coordinator waits to send RPCs
const heartbeatInterval = 10 * time.Second

// default port
const port = ":8080"

// checkError helper
func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// GetLocalIP gets the IP address on the connection
func GetLocalIP() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addresses {
		// check the address type and if it is not a loop-back the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// SystemInfo contains frequently required information about the system on which
// the software is running
type SystemInfo struct {
	OS       string
	hostname string
	username string
	nemonKey int64
}

var systemInfo SystemInfo

// InitSystemInfo initialises the SystemInfo struct for the system
func InitSystemInfo() {
	systemInfo.OS = os.Getenv("OS")
	systemInfo.hostname = os.Getenv("HOSTNAME")
	systemInfo.username = os.Getenv("USERNAME")
	key, err := strconv.ParseInt(os.Getenv("NEMONKEY"), 10, 64)
	if err != nil {
		log.Fatalf("issues with the key\n")
	}
	systemInfo.nemonKey = key
}

// WebSocket server structs

type DeleteApplicationRequest struct {
	ApplicationName string `json:"applicationName"`
	WorkerIp        string `json:"workerIp"`
}

type DeleteApplicationReply struct {
	Ok bool `json:"ok"`
}
