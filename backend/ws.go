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
	readLoop(conn)
}

func readLoop(conn *websocket.Conn) {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error read:", err)
			continue
		}
		fmt.Printf("Recieved: %s\n", msg)
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Println("Error write:", err)
			continue
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
