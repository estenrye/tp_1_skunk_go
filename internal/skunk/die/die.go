package die

import (
	"math/rand"
	"time"

	"github.com/estenrye/skunk/internal/skunk"
)

// Die implements a psuedorandom six sided die.
type Die struct {
	lastRoll int
}

// NewDie creates a new Die as an IRollable instance.
func NewDie() skunk.IRollable {
	return &Die{}
}

// Roll implements the IRollable.Roll() interface method.
func (d *Die) Roll() {
	rand.Seed(time.Now().UnixNano())
	d.lastRoll = rand.Intn(6) + 1
}

// GetLastRoll implements the IRollable.GetLastRoll() interface method.
func (d *Die) GetLastRoll() int {
	return d.lastRoll
}
