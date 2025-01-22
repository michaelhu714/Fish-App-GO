/*package fish

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
	team  int
	cards []string // cards in players hand
}

// struct to represent game state
type Game struct {
	Players []Player // slice of Player structs ^ in the game
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

	//shuffle deck for distrbution
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	fmt.Println(deck)
	return &Game{
		Players: []Player{}, // Initialize with no players for now

	}
}*/
