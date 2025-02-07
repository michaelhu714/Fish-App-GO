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

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error ugprading connection", err)
		return
	}
	defer conn.Close()
	fmt.Println("Client Connected", conn.RemoteAddr())
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
