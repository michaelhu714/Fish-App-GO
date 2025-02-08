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

type DeclareReq struct {
	CurrentPlayer string
	Set           string
}

type Players struct {
	name  string
	team  string
	cards []string // cards in players hand

}

// struct to represent game state
type Game1 struct {
	GameState   []Players // slice of Players structs ^ in the game
	Team1Points int
	Team2Points int
}
