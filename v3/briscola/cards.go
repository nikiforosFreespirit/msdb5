package briscola

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

// Card represents the Briscola Card
type Card struct{ card.Item }

type PlayedCards struct {
	*set.Cards
	nPlayers int
}

func NewPlayedCards(nPlayers int) *PlayedCards {
	return &PlayedCards{
		Cards:    &set.Cards{},
		nPlayers: nPlayers,
	}
}

func (c PlayedCards) Pile() *set.Cards {
	if len(*c.Cards) == c.nPlayers {
		return (*set.Cards)(c.Cards)
	}
	return &set.Cards{}
}

type Deck struct{ set.Cards }

func NewDeck() *Deck {
	return &Deck{Cards: set.Deck()}
}
