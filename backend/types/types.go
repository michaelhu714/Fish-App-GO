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
	Name  string `json:"name"`
	Team  int    `json:"team,omitempty"`
	Cards []Card `json:"cards,omitempty"`
}

type Game struct {
	Players      []*Player
	teamOneScore int
	teamTwoScore int
}
