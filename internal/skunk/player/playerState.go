package player

// State represents the state of a pair of dice in the game of skunk.
type State int

const (
	// TurnNotStarted represents a player state where the player has not started their turn.
	TurnNotStarted State = 0
	// ActiveTurn represents a player state where the player is actively playing their turn.
	ActiveTurn State = 1
	// CompleteTurn represents a player state where the player has finished their turn but have not triggered the endgame state.
	CompleteTurn State = 2
	// CompleteEndgame represents a player state where the player has finished their turn and triggered the endgame state.
	CompleteEndgame State = 3
	// UnknownState represents an undefined player state.
	UnknownState State = 4
)

func (state State) String() string {
	names := [...]string{
		"Turn Not Started",
		"Active Turn",
		"Complete Turn",
		"Complete Endgame",
		"Unknown State",
	}

	if state < TurnNotStarted || state > CompleteEndgame {
		return "Unknown State"
	}

	return names[state]
}
