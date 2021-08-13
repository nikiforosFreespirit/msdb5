package briscola

import (
	"fmt"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/register"
	"github.com/mcaci/msdb5/v2/dom/briscola"
)

// Game struct
type Game struct {
	opts         *Options
	players      misc.Players
	briscolaCard briscola.Card
	board        *briscola.PlayedCards
	registration func(string) error
	deck         *Deck
}

type Options struct {
	WithName string
}

var WithDefaultOptions = &Options{}

func NewGame(gOpts *Options) *Game {
	p, rf := register.NewWithRegistrator(2)
	g := Game{
		opts:         gOpts,
		players:      *p,
		deck:         NewDeck(),
		board:        briscola.NewPlayedCards(2),
		registration: rf,
	}
	return &g
}

// New func
func New() *Game { return &Game{} }

func (g *Game) Players() *misc.Players       { return &g.players }
func (g *Game) Deck() *Deck                  { return g.deck }
func (g *Game) Board() *briscola.PlayedCards { return g.board }
func (g *Game) BoardCards() *set.Cards       { return g.board.Cards }
func (g *Game) Briscola() *briscola.Card     { return &g.briscolaCard }
func (g *Game) Created(name string) bool     { return g.opts != nil && name == g.opts.WithName }
func Register(name string, g *Game) error    { return g.registration(name) }
func Set(card briscola.Card, g *Game)        { g.briscolaCard = card }

func Start(g *Game) {
	Distribute(&struct {
		Players  misc.Players
		Deck     *Deck
		HandSize int
	}{
		Players:  g.players,
		Deck:     g.deck,
		HandSize: 3,
	})
	Set(briscola.Card{Item: g.deck.Top()}, g)
}

func (g Game) String() string {
	return fmt.Sprintf("(Players: %v, Board: %v, Briscola: %v, Deck: %v)", g.players, g.board, g.briscolaCard, g.deck)
}
