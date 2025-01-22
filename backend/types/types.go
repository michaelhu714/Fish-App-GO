package types

type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

type Card struct {
	Suit  Suit
	Value int
}

type Player struct {
	Name  string
	Team  int
	Cards []Card
}

type Game struct {
	Players      []*Player
	teamOneScore int
	teamTwoScore int
}

type CreatePlayerReq struct {
	Name string
}

type AssignTeamReq struct {
	Name string
	Team int
}
