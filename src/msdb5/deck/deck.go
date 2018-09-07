package deck

import (
	"math/rand"
	"msdb5/card"
	"time"
)

const Size = 40

// Deck interface
type Deck interface {
	RemoveTop() *card.Card
	IsEmpty() bool
}

// Create func
func New() Deck {
	deck := new(concreteDeck)

	rand.Seed(time.Now().UnixNano())
	deck.cards = rand.Perm(Size)
	return deck
}
