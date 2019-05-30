package game

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
)

func TestCompletedGameReturningScoreInfo(t *testing.T) {
	gameTest := NewGame(false)
	for i, pl := range gameTest.players {
		pl.Hand().Clear()
		pl.Hand().Add(card.ID(2*i + 5))
		if i > 0 {
			pl.Fold()
		}
	}
	gameTest.Process("Join#A", "127.0.0.51")
	gameTest.Process("Join#B", "127.0.0.52")
	gameTest.Process("Join#C", "127.0.0.53")
	gameTest.Process("Join#D", "127.0.0.54")
	gameTest.Process("Join#E", "127.0.0.55")
	gameTest.Process("Auction#80", "127.0.0.51")
	gameTest.Process("Companion#7#Coin", "127.0.0.51")
	gameTest.Process("Card#5#Coin", "127.0.0.51")
	gameTest.Process("Card#7#Coin", "127.0.0.52")
	gameTest.Process("Card#9#Coin", "127.0.0.53")
	gameTest.Process("Card#1#Cup", "127.0.0.54")
	info := gameTest.Process("Card#3#Cup", "127.0.0.55")
	if info[0].Dest() == "" {
		t.Log(info[0].Err())
		t.Fatal("Expecting transition to end game and scoring")
	}
}
