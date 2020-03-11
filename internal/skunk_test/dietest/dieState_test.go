package dietest

import (
	"testing"

	"github.com/estenrye/skunk/internal/skunk/die"
)

func Test_given_state_constant_returns_expected_string_value(t *testing.T) {
	if die.Skunk.String() != "Skunk" {
		t.Errorf("Skunk.String() returned '%s', expected '%s'", die.Skunk.String(), "Skunk")
	}
	if die.Deuce.String() != "Deuce" {
		t.Errorf("Deuce.String() returned '%s', expected '%s'", die.Deuce.String(), "Deuce")
	}
	if die.Three.String() != "Three" {
		t.Errorf("Three.String() returned '%s', expected '%s'", die.Three.String(), "Three")
	}
	if die.Four.String() != "Four" {
		t.Errorf("Four.String() returned '%s', expected '%s'", die.Four.String(), "Four")
	}
	if die.Five.String() != "Five" {
		t.Errorf("Five.String() returned '%s', expected '%s'", die.Five.String(), "Five")
	}
	if die.Six.String() != "Six" {
		t.Errorf("Six.String() returned '%s', expected '%s'", die.Five.String(), "Six")
	}
	if die.State(0).String() != "Unknown State" {
		t.Errorf("die.State(0).String() returned '%s', expected '%s'", die.State(0).String(), "Unknown State")
	}
	if die.State(7).String() != "Unknown State" {
		t.Errorf("die.State(7).String() returned '%s', expected '%s'", die.State(7).String(), "Unknown State")
	}
}

func Test_given_state_constant_returns_expected_int_value(t *testing.T) {
	if die.Skunk.ToInt() != 1 {
		t.Errorf("Skunk.ToInt() returned '%d', expected '%d'", die.Skunk.ToInt(), 1)
	}
	if die.Deuce.ToInt() != 2 {
		t.Errorf("Deuce.ToInt() returned '%d', expected '%d'", die.Deuce.ToInt(), 2)
	}
	if die.Three.ToInt() != 3 {
		t.Errorf("Three.ToInt() returned '%d', expected '%d'", die.Three.ToInt(), 3)
	}
	if die.Four.ToInt() != 4 {
		t.Errorf("Four.ToInt() returned '%d', expected '%d'", die.Four.ToInt(), 4)
	}
	if die.Five.ToInt() != 5 {
		t.Errorf("Five.ToInt() returned '%d', expected '%d'", die.Five.ToInt(), 5)
	}
	if die.Six.ToInt() != 6 {
		t.Errorf("Six.ToInt() returned '%d', expected '%d'", die.Five.ToInt(), 6)
	}
	if die.State(0).ToInt() != 0 {
		t.Errorf("die.State(0).ToInt() returned '%d', expected '%d'", die.State(0).ToInt(), 0)
	}
	if die.State(7).ToInt() != 0 {
		t.Errorf("die.State(7).ToInt() returned '%d', expected '%d'", die.State(7).ToInt(), 0)
	}
}
