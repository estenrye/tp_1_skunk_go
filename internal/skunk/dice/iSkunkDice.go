package dice

import (
	"github.com/estenrye/skunk/internal/skunk/die"
)

// ISkunkDice implements an interface for performing a roll action and getting its result.
type ISkunkDice interface {
	Roll()
	GetLastRoll() int
	GetLastDie1() die.State
	GetLastDie2() die.State
	GetLastState() State
}
