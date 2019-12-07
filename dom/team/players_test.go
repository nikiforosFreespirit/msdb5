package team

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
)

var testPlayers Players

func init() {
	var a player.Player
	a.RegisterAs("A")
	testPlayers.Add(&a)
	var b player.Player
	b.RegisterAs("B")
	b.Hand().Add(*card.MustID(33))
	testPlayers.Add(&b)
}

func TestSuccessfulFindDataCorresponds(t *testing.T) {
	isPlayerACheck := func(p *player.Player) bool { return p.Name() == "A" }
	if p := testPlayers.At(testPlayers.MustFind(isPlayerACheck)); !isPlayerACheck(p) {
		t.Fatalf("%s and %v are expected to be the same player", "A", p)
	}
}
