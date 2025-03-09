package websockets

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/michaelhu714/Fish-App-GO/models"
	"github.com/michaelhu714/Fish-App-GO/services"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleLobbySocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("WebSocket upgrade failed\n")
	}
	rName := r.URL.Query().Get("room")
	pName := r.URL.Query().Get("name")
	player, err := services.CreatePlayer(pName, conn)
	if err != nil {
		fmt.Printf("Error in player creation\n")
	}
	room, exists := models.LobbyMgr.LobbiesByName[rName]
	if exists {
		services.JoinLobby(room, player)
	} else {
		fmt.Println("Lobby does not exist")
		conn.Close()
		return
	}
	for {
		var msg map[string]string
		err = conn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("Websocket read err: %s\n", err)
			break
		}
	}
	conn.Close()
}
