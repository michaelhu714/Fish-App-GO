package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var conns map[*websocket.Conn]bool = make(map[*websocket.Conn]bool)

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error ugprading connection", err)
		return
	}
	conns[conn] = true
	fmt.Println("Client Connected", conn.RemoteAddr())
	go readLoop(conn)
}

func readLoop(conn *websocket.Conn) {
	defer conn.Close()
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error read:", err)
			delete(conns, conn)
			break
		}
		fmt.Printf("Recieved: %s\n", msg)
		fullMsg := fmt.Sprintf("%s: %s", conn.RemoteAddr(), msg)
		for c := range conns {
			err = c.WriteMessage(msgType, []byte(fullMsg))
		}
		if err != nil {
			fmt.Println("Error write:", err)
			delete(conns, conn)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)
	port := "8080"
	fmt.Println("WebSocket Server listening on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server failed to start", err)
	}
}
