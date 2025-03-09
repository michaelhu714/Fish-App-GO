package services

import (
	"errors"

	"github.com/gorilla/websocket"
	"github.com/michaelhu714/Fish-App-GO/models"
)

var playerId = 1

func CreatePlayer(name string, conn *websocket.Conn) (*models.Player, error) {
	if name == "" {
		return nil, errors.New("player name is required")
	}

	newPlayer := &models.Player{
		ID:   playerId,
		Name: name,
	}

	models.PlayerMgr.Mutex.Lock()
	models.PlayerMgr.Players[conn] = newPlayer
	models.PlayerMgr.Mutex.Unlock()
	return newPlayer, nil
}
