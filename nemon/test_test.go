package nemon

import (
	"net"
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	myIP := GetLocalIP()
	var ip string
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		t.Error("error while getting the ip")
	}
	for _, address := range addresses {
		// check the address type and if it is not a loop-back the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	if ip == myIP {
		t.Errorf("ip was incorrect")
	}
}
