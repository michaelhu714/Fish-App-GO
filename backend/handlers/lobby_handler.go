package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/michaelhu714/Fish-App-GO/models"
	"net/http"
)

func CreateLobby(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Printf("Invalid request\n")
		return
	}
	defer r.Body.Close()
	var req models.CreateLobbyReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

}
