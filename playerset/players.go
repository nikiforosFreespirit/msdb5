package playerset

import (
	"errors"

	"github.com/nikiforosFreespirit/msdb5/player"
)

// Players struct
type Players []*player.Player

// Add func
func (playerSet *Players) Add(p player.Player) {
	*playerSet = append(*playerSet, &p)
}

// Find func
func (playerSet Players) Find(predicate func(p *player.Player) bool) (*player.Player, error) {
	for _, p := range playerSet {
		if predicate(p) {
			return p, nil
		}
	}
	return nil, errors.New("Player not found")
}

// Count func
func (playerSet Players) Count(predicate func(p *player.Player) bool) (count uint8) {
	for _, p := range playerSet {
		if predicate(p) {
			count++
		}
	}
	return
}

// All func
func (playerSet Players) All(predicate func(p *player.Player) bool) bool {
	for _, p := range playerSet {
		if !predicate(p) {
			return false
		}
	}
	return true
}
