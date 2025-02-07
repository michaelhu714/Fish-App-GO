package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type Client struct {
	conn *websocket.Conn
	room string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var conns map[*websocket.Conn]*Client

func initServer() {
	conns = make(map[*websocket.Conn]*Client)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error ugprading connection", err)
		return
	}
	curClient := &Client{conn: conn, room: ""}
	conns[conn] = curClient
	fmt.Println("Client Connected", conn.RemoteAddr())
	go readLoop(conn)
}

func readLoop(conn *websocket.Conn) {
	defer conn.Close()
	for {
		var recievedMsg Message
		err := conn.ReadJSON(&recievedMsg)
		if err != nil {
			fmt.Println("Error read:", err)
			delete(conns, conn)
			break
		}
		fmt.Printf("Recieved: %s\n", recievedMsg)
		switch recievedMsg.Type {
		case "JOIN":
			handleJoin(recievedMsg.Content, conn)
		case "CHAT":
			handleChat(recievedMsg, conn)
		case "LEAVE":
			handleLeave(conn)
		default:
			fmt.Println("?????")
		}
	}
}

func handleJoin(content string, conn *websocket.Conn) {
	client := conns[conn]
	client.room = content
	fmt.Printf("client: %s, room: %s\n", client.conn.RemoteAddr(), client.room)
}

func handleChat(msg Message, conn *websocket.Conn) {
	newContent := fmt.Sprintf("%s: %s", conn.RemoteAddr(), msg.Content)
	newMsg := Message{Type: msg.Type, Content: newContent}
	for c := range conns {
		if conns[c].room == conns[conn].room {
			err := c.WriteJSON(newMsg)
			if err != nil {
				fmt.Println("Error write:", err)
				delete(conns, conn)
				break
			}
		}
	}
}

func handleLeave(conn *websocket.Conn) {
	client := conns[conn]
	client.room = ""
}

func main() {
	initServer()
	http.HandleFunc("/ws", handleConnection)
	port := "8080"
	fmt.Println("WebSocket Server listening on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server failed to start", err)
	}
}
