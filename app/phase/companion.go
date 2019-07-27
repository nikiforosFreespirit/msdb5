package phase

import (
	"errors"

	"github.com/mcaci/msdb5/dom/player"
)

// ErrCardNotInHand error
var ErrCardNotInHand = errors.New("Card not in hand")

type companioner interface {
	Find(player.Predicate) (int, *player.Player)
}

func CardAction(rq cardProvider, comp companioner) Data {
	c, err := rq.Card()
	idx, _ := comp.Find(player.IsCardInHand(c))
	var errCardNotInHand error
	if idx < 0 {
		errCardNotInHand = ErrCardNotInHand
	}
	return Data{card: c, plIdx: uint8(idx), cardNotFound: err, cardNotInHand: errCardNotInHand}
}
