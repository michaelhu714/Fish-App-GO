package api

import (
	"encoding/json"
	"fmt"
	"github.com/michaelhu714/Fish-App-GO/internal/fish"
	"github.com/michaelhu714/Fish-App-GO/types"
	"net/http"
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
