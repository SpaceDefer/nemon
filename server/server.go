package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func sendMessage(conn *websocket.Conn) {
	var err error
	msg := `Hi, the handshake is complete!`

	if err = conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		log.Println(err)
		return
	} else {
		fmt.Println("Message sent")
	}

}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))
		sendMessage(conn)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("client connected")
	reader(ws)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Simple Server")
	if err != nil {
		return
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", serveWs)
}

func StartServer() {
	setupRoutes()
	go func() {
		err := http.ListenAndServe(":4000", nil)
		if err != nil {
			return
		}
	}()
	log.Println("ws server exited")
}
