package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	game := newGame()
	fmt.Println(game)
}

// struct to represent each player
type Player struct {
	name  string
	team  string
	cards []string // cards in players hand
}

// struct to represent game state
type Game struct {
	gameState []Player // slice of Player structs ^ in the game
}

func newGame() *Game {
	deck := []string{}
	ranks := "23456789TJQKA"
	suits := []string{"♠", "♥", "♦", "♣"}

	for _, rank := range ranks {
		for _, suit := range suits {
			deck = append(deck, string(rank)+suit)
		}
	}

	deck = append(deck, "red🃏", "black🃏")
	// figure out a better representation for jokers later

	// shuffle deck for distrbution
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	// create 6 players
	players := []Player{
		{name: "Player 1, ", team: "Team 1"},
		{name: "Player 2, ", team: "Team 1"},
		{name: "Player 3, ", team: "Team 1"},
		{name: "Player 4, ", team: "Team 2"},
		{name: "Player 5, ", team: "Team 2"},
		{name: "Player 6, ", team: "Team 2"},
	}

	// deal 9 cards to each player
	for i := 0; i < 9; i++ {
		for j := range players {
			players[j].cards = append(players[j].cards, deck[0])
			// update deck to get rid of first element
			deck = deck[1:]
		}
	}

	return &Game{gameState: players}
}
