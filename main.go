package main

import (
	"flag"
	"os"
	"runtime"
	"strings"

	"big_brother/bigBrother"
)

func main() {
	mode := flag.String("mode", "worker", "coordinator or worker")
	//port := flag.String("port", "1234", "enter a port")
	flag.Parse()
	err := os.Setenv("MODE", *mode)
	if err != nil {
		return
	}
	err = os.Setenv("OS", runtime.GOOS)
	if err != nil {
		return
	}
	if strings.ToLower(*mode) == "worker" {
		bigBrother.StartWorker()
	} else {
		bigBrother.StartCoordinator()
	}
}
