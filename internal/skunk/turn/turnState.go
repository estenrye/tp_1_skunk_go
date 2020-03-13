package turn

// State represents the state of a pair of dice in the game of skunk.
type State int

const (
	// NotStarted represents a turn state where the dice have not yet been rolled.
	NotStarted State = 0
	// Active represents a turn state where the user has decided to Roll as their first action.
	Active State = 1
	// Complete represents a turn state where the user passes their turn before they roll a skunk.
	Complete State = 2
	// CompleteResetScore represents a turn state where the user passes their turn before they roll a skunk.
	CompleteResetScore State = 3
	// UnknownState represents a dice state where the dice have not yet been rolled.
	UnknownState State = 4
)

func (state State) String() string {
	names := [...]string{
		"Not Started",
		"Active",
		"Complete",
		"Complete Reset Score",
		"Unknown State",
	}

	if state < NotStarted || state > CompleteResetScore {
		return "Unknown State"
	}

	return names[state]
}
