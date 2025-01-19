package fish

import (
	"fmt"
	"time"

	"math/rand"
)

// struct to represent each player
type Players struct {
	name  string
	team  string
	cards []string // cards in players hand
}

// struct to represent game state
type Game struct {
	GameState []Players // slice of Players structs ^ in the game
	Rng       *rand.Rand
}

func NewGame() *Game {
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
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("Random generator created:", rng)
	rng.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	// create 6 players
	player := []Players{
		{name: "Player 1, ", team: "Team 1"},
		{name: "Player 2, ", team: "Team 1"},
		{name: "Player 3, ", team: "Team 1"},
		{name: "Player 4, ", team: "Team 2"},
		{name: "Player 5, ", team: "Team 2"},
		{name: "Player 6, ", team: "Team 2"},
	}

	// deal 9 cards to each player
	for i := 0; i < 9; i++ {
		for j := range player {
			player[j].cards = append(player[j].cards, deck[0])
			// update deck to get rid of first element
			deck = deck[1:]
		}
	}

	return &Game{GameState: player}
}

// randommizes first player
func RandomizeFirstPlayer(numPlayers int, rng *rand.Rand) int {
	return 1
}
func MakeSets() (lH, lC, lS, lD, hH, hC, hS, hD, eJ []string) {
	ranks := "23456789TJQKA"
	suits := []string{"♥", "♣", "♠", "♦"}

	for _, rank := range ranks {
		for _, suit := range suits {
			card := string(rank) + suit

			// Low cards: 2-7
			if rank >= '2' && rank <= '7' {
				if suit == "♥" {
					lH = append(lH, card)
				} else if suit == "♣" {
					lC = append(lC, card)
				} else if suit == "♠" {
					lS = append(lS, card)
				} else if suit == "♦" {
					lD = append(lD, card)
				}
			}

			// High cards: 9-A
			if rank >= '9' && (rank <= 'K' || rank == 'A') {
				if suit == "♥" {
					hH = append(hH, card)
				} else if suit == "♣" {
					hC = append(hC, card)
				} else if suit == "♠" {
					hS = append(hS, card)
				} else if suit == "♦" {
					hD = append(hD, card)
				}
			}

			// Eights and Jokers
			if rank == '8' {
				eJ = append(eJ, card)
			}
		}
	}

	// Add Jokers to eights and jokers set
	eJ = append(eJ, "red🃏", "black🃏")

	return
}

// inital game setup (will make cleaner later)
func GameInit(player Players) {
	// display first player to go
	fmt.Println(player.name + player.team + " has the first move.")
	fmt.Println("Your Cards: ", player.cards)
	fmt.Print("It's your turn! \nChoose a card: ")
	var card string
	fmt.Scan(&card)
	fmt.Printf("You choose the %s", card)
}
