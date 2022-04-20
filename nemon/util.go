package nemon

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/1Password/srp"
)

const initialVector = "1234567890123456"

// heartbeatInterval is the duration the Coordinator waits to send RPCs
const heartbeatInterval = 10 * time.Second

const devHeartbeatInterval = 5 * time.Second

const username = "coordinator"

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
	Dev       bool
	OS        string
	hostname  string
	username  string
	nemonKey  int64
	AESKey    []byte       // deprecate with RSA handshake
	AESCipher cipher.Block // deprecate with RSA handshake
	SecretKey string
	Password  string
	ConfigDir string
	Cryptor   cipher.AEAD
}

// systemInfo is an instance of SystemInfo
var systemInfo SystemInfo

// SRPServerInfo stores the long term persistent info stored on the Worker
// during the runtime of the program
type SRPServerInfo struct {
	Verifier  *big.Int
	Group     int
	Salt      []byte
	Server    *srp.SRP
	ServerKey []byte
}

// srpServerInfo is an instance of SRPServerInfo
var srpServerInfo SRPServerInfo

// SRPClientInfo stores the srp info on the Coordinator
// during the runtime of the program
type SRPClientInfo struct {
	Verifier *big.Int
	Group    int
}

// srpClientInfo is an instance of SRPClientInfo
var srpClientInfo SRPClientInfo

// InitSystemInfo initialises the SystemInfo struct for the system
func InitSystemInfo() {
	systemInfo.OS = os.Getenv("OS")
	systemInfo.hostname = os.Getenv("HOSTNAME")
	systemInfo.username = os.Getenv("USERNAME")
	systemInfo.ConfigDir = os.Getenv("CONFIG_DIR")
	if os.Getenv("DEV") == "dev" {
		systemInfo.Dev = true
	} else {
		systemInfo.Dev = false
	}
	key, err := strconv.ParseInt(os.Getenv("NEMONKEY"), 10, 64)
	if err != nil {
		log.Fatalf("issues with the key\n")
	}
	systemInfo.nemonKey = key
}

func encrypt(plaintext []byte) []byte {
	cryptor := systemInfo.Cryptor
	nonce := make([]byte, cryptor.NonceSize(), cryptor.NonceSize()+len(plaintext)+cryptor.Overhead())
	if _, err := rand.Read(nonce); err != nil {
		fmt.Printf("couldn't encrypt: %v\n", err.Error())
		return nil
	}

	return cryptor.Seal(nonce, nonce, plaintext, nil)
}

func decrypt(ciphermsg []byte) []byte {
	cryptor := systemInfo.Cryptor
	nonce, ciphertext := ciphermsg[:cryptor.NonceSize()], ciphermsg[cryptor.NonceSize():]

	plaintext, err := cryptor.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		fmt.Printf("couldn't decrypt: %v\n", err.Error())
		return nil
	}

	return plaintext
}

// encrypt returns an AES encrypted byte array
func _encrypt(plaintext []byte) []byte {
	c := systemInfo.AESCipher

	if c == nil {
		return nil
	}

	ecb := cipher.NewCBCEncrypter(c, []byte(initialVector))
	plaintext = PKCS5Padding(plaintext, c.BlockSize())
	crypted := make([]byte, len(plaintext))
	ecb.CryptBlocks(crypted, plaintext)

	return crypted
}

// decrypt returns an AES decrypted byte array
func _decrypt(ciphertext []byte) []byte {
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
	if len(encrypt) < 1 {
		return nil
	}
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

// WebSocket Server structs

// Type of the WebsocketServer message
type Type string

// Type of messages to send through Websockets
const (
	Alert       Type = "ALT" // Alert the WebsocketServer client
	Info        Type = "INF" // Send Info to the WebsocketServer client
	Delete      Type = "DEL" // Delete application
	Acknowledge Type = "ACK" // Acknowledge a message sent or received
)

type AlertMessage struct {
	Type     Type   `json:"type"`
	Message  string `json:"message"`
	WorkerIp string `json:"workerIp"`
}

type DeleteApplicationRequest struct {
	Type            Type   `json:"type"`
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

type EnrollmentInfo struct {
	SRPGroup int64
	Salt     []byte
	Verifier []byte
}

type DeleteApplicationReply struct {
	Type    Type   `json:"type"`
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type WorkerStatusMessage struct {
	Type     Type   `json:"type"`
	WorkerIp string `json:"workerIp"`
	Status   Status `json:"status"`
}
