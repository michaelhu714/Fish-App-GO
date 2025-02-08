package fish

import (
	"math/rand"

	"github.com/michaelhu714/Fish-App-GO/internal/util"
	"github.com/michaelhu714/Fish-App-GO/types"
)

func CreatePlayer(name string) {
	np := types.Player{Name: name, Team: -1, Cards: make([]types.Card, 0)}
	g.Players = append(g.Players, &np) // g is curr game
}

func GetPlayer(name string) (*types.Player, error) {
	for p := range g.Players {
		if p.Name == name {
			return p, nil
		}
	}
	return nil, nil // not found error
}

func AssignTeam(p *types.Player, tn int) error {
	if tn != 1 || tn != 2 {
		return nil // invalid team error
	}
	p.Team = tn
	return nil
}

func ShuffleTeams() {
	countOne := 0
	countZero := 0
	usedSlice := make([]int, len(g.Players))
	for p := range g.Players {
		if countZero == len(g.Players)/2 {
			p.Team = 1
			countOne++
		} else if countOne == len(g.Players)/2+1 {
			p.Team = 0
			countZero++
		} else {
			rand := rand.Int31n(2)
			p.Team = rand
			if rand == 0 {
				countZero++
			} else {
				countOne++
			}
		}
	}
}

func PickCard(p1 *types.Player, p2 *types.Player, c types.Card) (*types.Player, error) {
	if !validatePick(p1, c) {
		return nil, nil // this should return an error
	}
	_, exists := p2.Cards[c]
	if !exists {
		return p2, nil
	}
	delete(p2.Cards, c)
	p1.Cards[c] = c
	return p1, nil
}

func validatePick(p1 *types.Player, c types.Card) bool {
	return len(util.Intersection(p1.Cards, *c.Set)) == 0
}

/*
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
*/
