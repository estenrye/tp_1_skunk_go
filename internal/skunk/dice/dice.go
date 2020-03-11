package dice

import (
	"fmt"

	"github.com/estenrye/skunk/internal/skunk/die"
)

// Dice Implements a pair of die that are pseudorandom.
type Dice struct {
	lastRoll int
	die1     die.ISkunkDie
	die2     die.ISkunkDie
}

// NewDice is the default initializer for Dice.
func NewDice() ISkunkDice {
	return NewDiceFromISkunkDice(die.NewDie(), die.NewDie())
}

// NewDiceFromISkunkDice is an initializer that enables injection of alternative die implementations.
func NewDiceFromISkunkDice(die1 die.ISkunkDie, die2 die.ISkunkDie) ISkunkDice {
	return &Dice{
		die1: die1,
		die2: die2,
	}
}

// GetLastDie1 returns the state of die1 to the caller.
func (d Dice) GetLastDie1() die.State {
	return d.die1.GetLastRoll()
}

// GetLastDie2 returns the state of die2 to the caller.
func (d Dice) GetLastDie2() die.State {
	return d.die2.GetLastRoll()
}

// GetLastState evaluates the states of die1 and die 2 returns the state of the dice roll.
func (d Dice) GetLastState() State {
	if d.GetLastDie1() == die.Skunk && d.GetLastDie2() == die.Skunk {
		return DoubleSkunk
	}

	if d.GetLastDie1() == die.Skunk {
		if d.GetLastDie2() == die.Deuce {
			return SkunkDeuce
		}
		return SingleSkunk
	}

	if d.GetLastDie2() == die.Skunk {
		if d.GetLastDie1() == die.Deuce {
			return SkunkDeuce
		}
		return SingleSkunk
	}

	if d.GetLastDie1() == die.UnknownState || d.GetLastDie2() == die.UnknownState {
		return UnknownState
	}

	return ScorableRoll
}

// Roll implements the roll action for a pair of dice
func (d *Dice) Roll() {
	d.die1.Roll()
	d.die2.Roll()
	d.lastRoll = d.GetLastDie1().ToInt() + d.GetLastDie2().ToInt()

	if d.GetLastDie1() == die.Skunk || d.GetLastDie2() == die.Skunk {
		d.lastRoll = 0
	}
}

// GetLastRoll returns the score of the last dice roll.
func (d Dice) GetLastRoll() int {
	return d.lastRoll
}

func (d Dice) String() string {
	return fmt.Sprintf("Die1: %s, Die2: %s", d.die1.GetLastRoll(), d.die2.GetLastRoll())
}
