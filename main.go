package main

import (
	"flag"
	"strings"
	"time"

	"scan_and_del/scanAndDel"
)

func StartWorker() {
	worker := scanAndDel.Worker{}
	go worker.Start()
	time.Sleep(5 * time.Second)
}

func StartCoordinator() {
	coordinator := scanAndDel.Coordinator{}
	go coordinator.Start()
	time.Sleep(5 * time.Second)
}

func main() {
	mode := flag.String("mode", "worker", "coordinator or worker")
	flag.Parse()
	if strings.ToLower(*mode) == "worker" {
		StartWorker()
	} else {
		StartCoordinator()
	}
}
