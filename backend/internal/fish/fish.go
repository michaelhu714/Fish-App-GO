package fish

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// struct to represent each player
type Players struct {
	Name  string
	team  string
	cards []string // cards in players hand

}

// struct to represent game state
type Game struct {
	GameState   []Players // slice of Players structs ^ in the game
	Team1Points int
	Team2Points int
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
	rng.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	// create 6 players
	player := []Players{
		{Name: "Player 1, ", team: "Team 1"},
		{Name: "Player 2, ", team: "Team 1"},
		{Name: "Player 3, ", team: "Team 1"},
		{Name: "Player 4, ", team: "Team 2"},
		{Name: "Player 5, ", team: "Team 2"},
		{Name: "Player 6, ", team: "Team 2"},
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
func RandomizeFirstPlayer(numPlayers int) int {

	// seed the random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// generate a random number between 1 and numPlayers
	return rng.Intn(numPlayers)
}
func MakeSets() (lH, lC, lS, lD, hH, hC, hS, hD, eJ []string) {
	ranks := "23456789TJQKA"
	suits := []string{"♥", "♣", "♠", "♦"}

	for _, rank := range ranks {
		for _, suit := range suits {
			card := string(rank) + suit

			// low cards: 2-7
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

			// high cards: 9-A
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

			// eights and Jokers
			if rank == '8' {
				eJ = append(eJ, card)
			}
		}
	}

	// add Jokers to eights and jokers set
	eJ = append(eJ, "red🃏", "black🃏")

	return
}

// inital game setup (will make cleaner later)
func GameInit(player Players, game *Game) {
	var choice string

	fmt.Println(player.Name + player.team + " has the first move.")
	fmt.Println("Your Cards:", player.cards)

	// loop until a valid choice is made
	for {
		fmt.Print("It's your turn! Type 'D' to declare or 'C' to choose a card: ")
		fmt.Scan(&choice)

		choice = strings.ToUpper(choice) // normalize input to uppercase

		if choice == "D" {
			Declare1(player, game)
			break
		} else if choice == "C" {
			fmt.Println("You chose to pick a card") // replace with actual logic
			break
		} else {
			fmt.Println("Invalid input. Please enter 'D' or 'C'.")
		}
	}
}

func Declare1(currentPlayer Players, game *Game) {
	var declaringSet string
	var team []Players

	// loop until a valid set is provided.
	var setCards []string
	for {
		fmt.Println("\nWhat set are you declaring for?\n(Choose one of: lH, lC, lS, lD, hH, hC, hS, hD, eJ)")
		fmt.Scan(&declaringSet)

		lH, lC, lS, lD, hH, hC, hS, hD, eJ := MakeSets()

		switch declaringSet {
		case "lH":
			setCards = lH
		case "lC":
			setCards = lC
		case "lS":
			setCards = lS
		case "lD":
			setCards = lD
		case "hH":
			setCards = hH
		case "hC":
			setCards = hC
		case "hS":
			setCards = hS
		case "hD":
			setCards = hD
		case "eJ":
			setCards = eJ
		default:
			fmt.Println("Invalid set name. Please try again.")
			continue // prompt again for a valid set name
		}
		// valid set was provided, so break out of the loop.
		break
	}
	team = DefineTeammates(currentPlayer, game)
	promptDeclare(currentPlayer, team, setCards, game)
}

// identifies teammates of current player
func DefineTeammates(currentPlayer Players, game *Game) []Players {
	var team []Players

	for _, p := range game.GameState {
		if p.team == currentPlayer.team {
			team = append(team, p)
		}
	}

	return team
}

// prints all declaration stuff
func promptDeclare(currentPlayer Players, team []Players, set []string, game *Game) {
	var teammateNum int
	var validDeclaration bool
	DisplayCurrentHand(currentPlayer, set)
	for _, card := range set {
		fmt.Printf("\nYou can select from the following teammates (Select teammate 1, 2, 3, or 4 for yourself):\n") // %v", team) if you ever want to display cards
		for i, mate := range team {
			fmt.Printf("%d: %s (%s)\n", i+1, mate.Name, mate.team)
		}
		fmt.Printf("Your card is %s, select which teammate has it: ", card)

		for {
			fmt.Scan(&teammateNum)

			// check if input is valid (
			if teammateNum == 1 || teammateNum == 2 || teammateNum == 3 || teammateNum == 4 {
				break // Valid input, exit loop
			}

			// invalid input, prompt again
			fmt.Printf("Invalid teammate number. Please enter 1, 2, 3, or 4 for yourself: ")
		}

		if teammateNum != 4 {
			if !ValidatePick(team, set, teammateNum) {
				validDeclaration = false
			}
		}
	}
	if validDeclaration {
		fmt.Println("Successful declaration! Your team gets a point.")
		if currentPlayer.team == "Team 1" {
			game.Team1Points++
			fmt.Println("Team 1 gets a point!")
		} else {
			game.Team2Points++
			fmt.Println("Team 2 gets a point!")
		}

	} else {
		fmt.Println("Failed declaration. You lose the point to your opponents.")
		if currentPlayer.team == "Team 1" {
			game.Team2Points++
			fmt.Println("Team 2 gets a point!")
		} else {
			game.Team1Points++
			fmt.Println("Team 1 gets a point!")
		}
	}
	CheckWin(game)

}

// check if the current player has any card from the declared set.
func DisplayCurrentHand(currentPlayer Players, set []string) {
	for _, card := range currentPlayer.cards {
		for _, setCard := range set {
			if card == setCard {
				fmt.Printf("I have %s\n", card)
			}
		}
	}
}

func ValidatePick(team []Players, set []string, teammateNum int) bool {
	target := team[teammateNum-1]

	// Check whether the target's hand contains any card from the declared set.
	found := false
	for _, card := range target.cards {
		for _, setCard := range set {
			if card == setCard {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	return found

}

// check if a team has reached 5 points and display scores
func CheckWin(game *Game) {
	// display current team points.
	fmt.Printf("Current Score -> Team 1: %d, Team 2: %d\n", game.Team1Points, game.Team2Points)

	// check if either team has reached 5 points to win.
	if game.Team1Points >= 5 {
		fmt.Println("Team 1 wins the game!")
	} else if game.Team2Points >= 5 {
		fmt.Println("Team 2 wins the game!")
	}
}

func Declare(player *Players, set string, game *Game) error {
	setCards, err := GetSetCards(set)
	if err != nil {
		return err
	}

	team := GetTeammates(player, game)
	if !ValidateDeclaration(team, setCards) {
		return fmt.Errorf("invalid declaration")
	}

	UpdateScore(player.team, game)
	return nil
}

func GetSetCards(set string) ([]string, error) {
	lH, lC, lS, lD, hH, hC, hS, hD, eJ := MakeSets()
	switch set {
	case "lH":
		return lH, nil
	case "lC":
		return lC, nil
	case "lS":
		return lS, nil
	case "lD":
		return lD, nil
	case "hH":
		return hH, nil
	case "hC":
		return hC, nil
	case "hS":
		return hS, nil
	case "hD":
		return hD, nil
	case "eJ":
		return eJ, nil
	default:
		return nil, fmt.Errorf("invalid set name")
	}
}

func GetTeammates(player *Players, game *Game) []*Players {
	var team []*Players
	for i := range game.GameState {
		if game.GameState[i].team == player.team {
			team = append(team, &game.GameState[i])
		}
	}
	return team
}

func ValidateDeclaration(team []*Players, set []string) bool {
	for _, p := range team {
		for _, card := range p.cards {
			for _, sCard := range set {
				if card == sCard {
					return true
				}
			}
		}
	}
	return false
}

func UpdateScore(team string, game *Game) {
	if team == "Team 1" {
		game.Team1Points++
	} else {
		game.Team2Points++
	}
}
