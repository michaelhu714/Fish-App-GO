package api

import (
	"encoding/json"
	"fmt"
	"github.com/michaelhu714/Fish-App-GO/types"
	"net/http"
)

func createPlayerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Printf("Invalid request\n")
		return
	}
	defer r.Body.Close()
	var player types.Player
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&player)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	player.Cards = make([]types.Card, 0)
	response := fmt.Sprintf("recieved: name: %s", player.Name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":` + response + `"}`))
}
