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
	Set   *map[Card]Card
}

type Player struct {
	Name  string
	Team  int
	Cards map[Card]Card
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

type ShuffleTeamReq struct {
	filler  string // assume these are gonna hold smth, idk what they are rn
	filler1 string
}

type PickCardReq struct {
	P1Name string
	P2Name string
	Card   Card
}
