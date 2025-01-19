package api

import (
	"encoding/json"
	"fmt"
	"github.com/michaelhu714/Fish-App-GO/types"
	"io"
	"net/http"
)

func createPlayerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Printf("Invalid request\n")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var player types.Player
	err = json.Unmarshal(body, &player)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	response := fmt.Sprintf("recieved: name: %s", player.Name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":` + response + `"}`))
}
