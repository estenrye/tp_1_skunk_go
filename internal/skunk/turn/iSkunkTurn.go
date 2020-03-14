package turn

import (
	"github.com/estenrye/skunk/internal/skunk/dice"
)

// ISkunkTurn provides an interface to interact with a turn object
type ISkunkTurn interface {
	Roll()
	Pass()
	GetScore() int
	GetPenalty() int
	GetState() State
	GetLastRoll() dice.ISkunkDiceResult
}
