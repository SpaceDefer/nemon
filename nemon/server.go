package nemon

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type WebsocketServer struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (ws *WebsocketServer) sendAlert(msg string) {
	var err error

	ws.mu.Lock()
	conn := ws.conn
	ws.mu.Unlock()

	if conn == nil {
		fmt.Printf("no client to send data to\n")
		return
	}
	if err = conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		log.Println(err)
		return
	} else {
		fmt.Println("message sent")
	}

}


func (ws *WebsocketServer) sendAppList(list ApplicationList) {
	var err error

	ws.mu.Lock()
	conn := ws.conn
	ws.mu.Unlock()

	if conn == nil {
		fmt.Printf("no client found\n")
		return
	}

	reply, err := json.Marshal(&list)
	if err != nil {
			return
	}

	if err := ws.conn.WriteMessage(websocket.TextMessage, reply); err != nil {
		log.Println(err)
		return
	}

}


func (ws *WebsocketServer) reader() {
	for {
		var req DeleteApplicationRequest
		err := ws.conn.ReadJSON(&req)
		if err != nil {
			log.Println(err)
			ws.conn.Close()
			break
		}
		fmt.Printf("app name: %v\ntarget ip: %v\n", req.ApplicationName, req.WorkerIp)
		// deleteChan <- req
		// reply, err := json.Marshal(&DeleteApplicationReply{Ok: true})
		// if err != nil {
		// 	return
		// }

		// if err := ws.conn.WriteMessage(websocket.TextMessage, reply); err != nil {
		// 	log.Println(err)
		// 	return
		// }
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
