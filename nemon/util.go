package nemon

import (
	"bytes"
	"crypto/cipher"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
	"strings"
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

	addrs, err := net.InterfaceAddrs()

    if err != nil {
        fmt.Println(err)
    }
	
    for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
				addrStr := ipnet.IP.String()
                if strings.Split(addrStr,".")[0] == "192" {
					return addrStr
				}
            }
        }
    }
	return ""
}

// SystemInfo contains frequently required information about the system on which
// the software is running
type SystemInfo struct {
	OS        string
	hostname  string
	username  string
	nemonKey  int64
	AESKey    []byte
	AESCipher cipher.Block
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

// encrypt returns an AES encrypted byte array
func encrypt(plaintext []byte) []byte {
	c := systemInfo.AESCipher

	if c == nil {
		return nil
	}
	blockSize := c.BlockSize()
	data := PKCS5Padding(plaintext, blockSize)
	enc := make([]byte, len(data))

	c.Encrypt(enc, data)
	return enc
}

// decrypt returns an AES decrypted byte array
func decrypt(ciphertext []byte) []byte {
	c := systemInfo.AESCipher

	if c == nil {
		return nil
	}

	dec := make([]byte, len(ciphertext))

	c.Decrypt(dec, ciphertext)

	data := PKCS5UnPadding(dec)
	return data
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// WebSocket server structs

type DeleteApplicationRequest struct {
	ApplicationName string `json:"applicationName"`
	WorkerIp        string `json:"workerIp"`
}

type ApplicationInfo struct {
	ApplicationName string `json:"applicationName"`
	Location        string `json:"location"`
}

type WorkerInfo struct {
	ApplicationList []ApplicationInfo `json:"applicationList"`
	WorkerIp        string            `json:"workerIp"`
	Username        string            `json:"username"`
	Hostname        string            `json:"hostname"`
	Os              string            `json:"os"`
}

type DeleteApplicationReply struct {
	Ok bool `json:"ok"`
}
