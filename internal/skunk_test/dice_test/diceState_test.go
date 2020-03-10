package dicetest

import (
	"testing"

	"github.com/estenrye/skunk/internal/skunk/dice"
)

func Test_given_state_constant_returns_expected_string_value(t *testing.T) {
	if dice.SingleSkunk.String() != "Single Skunk" {
		t.Errorf("dice.SingleSkunk.String() returned '%s', expected '%s'.", dice.SingleSkunk.String(), "Single Skunk")
	}
	if dice.DoubleSkunk.String() != "Double Skunk" {
		t.Errorf("dice.DoubleSkunk.String() returned '%s', expected '%s'.", dice.DoubleSkunk.String(), "Double Skunk")
	}
	if dice.SkunkDeuce.String() != "Skunk Deuce" {
		t.Errorf("dice.SkunkDeuce.String() returned '%s', expected '%s'.", dice.SkunkDeuce.String(), "Skunk Deuce")
	}
	if dice.ScorableRoll.String() != "Scorable Roll" {
		t.Errorf("dice.ScorableRoll.String() returned '%s', expected '%s'.", dice.ScorableRoll.String(), "Scorable Roll")
	}
	if dice.State(-1).String() != "Unknown State" {
		t.Errorf("dice.State(-1).String() returned '%s', expected '%s'.", dice.State(-1).String(), "Unknown State")
	}
	if dice.State(4).String() != "Unknown State" {
		t.Errorf("dice.State(4).String() returned '%s', expected '%s'.", dice.State(4).String(), "Unknown State")
	}
}

func Test_given_state_constant_returns_expected_penalty_value(t *testing.T) {
	if dice.SingleSkunk.GetPenalty() != 1 {
		t.Errorf("dice.SingleSkunk.GetPenalty() returned '%d', expected '%d'.", dice.SingleSkunk.GetPenalty(), 1)
	}
	if dice.DoubleSkunk.GetPenalty() != 4 {
		t.Errorf("dice.DoubleSkunk.GetPenalty() returned '%d', expected '%d'.", dice.DoubleSkunk.GetPenalty(), 4)
	}
	if dice.SkunkDeuce.GetPenalty() != 2 {
		t.Errorf("dice.SkunkDeuce.GetPenalty() returned '%d', expected '%d'.", dice.SkunkDeuce.GetPenalty(), 2)
	}
	if dice.ScorableRoll.GetPenalty() != 0 {
		t.Errorf("dice.ScorableRoll.GetPenalty() returned '%d', expected '%d'.", dice.ScorableRoll.GetPenalty(), 0)
	}
	if dice.State(-1).GetPenalty() != 0 {
		t.Errorf("dice.State(-1).GetPenalty() returned '%d', expected '%d'.", dice.State(-1).GetPenalty(), 0)
	}
	if dice.State(4).GetPenalty() != 0 {
		t.Errorf("dice.State(4).GetPenalty() returned '%d', expected '%d'.", dice.State(4).GetPenalty(), 0)
	}
}
