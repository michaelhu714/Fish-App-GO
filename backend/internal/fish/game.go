package fish

import (
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

func AssignTeam(p *types.Player, tn int) {
	p.Team = tn
}
