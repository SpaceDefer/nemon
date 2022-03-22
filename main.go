package main

import (
	"flag"
	"strings"

	"scan_and_del/bigBrother"
)

func main() {
	mode := flag.String("mode", "worker", "coordinator or worker")
	port := flag.String("port", "1234", "enter a port")
	flag.Parse()
	if strings.ToLower(*mode) == "worker" {
		bigBrother.StartWorker(port)
	} else {
		bigBrother.StartCoordinator()
	}
}
