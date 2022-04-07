package nemon

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "nemon/protos"

	"google.golang.org/grpc"
)

// Handshake performs our own handshake protocol with the established connection
func (c *Coordinator) Handshake(connection *grpc.ClientConn) (*pb.GetSysInfoResponse, pb.WorkerClient, error) {
	client := pb.NewWorkerClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetSysInfo(ctx, &pb.GetSysInfoRequest{Key: systemInfo.nemonKey})
	if err != nil {
		log.Printf("%v\n", err)
		if err := connection.Close(); err != nil {
			fmt.Printf("can't close connection\n")
		}
		return nil, nil, err
	}
	return response, client, nil
}
