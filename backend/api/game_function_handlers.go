package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/michaelhu714/Fish-App-GO/internal/fish"
	"github.com/michaelhu714/Fish-App-GO/types"
)

func PickCardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Printf("Invalid request\n")
		return
	}
	defer r.Body.Close()
	var pcr types.PickCardReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pcr)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	p1, err := fish.GetPlayer(pcr.P1Name)
	p2, err := fish.GetPlayer(pcr.P1Name)
	if err != nil {
		http.Error(w, "Player doesn't exist", http.StatusBadRequest)
	}
	fish.PickCard(p1, p2, pcr.Card)
	response := fmt.Sprintf("player %s took card {%d, %s} from player %s", pcr.P1Name, pcr.Card.Value, pcr.Card.Suit, pcr.P2Name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":` + response + `"}`))
}

func DeclareHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Printf("Invalid request\n")
		return
	}
	defer r.Body.Close()

	var req types.DeclareReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	// fetch the current player and game state
	game := fish.NewGame()
	if game == nil {
		http.Error(w, "Game state not found", http.StatusInternalServerError)
		return
	}

	// find the current player
	var currentPlayer *fish.Players
	for i := range game.GameState {
		if game.GameState[i].Name == req.CurrentPlayer {
			currentPlayer = &game.GameState[i]
			break
		}
	}

	if currentPlayer == nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	// Call Declare and handle errors
	if err := fish.Declare(currentPlayer, req.Set, game); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("Declaration processed. Current Score -> Team 1: %d, Team 2: %d", game.Team1Points, game.Team2Points)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"` + response + `"}`))
}
