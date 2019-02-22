package board

import (
	"log"
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/board/auction"
	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Action interface
func (b *Board) Action(request, origin string) {
	data := strings.Split(string(request), "#")
	switch data[0] {
	case "Join":
		b.Join(data[1], origin)
	case "Auction":
		b.RaiseAuction(data[1], origin)
	case "Companion":
		b.Nominate(data[1], data[2], origin)
	case "Card":
		b.Play(data[1], data[2], origin)
	}
}

const minScore = 61
const maxScore = 120

// RaiseAuction func
func (b *Board) RaiseAuction(score, origin string) {
	prevScore := int(b.AuctionScore())
	currentScore, _ := strconv.Atoi(score)
	currentScore = auction.Compose(currentScore, auction.NewAuction(prevScore, auction.LT), auction.NewAuction(minScore, auction.LT), auction.NewAuction(maxScore, auction.GT))
	b.SetAuctionScore(uint8(currentScore))
	currentScore = auction.Compose(currentScore, auction.NewAuctionWithReturnScore(prevScore, 0, auction.LT))
	isInfoPresent := func(p *player.Player) bool { return p.Host() == origin }
	p, _ := b.Players().Find(isInfoPresent)
	p.SetAuctionScore(uint8(currentScore))
}

// Play func
func (b *Board) Play(number, seed, origin string) {
	isInfoPresent := func(p *player.Player) bool { return p.Host() == origin }
	p, _ := b.Players().Find(isInfoPresent)
	c, _ := p.Play(number, seed)
	b.PlayedCards().Add(c)
	if len(*b.PlayedCards()) >= 5 {
		playerIndex := briscola.IndexOfWinningCard(*b.PlayedCards(), card.Coin)
		b.PlayedCards().Move(b.Players()[playerIndex].Pile())
	}
}

// Nominate func
func (b *Board) Nominate(number, seed, origin string) (card.ID, error) {
	card, err := card.Create(number, seed)
	isInfoPresent := func(p *player.Player) bool { return p.Has(card) }
	p, err := b.Players().Find(isInfoPresent)
	if err == nil {
		b.selectedCard = card
		b.selectedPlayer = *p
	}
	return card, err
}

// Join func
func (b *Board) Join(name, origin string) {
	isInfoPresent := func(p *player.Player) bool { return p.Name() == "" }
	p, err := b.Players().Find(isInfoPresent)
	if err == nil {
		p.SetName(name)
		p.MyHostIs(origin)
	} else {
		log.Println("All players have joined, no further players are expected: " + err.Error())
	}
}
