package robo

import (
	"math/rand"
)

type CardType int

const CardCount = 84

const (
	UTurn CardType = iota
	RotateLeft
	RotateRight
	Backup
	MoveOne
	MoveTwo
	MoveThree
)

type Card struct {
	Priority int
	Type     CardType
}

type Deck struct {
	Cards     []Card
	Discards  []Card
	drawOrder []int
	nextCard  int
}

func newDeck() *Deck {
	d := &Deck{
		Cards:     makeCards(),
		drawOrder: rand.Perm(CardCount),
		nextCard:  0,
	}
	return d
}

func makeCards() []Card {
	cards := make([]Card, CardCount)

	for i := 1; i <= 6; i++ {
		cards[i] = Card{
			Priority: i * 10,
			Type:     UTurn,
		}
	}

	for i := 7; i <= 41; i += 2 {
		cards[i] = Card{
			Priority: i * 10,
			Type:     RotateLeft,
		}
	}

	for i := 8; i <= 42; i += 2 {
		cards[i] = Card{
			Priority: i * 10,
			Type:     RotateRight,
		}
	}

	for i := 43; i <= 48; i++ {
		cards[i] = Card{
			Priority: i * 10,
			Type:     Backup,
		}
	}

	for i := 49; i <= 66; i++ {
		cards[i] = Card{
			Priority: i * 10,
			Type:     MoveOne,
		}
	}

	for i := 67; i <= 78; i++ {
		cards[i] = Card{
			Priority: i * 10,
			Type:     MoveTwo,
		}
	}

	for i := 79; i <= 84; i++ {
		cards[i] = Card{
			Priority: i * 10,
			Type:     MoveThree,
		}
	}
	return cards
}

func (d *Deck) DrawCard() Card {
	card := d.Cards[d.drawOrder[d.nextCard]]
	d.Discards = append(d.Discards, card)
	return card
}
