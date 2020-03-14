package player

import "github.com/estenrye/skunk/internal/skunk/turn"

// ISkunkPlayer represents a
type ISkunkPlayer interface {
	NewTurn()
	Roll()
	Pass()
	GetName() string
	GetLastScore() int
	GetLastChips() int
	GetLastTurn() turn.ISkunkTurnResult
}
