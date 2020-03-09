package skunk

// Die implements a psuedorandom six sided die.
type Die struct {
}

// Roll implements the IRollable.Roll() interface method.
func (d Die) Roll() {}

// GetLastRoll implements the IRollable.GetLastRoll() interface method.
func (d Die) GetLastRoll() int {
	return -1
}
