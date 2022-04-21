package nemon

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	pb "nemon/protos"

	"golang.org/x/crypto/chacha20poly1305"

	"github.com/1Password/srp"

	"google.golang.org/grpc"
)

// Handshake performs our own handshake protocol (SRP with XChaCha20-Poly1305) with the established connection
// returns a GetSysInfoResponse struct containing SystemInfo struct of the worker,
// a WorkerClient for the connection and an error if unsuccessful
func (c *Coordinator) Handshake(connection *grpc.ClientConn) (*pb.GetSysInfoResponse, pb.WorkerClient, error) {
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
		fmt.Printf("enrolling\n")
		err := c.Enrollment(client)
		if err != nil {
			return nil, nil, err
		}
		fmt.Printf("enrollment successful\n")
	}
	// authenticate and verify
	fmt.Printf("authenticating\n")
	err = c.Authentication(client)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("authentication and verification successful\n")

	sysInfoResponse, err := client.GetSysInfo(ctx, &pb.GetSysInfoRequest{})

	fmt.Println(sysInfoResponse)

	return sysInfoResponse, client, nil
}

// Enrollment enrolls the Coordinator with the Worker
func (c *Coordinator) Enrollment(client pb.WorkerClient) error {
	group := srp.RFC5054Group3072
	enrollmentCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	salt := make([]byte, 8)
	if n, err := rand.Read(salt); err != nil {
		return err
	} else if n != 8 {
		return fmt.Errorf("couldn't generate an 8 byte salt")
	}

	// TODO: REMOVE THIS!
	x := KDF(salt)

	//str := username + ":" + pw

	//hash := sha256.New()
	//xBytes := pbkdf2.Key([]byte(str), salt, 4096, 32, sha1.New)
	//x := new(big.Int)
	//x.SetBytes(xBytes)
	//fmt.Println(x)
	//fmt.Println(_x)

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
	authCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.GetSaltAndSRP(authCtx, &pb.GetSaltAndSRPRequest{})
	group, salt := response.SRPGroup, response.Salt
	if err != nil {
		return err
	}
	x := srp.KDFRFC5054(salt, username, pw)

	srpClient := srp.NewSRPClient(srp.KnownGroups[int(group)], x, nil)

	A := srpClient.EphemeralPublic()

	// TODO: handle different errors and codes
	exchangeEphemeralResponse, err := client.ExchangeEphemeralPublic(authCtx, &pb.ExchangeEphemeralPublicRequest{
		A: A.Bytes(),
	})
	BBytes, serverProof := exchangeEphemeralResponse.B, exchangeEphemeralResponse.ServerProof
	B := new(big.Int)
	B.SetBytes(BBytes)
	if err = srpClient.SetOthersPublic(B); err != nil {
		return err
	}

	clientKey, err := srpClient.Key()

	if err != nil || clientKey == nil {
		return fmt.Errorf("couldn't make the client key\n%v\n", err.Error())
	}

	fmt.Printf("authentication successful, continuing with verification... %v\n", len(serverProof))
	if !srpClient.GoodServerProof(salt, username, serverProof) {
		return fmt.Errorf("bad proof from server")
	}

	clientProof, err := srpClient.ClientProof()
	if err != nil {
		return err
	}

	_, err = client.VerifyClientProof(authCtx, &pb.VerifyClientProofRequest{
		ClientProof: clientProof,
	})
	if err != nil {
		return err
	}

	fmt.Printf("verification successful!\n")

	clientCryptor, _ := chacha20poly1305.NewX(clientKey)

	systemInfo.Cryptor = clientCryptor
	return nil
}
