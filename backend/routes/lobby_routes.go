package routes

import (
	"github.com/gorilla/mux"
	"github.com/michaelhu714/Fish-App-GO/websockets"
)

func HandleCreateLobby() {
}

func HandleJoinLobby() {
}

func RegisterLobbyRoutes(router *mux.Router) {
	router.HandleFunc("/lobby/{lobbyName}/ws", websockets.HandleLobbySocket).Methods("GET")
}
