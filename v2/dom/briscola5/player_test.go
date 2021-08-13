package briscola5

import (
	"testing"
)

func testP(name string) Player {
	return *NewB5Player(name)
}

func TestJoinPlayerName(t *testing.T) {
	if p := testP("Me"); p.Name() != "Me" {
		t.Fatal("Unexpected name")
	}
}

func TestJoinPlayerPileIsEmpty(t *testing.T) {
	if p := testP(""); len(*p.Pile()) != 0 {
		t.Fatal("Pile should be empty")
	}
}
