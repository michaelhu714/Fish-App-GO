package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Player struct {
	ID   int
	Name string
	Conn *websocket.Conn
}

type PlayerManager struct {
	Players map[*websocket.Conn]*Player
	Mutex   sync.Mutex
}

var PlayerMgr PlayerManager = PlayerManager{Players: make(map[*websocket.Conn]*Player)}
