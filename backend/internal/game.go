package fish

import (
	"github.com/michaelhu714/Fish-App-GO/types"
)

func createPlayer(name string) {
	np := types.Player{Name: name, Team: -1, Cards: make([]types.Card, 0)}
	g.Players = append(g.Players, &np) // g is curr game
}

func assignTeam(p *types.Player, tn int) {
	p.Team = tn
}
