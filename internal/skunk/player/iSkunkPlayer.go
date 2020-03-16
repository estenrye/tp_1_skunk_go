package player

import "github.com/estenrye/skunk/internal/skunk/turn"

// ISkunkPlayer represents a
type ISkunkPlayer interface {
	NewTurn()
	NewTurnFromISkunkTurn(turn turn.ISkunkTurn)
	Roll()
	Pass()
	GetName() string
	GetLastScore() int
	GetLastChips() int
	GetLastTurn() turn.ISkunkTurnResult
	GetLastState() State
}
