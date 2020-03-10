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

func (state State) String() string {
	names := [...]string{
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
