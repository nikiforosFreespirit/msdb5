package team

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/player"
)

var testPlayers Players

func init() {
	testPlayers = *New(2)
	a := player.New(&player.Options{For2P: true, Name: "A"})
	a.Hand().Add(*card.MustID(34))
	testPlayers[0] = a
	b := player.New(&player.Options{For2P: true, Name: "B"})
	b.Hand().Add(*card.MustID(33))
	b.Hand().Add(*card.MustID(34))
	testPlayers[1] = b
}

func TestSuccessfulFindDataCorrespondsToA(t *testing.T) {
	i, err := testPlayers.SelectIndex(testPredicateA)
	if err != nil {
		t.Fatal(err)
	}
	if p := testPlayers[i]; !testPredicateA(p) {
		t.Fatalf("%s and %v are expected to be the same player", "A", p)
	}
}

func TestSuccessfulFindDataCorrespondsToB(t *testing.T) {
	testPredicateB := func(p player.Player) bool { return p.Name() == "B" }
	i, err := testPlayers.SelectIndex(testPredicateB)
	if err != nil {
		t.Fatal(err)
	}
	if p := testPlayers[i]; !testPredicateB(p) {
		t.Fatalf("%s and %v are expected to be the same player", "B", p)
	}
}
