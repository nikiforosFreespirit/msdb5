package game

import (
	"fmt"
	"os"

	"github.com/mcaci/msdb5/dom/briscola"

	"github.com/mcaci/msdb5/app/action"
	"github.com/mcaci/msdb5/app/msg"
	"github.com/mcaci/msdb5/app/next"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

// Process func
func (g *Game) Process(inputRequest, origin string) []PlMsg {
	printer := message.NewPrinter(g.Lang())
	rq := NewReq(inputRequest, origin)
	pr := proc{}

	// verify phase step
	if pr.err == nil {
		// err = msg.UnexpectedPhaseErr(phase.MustID(rq), g.Phase(), g.Lang())
		pr.reportErr(g, rq, phase.Check(g, rq))
	}

	// verify player step
	if pr.err == nil {
		// err = msg.UnexpectedPlayerErr(g.CurrentPlayer().Name(), g.Lang())
		pr.reportErr(g, rq, team.CheckOrigin(g, rq))
	}

	// play step
	if pr.err == nil {
		pr.reportErr(g, rq, action.Play(g, rq))
	}

	if pr.err == nil {
		cardN := auction.SideCards(*g.AuctionScore())
		if phase.InsideAuction == g.Phase() && len(*g.SideDeck()) != 0 && cardN > 0 {
			for _, pl := range g.Players() {
				pl, plMsg := pl, printer.Sprintf("Side deck section: (%s)\n", msg.TranslateCards((*g.SideDeck())[:cardN], printer))
				pr.reports = append(pr.reports, PlMsg{pl, plMsg})
			}
		}

		// end round: next player and next phase
		plInfo := next.NewPlInfo(g.Phase(), g.Players(), g.PlayedCards(), g.Briscola(),
			len(*g.SideDeck()) > 0, len(*g.PlayedCards()) < 5, rq.From())
		plIndex := next.Player(plInfo)

		phInfo := next.NewPhInfo(g.Phase(), g.Players(), g.Caller(), g.Companion(), g.Briscola(),
			len(*g.SideDeck()) > 0, len(*g.PlayedCards()) < 5, rq.Value())
		nextPhase := next.Phase(phInfo)

		current := g.Phase()
		if current == phase.InsideAuction && nextPhase > current {
			_, p := g.Players().Find(player.NotFolded)
			g.caller = p
		}

		if !(current != phase.PlayingCards || g.IsRoundOngoing()) {
			pile := g.Players()[plIndex].Pile()
			move(g.PlayedCards(), pile)
			if team.Count(g.Players(), player.IsHandEmpty) == 5 && g.IsSideUsed() {
				move(g.SideDeck(), pile)
			}
		}
		g.setPhase(nextPhase)
		track.Player(g.LastPlaying(), g.Players()[plIndex])

		// log action to console
		cons, consMsg := os.Stdout, fmt.Sprintf("New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", sender(g, rq).Name(), *rq, sender(g, rq), g)
		pr.reports = append(pr.reports, PlMsg{cons, consMsg})
		for _, pl := range g.Players() {
			pl, plMsg := pl, "-----"
			pr.reports = append(pr.reports, PlMsg{pl, plMsg})
		}
		pl, plMsg := g.LastPlayer(), msg.CreateInGameMsg(g, g.LastPlayer())
		pr.reports = append(pr.reports, PlMsg{pl, plMsg})
		for _, pl := range g.Players() {
			pl, plMsg := pl, msg.TranslateGameStatus(g, printer)
			pr.reports = append(pr.reports, PlMsg{pl, plMsg})
		}
		pl, plMsg = g.CurrentPlayer(), msg.CreateInGameMsg(g, g.CurrentPlayer())
		pr.reports = append(pr.reports, PlMsg{pl, plMsg})

		if g.phase != phase.End {
			return pr.reports
		}

		// process end phase
		remainingCards := len(*g.Players()[0].Hand())
		if remainingCards > 0 {
			collect(g)
		}
		// compute score
		pilers := make([]team.Piler, 0)
		for _, p := range g.Players() {
			pilers = append(pilers, p)
		}
		scoreTeam1, scoreTeam2 := team.Score(g.Caller(), g.Companion(), pilers, briscola.Points)
		for _, pl := range g.Players() {
			pl, plMsg := pl, printer.Sprintf("The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2)
			pr.reports = append(pr.reports, PlMsg{pl, plMsg})
		}
	}
	pr.reports = append(pr.reports, g.handleMLData()) // placeholder for ml data
	return pr.reports
}

type proc struct {
	rq      *Req
	reports []PlMsg
	err     error
}

func (pr *proc) reportErr(g interface{ Players() team.Players }, rq interface{ From() string }, err error) {
	cons, consMsg := os.Stdout, fmt.Sprintf("New Action by %s: %s\nError raised: %+v\n", sender(g, rq).Name(), rq, err)
	pl, plMsg := sender(g, rq), fmt.Sprintf("Error: %+v\n", err)
	pr.reports = append(pr.reports, PlMsg{cons, consMsg})
	pr.reports = append(pr.reports, PlMsg{pl, plMsg})
	pr.err = err
}

func sender(g interface{ Players() team.Players }, rq interface{ From() string }) *player.Player {
	_, p := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(rq.From()) })
	return p
}
