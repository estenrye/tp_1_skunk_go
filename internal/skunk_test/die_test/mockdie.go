package dietest

import (
	"github.com/estenrye/skunk/internal/skunk/die"
)

// MockDie implements a psuedorandom six sided die.
type MockDie struct {
	lastRoll      die.State
	expectedRolls []int
	index         int
}

// NewDieFromInt creates a new Die as an ISkunkDie instance.
func NewDieFromInt(rolls int) die.ISkunkDie {
	return NewDieFromArray([]int{rolls})
}

// NewDieFromArray creates a new Die as an ISkunkDie instance.
func NewDieFromArray(rolls []int) die.ISkunkDie {
	return &MockDie{
		expectedRolls: rolls,
		index:         -1,
	}
}

// Roll implements the ISkunkDie.Roll() interface method.
func (d *MockDie) Roll() {
	d.index = (d.index + 1) % len(d.expectedRolls)
	d.lastRoll = die.State(d.expectedRolls[d.index])
}

// GetLastRoll implements the ISkunkDie.GetLastRoll() interface method.
func (d *MockDie) GetLastRoll() die.State {
	return d.lastRoll
}
