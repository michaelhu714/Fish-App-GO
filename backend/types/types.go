package types

type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

type Card struct {
	suit  Suit
	value int
}

type Player struct {
	name  string
	team  int
	cards []Card
}
