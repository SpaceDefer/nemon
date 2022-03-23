package bigBrother

import (
	"fmt"
	"net"
	"time"
)

const heartbeatInterval = 10 * time.Second
const port = ":8080"

type ApplicationInfo struct {
	Name     string
	Location string
}

func checkError(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}

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
