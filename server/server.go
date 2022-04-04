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
	CheckOrigin := func(r *http.Request) bool {
		return true
	}
	upgrader.CheckOrigin = CheckOrigin

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	http.HandleFunc("/ws", serveWs)
}

func StartServer() {
	setupRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
