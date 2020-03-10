package die

// State represents the various states of a Die.
type State int

const (
	// Skunk represents the state of a one.
	Skunk State = 1
	// Deuce represents the state of a two.
	Deuce State = 2
	// Three represents the state of a three.
	Three State = 3
	// Four represents the state of a four.
	Four State = 4
	// Five represents the state of a five.
	Five State = 5
	// Six represents the state of a six.
	Six State = 6
)

// String() converts the die.State to a human readable string.
func (state State) String() string {
	names := [...]string{
		"Unknown State",
		"Skunk",
		"Deuce",
		"Three",
		"Four",
		"Five",
		"Six",
	}

	if state < Skunk || state > Six {
		return "Unknown State"
	}

	return names[state]
}

// ToInt converts a die.State type to an int type.
func (state State) ToInt() int {
	values := [...]int{0, 1, 2, 3, 4, 5, 6}

	if state < Skunk || state > Six {
		return 0
	}

	return values[state]
}
