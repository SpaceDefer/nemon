package main

import (
	"flag"
	"log"
	"os"
	"os/user"
	"runtime"
	"strings"

	"nemon/nemon"
)

var (
	key  = flag.String("key", "", "worker key, same as the coordinator") // key is the nemon product key
	mode = flag.String("mode", "worker", "coordinator or worker")        // mode is the mode nemon runs in
	dev  = flag.Bool("dev", false, "start dev")                          // dev decides whether to start dev environment or not
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
	if err = os.Setenv("NEMONKEY", *key); err != nil {
		return err
	}
	if *dev {
		if err = os.Setenv("DEV", "dev"); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := Init()
	if err != nil {
		log.Fatal(err)
	}
	if strings.ToLower(*mode) == "worker" {
		nemon.StartWorker()
	} else {
		nemon.StartCoordinator()
	}
}
