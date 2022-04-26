package nemon

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WebsocketServer stores a websocket.Conn and a sync.Mutex
type WebsocketServer struct {
	conn *websocket.Conn // conn is the websocket connection to the client
	mu   sync.Mutex      // mu is a sync.Mutex to prevent data races
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Status string

const (
	Offline      Status = "offline"
	Online       Status = "online"
	Reconnecting Status = "reconnecting"
)

func (ws *WebsocketServer) sendWorkerStatus(ip string, status Status) {
	var err error
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if ws.conn == nil {
		Debug(dInfo, "no client to send data to\n")
		return
	}

	res, err := json.Marshal(&WorkerStatusMessage{Type: Alert, Status: status, WorkerIp: ip})
	if err != nil {
		return
	}
	if err = ws.conn.WriteMessage(websocket.TextMessage, res); err != nil {
		log.Println(err)
		return
	}
}

// sendAlert to the WebsocketServer client
func (ws *WebsocketServer) sendAlert(msg string, ip string) {
	var err error

	ws.mu.Lock()
	defer ws.mu.Unlock()

	if ws.conn == nil {
		Debug(dInfo, "no client to send data to\n")
		return
	}
	res, err := json.Marshal(&AlertMessage{Type: Alert, Message: msg, WorkerIp: ip})
	if err != nil {
		return
	}
	if err = ws.conn.WriteMessage(websocket.TextMessage, res); err != nil {
		log.Println(err)
		return
	}
}

// Cleanup and close WebsocketServer connections
func (ws *WebsocketServer) Cleanup() {
	ws.mu.Lock()
	ws.mu.Unlock()
	Debug(dInfo, "websocket server exiting gracefully...\n")
	if ws.conn == nil {
		return
	}
	err := ws.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("Error during closing websocket:", err)
		return
	}
}

// sendAppList to the WebsocketServer client
func (ws *WebsocketServer) sendAppList(workerInfo *WorkerInfo) {
	var err error

	ws.mu.Lock()
	defer ws.mu.Unlock()
	conn := ws.conn

	if conn == nil {
		Debug(dInfo, "no client found\n")
		return
	}

	reply, err := json.Marshal(workerInfo)
	if err != nil {
		return
	}

	if err := ws.conn.WriteMessage(websocket.TextMessage, reply); err != nil {
		log.Println(err)
		return
	}

}

// reader is the default handler of incoming requests
func (ws *WebsocketServer) reader() {
	for {
		var req DeleteApplicationRequest
		err := ws.conn.ReadJSON(&req)
		if err != nil {
			log.Println(err)
			ws.mu.Lock()
			err := ws.conn.Close()
			if err != nil {
				ws.mu.Unlock()
				log.Println(err)
				return
			}
			ws.mu.Unlock()
			break
		}

		switch req.Type {
		case Delete:
			Debug(dInfo, "app name: %v\ntarget ip: %v\n", req.ApplicationName, req.WorkerIp)

			if !systemInfo.Dev {
				deleteChan <- req
			}
			reply, err := json.Marshal(&DeleteApplicationReply{
				Type:    Acknowledge,
				Ok:      true,
				Message: fmt.Sprintf("removed %v successfully!", req.ApplicationName),
			})
			if err != nil {
				return
			}

			if err := ws.conn.WriteMessage(websocket.TextMessage, reply); err != nil {
				log.Println(err)
				return
			}
		default:
			Debug(dInfo, "don't recognise this currently\n")
		}
	}
}

func (ws *WebsocketServer) serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("client connected")
	ws.mu.Lock()
	ws.conn = conn
	ws.mu.Unlock()
	ws.reader()
}

func (ws *WebsocketServer) homePage(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Simple Server")
	if err != nil {
		return
	}
}

// setupRoutes for nemon server
func (ws *WebsocketServer) setupRoutes() {
	http.HandleFunc("/", ws.homePage)
	http.HandleFunc("/ws", ws.serveWs)
}

// StartServer starts a websocket server
func (ws *WebsocketServer) StartServer() {
	ws.setupRoutes()
	go func() {
		err := http.ListenAndServe(":4000", nil)
		if err != nil {
			return
		}
		log.Println("ws server exited")
	}()
}
