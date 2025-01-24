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
	response := fmt.Sprintf("recieved: name: %s", pr.Name)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":` + response + `"}`))
}
