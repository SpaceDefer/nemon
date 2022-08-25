package nemon

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	pb "nemon/protos"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"
)

// Coordinator struct implements the coordinator
type Coordinator struct {
	workers  map[string]*Worker // workers map ip addresses to Worker structs
	nWorkers uint               // nWorkers gives the number of workers
	allowed  map[string]bool    // list of all allowed applications
	mu       sync.Mutex         // mu mutex to prevent data races in the Coordinator's data
	//screenMu    sync.Mutex         // screenMu to print on the screen exclusively
	discoveryMu sync.Mutex      // discoveryMu to exclusively discover or send heartbeats
	pending     map[string]uint // pending checks if a request to a Worker is pending
	stopCh      chan bool       // stopCh blocking receive to stop the main process
}

// deleteChan sends DeleteApplicationRequest's from the wsServer to ListenDeleteApplication goroutine
var deleteChan chan DeleteApplicationRequest

var notificationChan chan NotifyRequest

// wsServer is an instance of WebsocketServer, started when the coordinator starts up
var wsServer *WebsocketServer

func (c *Coordinator) ListenNotification() {
	for {
		var req NotifyRequest
		req, ok := <-notificationChan

		if !ok {
			fmt.Println("nch error")
			continue
		}
		go func(request *NotifyRequest) {
			Debug(dInfo, "notifying %v\n", request)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			worker := c.workers[request.WorkerIp]
			if worker == nil {
				Debug(dInfo, "didn't find that specific ip\n")
				return
			}
			Debug(dInfo, "worker found %v\n", worker)
			_, err := worker.client.Notify(
				ctx,
				&pb.NotifyRequest{
					Message: encrypt([]byte("come here")),
				},
			)
			st, ok := status.FromError(err)
			checkCodeParseOk(ok)

			if err != nil {
				// TODO: handle various error codes
				fmt.Println(st.Code())
				return
			}

		}(&req)
	}
}

// ListenDeleteApplication wrapper for the server to call
func (c *Coordinator) ListenDeleteApplication() {
	for {
		var req DeleteApplicationRequest
		req, ok := <-deleteChan
		if !ok {
			fmt.Println("ch error")
			continue
		}
		go func(request *DeleteApplicationRequest) {
			Debug(dInfo, "received an rpc, going to delete %v on %v at %v\n", req.ApplicationName, req.WorkerIp, req.Location)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			worker := c.workers[req.WorkerIp]
			if worker == nil {
				Debug(dInfo, "didn't find that specific ip\n")
				return
			}
			Debug(dInfo, "worker found %v\n", worker)
			response, err := worker.client.DeleteApp(
				ctx,
				&pb.DeleteAppsRequest{
					Name:     encrypt([]byte(req.ApplicationName)),
					Location: encrypt([]byte(req.Location)),
				},
			)
			st, ok := status.FromError(err)
			checkCodeParseOk(ok)

			if err != nil {
				// TODO: handle various error codes
				fmt.Println(st.Code())
				return
			}
			Debug(dInfo, "%v\n", response.Ok)
		}(&req)
	}
}

// CheckTimeout checks if a Worker hasn't responded to 3 or more consecutive heartbeats
func (c *Coordinator) CheckTimeout(ip string, username string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	pending := c.pending[ip]
	if pending >= 4 {
		// issue an alert
		Debug(dInfo, "%v's computer hasn't reponsed in ages\n", ip)
		if c.workers[ip] != nil && c.workers[ip].status == Reconnecting {
			return
		}
		c.workers[ip].status = Offline
		wsServer.sendWorkerStatus(ip, c.workers[ip].status)

	}
}

