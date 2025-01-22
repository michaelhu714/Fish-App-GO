package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/michaelhu714/Fish-App-GO/internal/fish"
	"github.com/michaelhu714/Fish-App-GO/types"
)

func createPlayerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Printf("Invalid request\n")
		return
	}
	defer r.Body.Close()
	var pr types.CreatePlayerReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pr)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	fish.CreatePlayer(pr.Name)
	response := fmt.Sprintf("recieved: name: %s", pr.Name)
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
	var atr types.AssignTeamReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&atr)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	p := fish.GetPlayer(atr.Name)
	fish.AssignTeam(p, atr.Team)
	response := fmt.Sprintf("recieved: name: %s", pn)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":` + response + `"}`))
}
