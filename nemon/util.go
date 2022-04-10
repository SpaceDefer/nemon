package nemon

import (
	"bytes"
	"crypto/cipher"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const initialVector = "1234567890123456"

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
				if strings.Split(addrStr, ".")[0] == "192" {
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
	// blockSize := c.BlockSize()
	// data := PKCS5Padding(plaintext, blockSize)
	// enc := make([]byte, len(data))

	ecb := cipher.NewCBCEncrypter(c, []byte(initialVector))
	content := []byte(plaintext)
	content = PKCS5Padding(content, c.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

// decrypt returns an AES decrypted byte array
func decrypt(ciphertext []byte) []byte {
	c := systemInfo.AESCipher

	if c == nil {
		return nil
	}
	ecb := cipher.NewCBCDecrypter(c, []byte(initialVector))
	decrypted := make([]byte, len(ciphertext))
	ecb.CryptBlocks(decrypted, ciphertext)

	data := PKCS5UnPadding(decrypted)
	return data
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

// WebSocket server structs

type Type string

const (
	Alert Type = "ALT"
	Info  Type = "INF"
)

type AlertMessage struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

type DeleteApplicationRequest struct {
	ApplicationName string `json:"applicationName"`
	WorkerIp        string `json:"workerIp"`
	Location        string `json:"location"`
}

type ApplicationInfo struct {
	ApplicationName string `json:"applicationName"`
	Location        string `json:"location"`
}

type WorkerInfo struct {
	Type            Type              `json:"type"`
	ApplicationList []ApplicationInfo `json:"applicationList"`
	WorkerIp        string            `json:"workerIp"`
	Username        string            `json:"username"`
	Hostname        string            `json:"hostname"`
	Os              string            `json:"os"`
}

type DeleteApplicationReply struct {
	Ok bool `json:"ok"`
}