// SendHeartbeat to a single Worker
func (c *Coordinator) SendHeartbeat(worker *Worker) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// TODO: check if error is session expired. add error codes too, easier to check
	response, err := worker.client.GetApps(ctx, &pb.GetAppsRequest{})
	st, ok := status.FromError(err)
	checkCodeParseOk(ok)

	ip := worker.ip
	if err != nil {
		Debug(dInfo, "%v\n", st.Message())
		if st.Code() == codes.Unauthenticated {
			// remove worker info currently and start authentication again with handshake
			c.workers[ip].status = Reconnecting
			wsServer.sendWorkerStatus(ip, c.workers[ip].status)
			//delete(c.workers, ip)
			go c.SendDiscoveryPing(ip)

			Debug(dInfo, "restarting auth with ip %v\n", ip)
		}
		return
	}
	c.pending[ip] = 0

	Debug(dInfo, "app list received from worker %v\n", ip)

	var ApplicationList []ApplicationInfo

	for _, app := range response.Applications {
		if !c.allowed[string(decrypt(app.GetName()))] {
			Debug(dInfo,
				"found an app on %v's at ip [%v] which isn't allowed: %v\n",
				string(decrypt(response.Username)),
				ip,
				string(decrypt(app.GetName())),
			)

			app := ApplicationInfo{
				ApplicationName: string(decrypt(app.GetName())),
				Location:        string(decrypt(app.GetLocation())),
			}

			ApplicationList = append(ApplicationList, app)
		}
	}

	wsServer.sendAppList(&WorkerInfo{
		Type:            Info,
		ApplicationList: ApplicationList,
		WorkerIp:        ip,
		Username:        worker.username,
		Hostname:        worker.hostname,
		Os:              worker.os,
	})

	fmt.Println("heartbeat sent")

}

// BroadcastHeartbeats to multiple workers, cycle signifies the heartbeat cycle
func (c *Coordinator) BroadcastHeartbeats(cycle int) {
	fmt.Print("\033[H\033[2J")
	Debug(dInfo, "heartbeat cycle number %v\n", cycle)
	c.mu.Lock()
	defer c.mu.Unlock()
	workers := c.workers

	for _, worker := range workers {
		Debug(dInfo, "coordinator sending a heartbeat to ip %v\n", worker.ip)
		c.pending[worker.ip]++
		go c.SendHeartbeat(worker)
		go c.CheckTimeout(worker.ip, worker.username)
	}
}

// Cleanup closes the connections in the event of a shutdown
func (c *Coordinator) Cleanup() {
	Debug(dInfo, "running cleanup...\n")
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, worker := range c.workers {
		err := worker.connection.Close()
		checkError(err)
	}
}

// StartCoordinator starts up a Coordinator process
func StartCoordinator() {
	InitSystemInfo()
	wsServer = &WebsocketServer{}
	wsServer.StartServer()
	deleteChan = make(chan DeleteApplicationRequest)
	Debug(dInfo, "%v started as a coordinator\n", os.Getpid())
	coordinator := Coordinator{
		workers:  map[string]*Worker{},
		nWorkers: 0,
		allowed:  map[string]bool{},
		pending:  map[string]uint{},
	}

	// Listen for an exit syscall to perform the cleanup and exit
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)
	go func() {
		<-sigCh
		Debug(dInfo, "\ncoordinator exiting gracefully...\n")
		coordinator.Cleanup()
		wsServer.Cleanup()
		os.Exit(1)
	}()

	if systemInfo.Dev {
		Debug(dInfo, "sending discovery pings over ports\n")
		wsServer.sendDiscoveryNotification()
		coordinator.SendDiscoveryPing("localhost")
	} else {
		coordinator.BroadcastDiscoveryPings()
	}

	coordinator.stopCh = make(chan bool, 1)
	coordinator.mu.Lock()
	nWorkers := coordinator.nWorkers
	workers := coordinator.workers
	coordinator.mu.Unlock()
	Debug(dInfo, "number of workers found: %v\nworkers: %v\n", nWorkers, workers)
	go coordinator.ListenDeleteApplication()
	if nWorkers >= 0 {
		go coordinator.HeartbeatRoutine()
		go coordinator.DiscoveryRoutine()
	}
	<-coordinator.stopCh
}

func (c *Coordinator) DiscoveryRoutine() {
	for {
		if systemInfo.Dev {
			time.Sleep(devDiscoveryPeriod)
		} else {
			time.Sleep(discoveryPeriod)
		}
		c.discoveryMu.Lock()
		if systemInfo.Dev {
			Debug(dInfo, "sending discovery pings over ports\n")
			wsServer.sendDiscoveryNotification()
			c.SendDiscoveryPing("localhost")
		} else {
			c.BroadcastDiscoveryPings()
		}
		c.discoveryMu.Unlock()
	}
}

func (c *Coordinator) HeartbeatRoutine() {
	cycle := 0
	for cycle < 1000 {
		c.discoveryMu.Lock()
		c.BroadcastHeartbeats(cycle)
		cycle++
		c.discoveryMu.Unlock()
		if systemInfo.Dev {
			time.Sleep(devHeartbeatInterval)
		} else {
			time.Sleep(heartbeatInterval)
		}
	}
	c.stopCh <- true
}
