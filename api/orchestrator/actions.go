package orchestrator

import (
	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/api/action/execute/auction"
	"github.com/nikiforosFreespirit/msdb5/api/action/execute/exchange"
	"github.com/nikiforosFreespirit/msdb5/api/action/execute/join"
	"github.com/nikiforosFreespirit/msdb5/api/action/execute/nominate"
	"github.com/nikiforosFreespirit/msdb5/api/action/execute/play"
	"github.com/nikiforosFreespirit/msdb5/api/action/find"

	"github.com/nikiforosFreespirit/msdb5/player"
)

func NewFinder(requestname, request, origin string, currentPlayer *player.Player) (finder action.Finder) {
	switch requestname {
	case "Join":
		finder = find.NewJoinFinder()
	default:
		finder = find.NewPlayerFinder(origin, currentPlayer)
	}
	return
}

func NewExecuter(requestname, request, origin string, o *Orchestrator) (executer action.Executer) {
	switch requestname {
	case "Join":
		executer = join.NewJoin(request, origin)
	case "Auction":
		executer = auction.NewAuction(request, origin, o.game.Board())
	case "Exchange":
		executer = exchange.NewExchangeCards(request, origin, o.game.Board().SideDeck())
	case "Companion":
		executer = nominate.NewCompanion(request, origin, o.game.Players(), o.game.SetCompanion)
	case "Card":
		executer = play.NewPlay(request, origin, o.game.Players(),
			o.game.Board().PlayedCards(), o.game.Board().SideDeck(), o.game.BriscolaSeed())
	}
	return
}
