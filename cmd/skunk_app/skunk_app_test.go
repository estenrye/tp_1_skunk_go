package main

import (
	"testing"

	"github.com/estenrye/skunk/internal/skunk/player"
)

func Test_getTurnInput(t *testing.T) {
	p := player.NewPlayer("playerOne")
	p.NewTurn()

	getTurnInput(p, "a\r\n")
	if p.GetLastState() != player.TurnNotStarted {
		t.Errorf("Expected input of string value other than 'r' or 'p' to noop, player state: %s.", p.GetLastState())
	}

	p.Roll()
	if p.GetLastState() == player.TurnNotStarted {
		t.Errorf("Expected input of string value 'r' to change player state, got %s.", p.GetLastState())
	}

	p = player.NewPlayer("playerOne")
	p.NewTurn()
	p.Pass()
	if p.GetLastState() != player.CompleteTurn {
		t.Errorf("Expected input of string value 'p' to change player state to 'Complete Turn', got %s", p.GetLastState())
	}
}
