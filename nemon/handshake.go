package nemon

import (
	"context"
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "nemon/protos"

	"google.golang.org/grpc"
)

// Handshake performs our own handshake protocol with the established connection
// returns a GetSysInfoResponse struct containing SystemInfo struct of the worker and an AESKey if successful,
// a WorkerClient for the connection and an error if unsuccessful
func (c *Coordinator) Handshake(connection *grpc.ClientConn) (*pb.GetSysInfoResponse, pb.WorkerClient, error) {
	client := pb.NewWorkerClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return nil, nil, err
	}

	publicKey := privateKey.PublicKey

	response, err := client.GetSysInfo(ctx, &pb.GetSysInfoRequest{
		Key:        systemInfo.nemonKey,
		PublicKeyN: publicKey.N.String(),
		PublicKeyE: strconv.Itoa(publicKey.E),
	})

	if err != nil {
		return nil, nil, err
	}

	fmt.Printf("%v\n", response)

	hash := sha512.New()
	AESKey, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, response.AESKey, nil)

	AESCipher, err := aes.NewCipher(AESKey)

	if err != nil {
		return nil, nil, err
	}

	systemInfo.AESCipher = AESCipher
	systemInfo.AESKey = AESKey

	if err != nil {
		log.Printf("%v\n", err)
		if err := connection.Close(); err != nil {
			fmt.Printf("can't close connection\n")
		}
		return nil, nil, err
	}
	return response, client, nil
}
