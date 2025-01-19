package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/michaelhu714/Fish-App-GO/api"
	fish "github.com/michaelhu714/Fish-App-GO/internal"
)

const portNum string = ":8000"

func main() {

	http.HandleFunc("/", api.Home)
	http.HandleFunc("/new", api.New)
	fmt.Printf("Server started on port %s\n", portNum)
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}

	game := fish.NewGame()
	fmt.Println(game)
	// generate the sets for use later
	lowHearts, lowClubs, lowSpades, lowDiamonds, highHearts, highClubs, highSpades, highDiamonds, eightsJokers := fish.MakeSets()

	fmt.Println(lowHearts, lowClubs, lowSpades, lowDiamonds, highHearts, highClubs, highSpades, highDiamonds, eightsJokers)

	firstPlayerIndex := fish.RandomizeFirstPlayer(len(game.GameState), game.Rng)
	fish.GameInit(game.GameState[firstPlayerIndex])

}
