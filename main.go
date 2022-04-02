package main

import (
	"flag"
	"log"
	"os"
	"os/user"
	"runtime"
	"strings"

	"big_brother/bigBrother"
)

var (
	mode = flag.String("mode", "worker", "coordinator or worker")
)

func Init() error {
	flag.Parse()
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	currentUser, err := user.Current()
	if err != nil {
		return err
	}
	username := currentUser.Username
	if err = os.Setenv("MODE", *mode); err != nil {
		return err
	}
	if err = os.Setenv("OS", runtime.GOOS); err != nil {
		return err
	}
	if err = os.Setenv("USERNAME", username); err != nil {
		return err
	}
	if err = os.Setenv("HOSTNAME", hostname); err != nil {
		return err
	}
	return nil
}

func main() {
	err := Init()
	if err != nil {
		log.Fatal(err)
	}
	if strings.ToLower(*mode) == "worker" {
		bigBrother.StartWorker()
	} else {
		bigBrother.StartCoordinator()
	}
}
