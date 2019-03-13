package action

import (
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type PhaseSupplier interface {
	Phase() game.Phase
}

type Finder interface {
	Find(*player.Player) bool
}

type Executer interface {
	Do(*player.Player) error
}

type NextPlayerSelector interface {
	NextPlayer(uint8) uint8
}

type NextPhaseChanger interface {
	NextPhase(playerset.Players, PlayerPredicate) bool
}

type PlayerPredicate interface {
	NextPhasePlayerInfo(*player.Player) bool
}

type Action interface {
	PhaseSupplier
	Finder
	Executer
	NextPlayerSelector
	NextPhaseChanger
	PlayerPredicate
}

func nextPlayerInTurn(playerInTurn uint8) uint8 { return (playerInTurn + 1) % 5 }
