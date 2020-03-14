package dice

import (
	"github.com/estenrye/skunk/internal/skunk/die"
)

// ISkunkDiceResult implements an interface for passing a roll data to the presentation layer.
type ISkunkDiceResult interface {
	GetLastRoll() int
	GetLastDie1() die.State
	GetLastDie2() die.State
	GetLastState() State
}
