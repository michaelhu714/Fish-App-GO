package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Room    string `json:"room"`
}

type RoomReq struct {
	Room string `json:"room"`
}

type Room struct {
	Name    string
	Clients map[*websocket.Conn]bool
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var rooms map[string]*Room
var roomsMutex sync.Mutex

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func initServer() {
	rooms = make(map[string]*Room)
	roomsMutex = sync.Mutex{}
}

func handleSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	roomName := r.URL.Query().Get("room")
	fmt.Println("room:", roomName)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error ugprading connection", err)
		return
	}
	fmt.Println("Client Connected", conn.RemoteAddr())
	rooms[roomName].Clients[conn] = true
	go readLoop(conn, rooms[roomName])
}

func readLoop(conn *websocket.Conn, room *Room) {
	defer conn.Close()
	for {
		var recievedMsg Message
		err := conn.ReadJSON(&recievedMsg)
		if err != nil {
			fmt.Println("Error read:", err)
			delete(room.Clients, conn)
			break
		}
		fmt.Printf("Recieved: %s\n", recievedMsg)
		fmt.Printf("msg type: %s, content: %s\n", recievedMsg.Type, recievedMsg.Content)
		switch recievedMsg.Type {
		case "CHAT":
			handleChat(recievedMsg, conn, room)
		case "LEAVE":
			handleLeave(conn, room)
		default:
			fmt.Println("?????")
		}
	}
}

func handleChat(msg Message, conn *websocket.Conn, room *Room) {
	newContent := fmt.Sprintf("%s: %s", conn.RemoteAddr(), msg.Content)
	newMsg := Message{Type: msg.Type, Content: newContent}
	for c := range room.Clients {
		err := c.WriteJSON(newMsg)
		if err != nil {
			fmt.Println("Error write:", err)
			delete(room.Clients, conn)
			break
		}
	}
}

func Join(roomName string) {
	roomsMutex.Lock()
	if _, exists := rooms[roomName]; !exists {
		rooms[roomName] = &Room{Name: roomName, Clients: make(map[*websocket.Conn]bool)}
	}
	roomsMutex.Unlock()
}

func handleJoin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Printf("Invalid request\n")
		return
	}
	defer r.Body.Close()
	var reqBody RoomReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)
	if err != nil || reqBody.Room == "" {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	Join(reqBody.Room)
	response := map[string]string{"room": reqBody.Room}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleLeave(conn *websocket.Conn, room *Room) {
	fmt.Printf("client %s leaving room %s\n", conn.RemoteAddr(), room.Name)
	delete(room.Clients, conn)
}

func main() {
	initServer()
	http.HandleFunc("/ws", handleSocket)
	http.HandleFunc("/api/join", enableCORS(handleJoin))
	port := "8080"
	fmt.Println("WebSocket Server listening on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server failed to start", err)
	}
}
