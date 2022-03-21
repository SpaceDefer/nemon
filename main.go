package main

import (
	"flag"
	"strings"

	"scan_and_del/bigBrother"
)

func main() {
	mode := flag.String("mode", "worker", "coordinator or worker")
	flag.Parse()
	if strings.ToLower(*mode) == "worker" {
		bigBrother.StartWorker()
	} else {
		bigBrother.StartCoordinator()
	}
}
