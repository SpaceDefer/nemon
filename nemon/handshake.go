package nemon

import (
	"context"
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
func (c *Coordinator) Handshake(connection *grpc.ClientConn) (*pb.GetSysInfoResponse, pb.WorkerClient, error) {
	client := pb.NewWorkerClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
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

	hash := sha512.New()
	AESKey, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, response.AESKey, nil)

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
