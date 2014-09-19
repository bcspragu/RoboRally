package robo

import (
	"errors"
)

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

func newGame() *Game {
	g := &Game{
		GameBoard: Board{},
		GameDeck:  newDeck(),
		Players:   []Player{},
		GamePhase: Deal,
	}
	return g
}

func (g *Game) addPlayer(p Player) error {
	if len(g.Players) < 8 {
		g.Players = append(g.Players, p)
	} else {
		return errors.New("Too many players")
	}
	return nil
}
