package dice

// State represents the state of a pair of dice in the game of skunk.
type State int

const (
	// SingleSkunk represents a dice state where one die is a skunk and the other die is >= 3
	SingleSkunk State = 0
	// DoubleSkunk represents a dice state where both dice are a skunk.
	DoubleSkunk State = 1
	// SkunkDeuce represents a dice state where one die is a skunk and the other is a deuce.
	SkunkDeuce State = 2
	// ScorableRoll represents a dice state where the sum is scorable.
	ScorableRoll State = 3
	// UnknownState represents a dice state where the dice have not yet been rolled.
	UnknownState State = 4
)

func (state State) String() string {
	names := [...]string{
		"Single Skunk",
		"Double Skunk",
		"Skunk Deuce",
		"Scorable Roll",
		"Unknown State",
	}

	if state < SingleSkunk || state > ScorableRoll {
		return "Unknown State"
	}

	return names[state]
}

// GetPenalty converts the dice.State object into an integer penalty value.
func (state State) GetPenalty() int {
	values := [...]int{1, 4, 2, 0}

	if state < SingleSkunk || state > ScorableRoll {
		return 0
	}

	return values[state]
}
