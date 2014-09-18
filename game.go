package robo

type Phase int

// These are just straight up wrong, fix them
const (
	Deal Phase = iota
	Move1
	Move2
	Fix
)

type Game struct {
	GameBoard Board
	GameDeck  Deck
	Players   []Player
	GamePhase Phase
}

func newGame(numPlayers int) *Game {
	g := &Game{
		GameBoard: Board{},
		GameDeck:  newDeck(),
		Players:   []Player{},
		GamePhase: Deal,
	}
	return g
}
