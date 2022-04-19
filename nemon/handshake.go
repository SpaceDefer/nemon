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

	"github.com/1Password/srp"

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

func (c *Coordinator) _Handshake(connection *grpc.ClientConn) (*pb.GetSysInfoResponse, pb.WorkerClient, error) {
	client := pb.NewWorkerClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.IsEnrolled(ctx, &pb.IsEnrolledRequest{
		Key: systemInfo.nemonKey,
	})
	if err != nil {
		return nil, nil, err
	}
	if !response.Enrolled {
		err := c.Enrollment(client)
		if err != nil {
			return nil, nil, err
		}
	}
	// authenticate and verify
	err = c.Authentication(client)
	if err != nil {
		return nil, nil, nil
	}
	return nil, nil, nil
}

// Enrollment enrolls the Coordinator with the Worker
func (c *Coordinator) Enrollment(client pb.WorkerClient) error {
	group := srp.RFC5054Group3072
	pw := "temp!" // make a random password

	enrollmentCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	salt := make([]byte, 8)
	if n, err := rand.Read(salt); err != nil {
		return err
	} else if n != 8 {
		return fmt.Errorf("couldn't generate an 8 byte salt")
	}

	username := "random" // random username

	// save the
	// password and the
	// username in worker(server) sysInfo for later use

	x := srp.KDFRFC5054(salt, username, pw)

	firstClient := srp.NewSRPClient(srp.KnownGroups[group], x, nil)
	if firstClient == nil {
		return fmt.Errorf("couldn't create a srpClient")
	}
	v, err := firstClient.Verifier() // Verifier, err
	if err != nil {
		return err
	}
	// make a grpc call to save v on the Worker
	_, err = client.SaveEnrollmentInfo(enrollmentCtx, &pb.SaveEnrollmentInfoRequest{
		Salt:     salt,
		Verifier: v.Bytes(),
		SRPGroup: int64(group),
	})

	return err
}

func (c *Coordinator) Authentication(client pb.WorkerClient) error {
	// grpc call to request the salt and SRP Group from the Worker
	authCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := client.GetSaltAndSRP(authCtx, &pb.GetSaltAndSRPRequest{})
	if err != nil {
		return err
	}
	fmt.Printf("salt %v\ngroup %v\n", string(response.Salt), response.SRPGroup)
	// fetch the master password and secret key (username?) from the sysInfo
	//pw, secretKey := systemInfo.Password, systemInfo.SecretKey
	//x := srp.KDFRFC5054(response.Salt, secretKey, pw)

	return nil
}
