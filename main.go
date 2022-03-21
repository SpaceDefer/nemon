package main

import (
	"flag"
	"strings"

	"scan_and_del/scanAndDel"
)

func main() {
	mode := flag.String("mode", "worker", "coordinator or worker")
	flag.Parse()
	if strings.ToLower(*mode) == "worker" {
		scanAndDel.StartWorker()
	} else {
		scanAndDel.StartCoordinator()
	}
}
