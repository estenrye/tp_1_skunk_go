package dietest

import (
	"github.com/estenrye/skunk/internal/skunk"
)

// MockDie implements a psuedorandom six sided die.
type MockDie struct {
	lastRoll      int
	expectedRolls []int
	index         int
}

// NewDieFromInt creates a new Die as an IRollable instance.
func NewDieFromInt(rolls int) skunk.IRollable {
	return NewDieFromArray([]int{rolls})
}

// NewDieFromArray creates a new Die as an IRollable instance.
func NewDieFromArray(rolls []int) skunk.IRollable {
	return &MockDie{
		expectedRolls: rolls,
		index:         -1,
	}
}

// Roll implements the IRollable.Roll() interface method.
func (d *MockDie) Roll() {
	d.index = (d.index + 1) % len(d.expectedRolls)
	d.lastRoll = d.expectedRolls[d.index]
}

// GetLastRoll implements the IRollable.GetLastRoll() interface method.
func (d *MockDie) GetLastRoll() int {
	return d.lastRoll
}
