package turn

import (
	"fmt"

	"github.com/estenrye/skunk/internal/skunk/dice"
)

// Turn represents a single player's actions in a round.
type Turn struct {
	skunkDice dice.ISkunkDice
	score     int
	penalty   int
	state     State
}

// NewTurn creates a player turn.
func NewTurn() ISkunkTurn {
	return NewTurnFromISkunkDice(dice.NewDice())
}

// NewTurnFromISkunkDice creates a turn from an initialized ISkunkDice object
func NewTurnFromISkunkDice(roll dice.ISkunkDice) ISkunkTurn {
	return &Turn{
		skunkDice: roll,
		penalty:   0,
		state:     NotStarted,
	}
}

// Roll performs the player's roll action in the turn and rolls the dice.
func (t *Turn) Roll() {
	if t.state != Active && t.state != NotStarted {
		return
	}

	t.skunkDice.Roll()
	if t.skunkDice.GetLastState() == dice.ScorableRoll {
		t.score += t.skunkDice.GetLastRoll()
		t.state = Active
	}
	if t.skunkDice.GetLastState() == dice.SingleSkunk {
		t.score = t.skunkDice.GetLastRoll()
		t.penalty = 1
		t.state = Complete
	}
	if t.skunkDice.GetLastState() == dice.SkunkDeuce {
		t.score = t.skunkDice.GetLastRoll()
		t.penalty = 2
		t.state = Complete
	}
	if t.skunkDice.GetLastState() == dice.DoubleSkunk {
		t.score = t.skunkDice.GetLastRoll()
		t.penalty = 4
		t.state = CompleteResetScore
	}
}

// Pass performs the player's pass action in the turn, completing their turn.
func (t *Turn) Pass() {
	if t.state != Active && t.state != NotStarted {
		return
	}

	t.state = Complete
}

// GetScore returns the player's current turn score.
func (t Turn) GetScore() int {
	return t.score
}

// GetPenalty returns the player's current turn penalty.
func (t Turn) GetPenalty() int {
	return t.penalty
}

// GetState returns the player's turn state.
func (t Turn) GetState() State {
	return t.state
}

// GetLastRoll returns the string representation of the last dice roll.
func (t Turn) GetLastRoll() string {
	return fmt.Sprintf("Die1: %d, Die2: %d", t.skunkDice.GetLastDie1().ToInt(), t.skunkDice.GetLastDie2().ToInt())
}
