package auction

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/mcaci/msdb5/v2/dom/auction"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Run(players team.Players, listenFor func(context.Context, func())) struct {
	Score  auction.Score
	Caller *player.Player
} {
	ctx, cancel := context.WithCancel(context.Background())
	numbers := make(chan int)
	done := make(chan struct{})
	go listenFor(ctx, func() { numbers <- 60 + rand.Intn(60) })
	go func() {
		<-done
		cancel()
		close(numbers)
	}()

	var curr auction.Score
	var currID uint8

	for score := range numbers {
		pl := players[currID]
		next := auction.Score(score)
		curr = auction.Max120(curr, next)

		switch {
		// Player has folded already
		case player.Folded(pl):
			break
		// Player for scoring less than current
		case auction.ScoreCmp(curr, next) >= 0:
			pl.Fold()
		// Fold everyone if score is 120 or more
		case auction.ScoreCmp(120, next) >= 0:
			for _, p := range players {
				if p == pl {
					continue
				}
				p.Fold()
			}
		}

		// End the loop if only one not folded players is left
		if team.Count(players, notFolded) != 1 {
			// else search next player
			id, err := rotateOn(players, currID, notFolded)
			if err != nil {
				log.Fatalf("error found: %v. Exiting.", err)
			}
			currID = id
		}

		// next phase
		done <- struct{}{}
		close(done)
	}
	return struct {
		Score  auction.Score
		Caller *player.Player
	}{
		Score:  curr,
		Caller: players[players.MustIndex(notFolded)],
	}
}

func notFolded(p *player.Player) bool { return !player.Folded(p) }

func rotateOn(players team.Players, idx uint8, appliesTo player.Predicate) (uint8, error) {
	for i := 0; i < 2*len(players); i++ {
		idx = (idx + 1) % uint8(len(players))
		if !appliesTo(players[idx]) {
			continue
		}
		return idx, nil
	}
	return 0, fmt.Errorf("rotated twice on the number of players and no player found in play.")
}
