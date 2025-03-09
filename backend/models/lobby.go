package models

import "sync"

type Lobby struct {
	ID         int
	Name       string
	Owner      *Player
	NumPlayers int
	Players    map[int]*Player
	Mutex      sync.Mutex
}

type CreateLobbyReq struct {
	Name       string `json:"name" validate:"required"`
	NumPlayers int    `json:"num_players"` // default is 6
	OwnerName  string
}

type CreateLobbyResp struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Owner Player `json:"owner"`
}

type JoinLobbyReq struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type JoinLobbyResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LobbyManager struct {
	Lobbies map[int]*Lobby
	Mutex   sync.Mutex
}

var LobbyMgr = LobbyManager{Lobbies: make(map[int]*Lobby)}
