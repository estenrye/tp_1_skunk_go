package turn

import (
	"github.com/estenrye/skunk/internal/skunk/dice"
)

// Turn represents a single player's actions in a round.
type Turn struct {
	skunkDice dice.ISkunkDice
}

// NewTurnFromISkunkDice creates a turn from an initialized ISkunkDice object
func NewTurnFromISkunkDice(roll dice.ISkunkDice) ISkunkTurn {
	return &Turn{
		skunkDice: roll,
	}
}

// Roll performs the player's roll action in the turn and rolls the dice.
func (t *Turn) Roll() {

}

// Pass performs the player's pass action in the turn, completing their turn.
func (t *Turn) Pass() {

}

// GetScore returns the player's current turn score.
func (t Turn) GetScore() int {
	return -1
}

// GetPenalty returns the player's current turn penalty.
func (t Turn) GetPenalty() int {
	return -5
}

// GetState returns the player's turn state.
func (t Turn) GetState() State {
	return UnknownState
}
