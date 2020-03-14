package turn

import (
	"github.com/estenrye/skunk/internal/skunk/dice"
)

// ISkunkTurnResult provides an interface for passing turn data back to the presentation layer.
type ISkunkTurnResult interface {
	GetScore() int
	GetPenalty() int
	GetState() State
	GetLastRoll() dice.ISkunkDiceResult
}
