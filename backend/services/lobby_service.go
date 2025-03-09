package services

import (
	"errors"

	"github.com/michaelhu714/Fish-App-GO/models"
)

var lobbyIds = 1

func CreateLobby(name string, owner *models.Player) (*models.Lobby, error) {
	if name == "" {
		return nil, errors.New("Error")
	}

	newLobby := &models.Lobby{
		ID:      lobbyIds,
		Name:    name,
		Owner:   owner,
		Players: make(map[int]*models.Player),
	}

	newLobby.Players[owner.ID] = owner
	models.LobbyMgr.Mutex.Lock()
	models.LobbyMgr.LobbiesById[newLobby.ID] = newLobby
	models.LobbyMgr.LobbiesByName[newLobby.Name] = newLobby.ID
	models.LobbyMgr.Mutex.Unlock()
	lobbyIds++
	return newLobby, nil
}

func JoinLobby(id int, player *models.Player) error {
	lobby, exists := models.LobbyMgr.LobbiesById[id]
	if !exists {
		return errors.New("Lobby does not exist")
	}
	lobby.Players[player.ID] = player
	return nil
}
