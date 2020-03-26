package die

import (
	rand "math/rand"
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
	var src cryptoSource
	rnd := rand.New(src)
	d.lastRoll = State(rnd.Intn(6) + 1)
}

// GetLastRoll implements the ISkunkDie.GetLastRoll() interface method.
func (d *Die) GetLastRoll() State {
	return d.lastRoll
}
