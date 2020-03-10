package die

import (
	"math/rand"
	"time"
)

// Die implements a psuedorandom six sided die.
type Die struct {
	lastRoll State
}

// NewDie creates a new Die as an ISkunkDie instance.
func NewDie() ISkunkDie {
	return &Die{}
}

// Roll implements the ISkunkDie.Roll() interface method.
func (d *Die) Roll() {
	rand.Seed(time.Now().UnixNano())
	d.lastRoll = State(rand.Intn(6) + 1)
}

// GetLastRoll implements the ISkunkDie.GetLastRoll() interface method.
func (d *Die) GetLastRoll() State {
	return d.lastRoll
}
