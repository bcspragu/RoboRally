package robo

import (
	"math/rand"
)

type CardType int

const CardCount = 84

// All the possible move card types
const (
	UTurn CardType = iota
	RotateLeft
	RotateRight
	Backup
	MoveOne
	MoveTwo
	MoveThree
)

// ProgramCard represents a movement card
type ProgramCard struct {
	Priority int
	Type     CardType
}

type Deck struct {
	Cards     []ProgramCard
	drawOrder []int
	nextCard  int
}

func newDeck() Deck {
	d := Deck{
		Cards:     makeCards(),
		drawOrder: rand.Perm(CardCount),
		nextCard:  0,
	}
	return d
}

// makeCards generates a slice of cards with the priorities and types set by the game.
func makeCards() []ProgramCard {
	cards := make([]ProgramCard, CardCount)

	for i := 1; i <= 6; i++ {
		cards[i] = ProgramCard{
			Priority: i * 10,
			Type:     UTurn,
		}
	}

	for i := 7; i <= 41; i += 2 {
		cards[i] = ProgramCard{
			Priority: i * 10,
			Type:     RotateLeft,
		}
	}

	for i := 8; i <= 42; i += 2 {
		cards[i] = ProgramCard{
			Priority: i * 10,
			Type:     RotateRight,
		}
	}

	for i := 43; i <= 48; i++ {
		cards[i] = ProgramCard{
			Priority: i * 10,
			Type:     Backup,
		}
	}

	for i := 49; i <= 66; i++ {
		cards[i] = ProgramCard{
			Priority: i * 10,
			Type:     MoveOne,
		}
	}

	for i := 67; i <= 78; i++ {
		cards[i] = ProgramCard{
			Priority: i * 10,
			Type:     MoveTwo,
		}
	}

	for i := 79; i <= 84; i++ {
		cards[i] = ProgramCard{
			Priority: i * 10,
			Type:     MoveThree,
		}
	}
	return cards
}

// DrawCards pulls numCards from the deck, shuffling if there aren't enough left in the deck to grab numCards cards.
func (d *Deck) DrawCards(numCards int) []ProgramCard {
	cards := make([]ProgramCard, numCards)

	if numCards > CardCount-d.nextCard {
		d.shuffleDeck()
	}

	for i := 0; i < numCards; i++ {
		cards[i] = d.drawCard()
	}

	return cards
}

// drawCard pulls a single card from the deck.
func (d *Deck) drawCard() ProgramCard {
	d.nextCard++
	return d.Cards[d.drawOrder[d.nextCard-1]]
}

// shuffleDeck generates a new card order and moves our index to the top of the deck.
func (d *Deck) shuffleDeck() {
	d.drawOrder = rand.Perm(CardCount)
	d.nextCard = 0
}
