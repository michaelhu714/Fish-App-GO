package api

import (
	"encoding/json"
	"fmt"
	"github.com/michaelhu714/Fish-App-GO/internal/fish"
	"net/http"
)

func createPlayerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Printf("Invalid request\n")
		return
	}
	defer r.Body.Close()
	var pn string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pn)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	fish.CreatePlayer(pn)
	response := fmt.Sprintf("recieved: name: %s", pn)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":` + response + `"}`))
}

func joinTeamHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Printf("Invalid request\n")
		return
	}
	defer r.Body.Close()
}
