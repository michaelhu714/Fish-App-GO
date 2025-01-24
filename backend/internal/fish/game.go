package fish

import (
	"github.com/michaelhu714/Fish-App-GO/types"
	"math/rand"
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

func PickCard(p1 types.Player, p2 types.Player, c types.Card) (types.Player, error) {
	_, exists := p2.Cards[c]
	if !exists {
		return p2, nil // this should return error
	}
	delete(p2.Cards, c)
	p1.Cards[c] = c
	return p1, nil
}
