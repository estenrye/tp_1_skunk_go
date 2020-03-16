package playertest

import (
	"testing"

	"github.com/estenrye/skunk/internal/skunk/dice"
	"github.com/estenrye/skunk/internal/skunk/player"
	"github.com/estenrye/skunk/internal/skunk/turn"
	"github.com/estenrye/skunk/internal/skunk_test/dietest"
)

func assertPlayer(t *testing.T, p player.ISkunkPlayer, when string, expectedPlayerName string,
	expectedPlayerState player.State, expectedPlayerScore int, expectedPlayerChips int,
	expectTurnIsNil bool) {

	if p.GetName() != expectedPlayerName {
		t.Errorf("When %s, expected player name to be %s, not %s.", when, expectedPlayerName, p.GetName())
	}
	if p.GetLastState() != expectedPlayerState {
		t.Errorf("When %s, expected turn state to be %s, not %s", when, expectedPlayerState, p.GetLastState())
	}
	if p.GetLastScore() != expectedPlayerScore {
		t.Errorf("When %s, expected player's score to be %d, got %d", when, expectedPlayerScore, p.GetLastScore())
	}
	if p.GetLastChips() != expectedPlayerChips {
		t.Errorf("When %s, expected player's score to be %d, got %d", when, expectedPlayerChips, p.GetLastChips())
	}
	if expectTurnIsNil && p.GetLastTurn() != nil {
		t.Errorf("When player is created, expected turn to be nil.")
	}
	if !expectTurnIsNil && p.GetLastTurn() == nil {
		t.Errorf("When player is created, expected turn to be not nil.")
	}
}

func assertPlayerTurn(t *testing.T, p player.ISkunkPlayer, when string,
	expectedTurnState turn.State, expectedTurnScore int, expectedTurnPenalty int) {
	playerTurn := p.GetLastTurn()
	if playerTurn == nil {
		t.Errorf("When %s, expected turn to be not nil.", when)
	} else {
		if playerTurn.GetState() != expectedTurnState {
			t.Errorf("When %s, expected turn state to be %s, not %s.", when, expectedTurnState, playerTurn.GetState())
		}
		if playerTurn.GetScore() != expectedTurnScore {
			t.Errorf("When %s, expected turn state to be %d, not %d.", when, expectedTurnScore, playerTurn.GetScore())
		}
		if playerTurn.GetPenalty() != expectedTurnPenalty {
			t.Errorf("When %s, expected turn state to be %d, not %d.", when, expectedTurnPenalty, playerTurn.GetPenalty())
		}
	}
}

func Test_when_player_is_created_turn_is_name_and_defaults_are_initialized(t *testing.T) {
	p := player.NewPlayer("playerOne")

	assertPlayer(t, p, "player is created", "playerOne", player.TurnNotStarted, 0, 50, true)
}

func Test_when_player_is_created_and_turn_is_nil_roll_does_nothing(t *testing.T) {
	p := player.NewPlayer("playerOne")
	p.Roll()
	assertPlayer(t, p, "player is created and turn is nil and roll action is selected", "playerOne", player.TurnNotStarted, 0, 50, true)
}

func Test_when_player_is_created_and_turn_is_nil_pass_does_nothing(t *testing.T) {
	p := player.NewPlayer("playerOne")
	p.Pass()
	assertPlayer(t, p, "player is created and turn is nil and pass action is selected", "playerOne", player.TurnNotStarted, 0, 50, true)
}

func Test_when_turn_is_initialized_player_state_is_TurnNotStarted_and_turn_state_is_NotStarted(t *testing.T) {
	p := player.NewPlayer("playerOne")
	p.NewTurn()

	assertPlayer(t, p, "player is created and turn is initialized", "playerOne", player.TurnNotStarted, 0, 50, false)
	assertPlayerTurn(t, p, "player is created and turn is initialized", turn.NotStarted, 0, 0)
}

func Test_when_player_rolls_and_roll_is_not_skunk_player_state_is_active_turn_state_active(t *testing.T) {
	d1 := dietest.NewDieFromInt(5)
	d2 := dietest.NewDieFromInt(3)
	skunkDice := dice.NewDiceFromISkunkDice(d1, d2)
	playerTurn := turn.NewTurnFromISkunkDice(skunkDice)
	p := player.NewPlayer("playerOne")
	p.NewTurnFromISkunkTurn(playerTurn)
	assertPlayer(t, p, "player rolls and roll is not a skunk", "playerOne", player.TurnNotStarted, 0, 50, false)
	p.Roll()
	assertPlayer(t, p, "player rolls and roll is not a skunk", "playerOne", player.ActiveTurn, 0, 50, false)
	assertPlayerTurn(t, p, "player rolls and roll is not a skunk", turn.Active, 8, 0)
}

func Test_when_player_first_action_is_pass_player_state_is_complete_turn_state_active(t *testing.T) {
	d1 := dietest.NewDieFromInt(5)
	d2 := dietest.NewDieFromInt(3)
	skunkDice := dice.NewDiceFromISkunkDice(d1, d2)
	playerTurn := turn.NewTurnFromISkunkDice(skunkDice)
	p := player.NewPlayer("playerOne")
	p.NewTurnFromISkunkTurn(playerTurn)
	assertPlayer(t, p, "player passes on first action", "playerOne", player.TurnNotStarted, 0, 50, false)
	p.Pass()
	assertPlayer(t, p, "player passes on first action", "playerOne", player.CompleteTurn, 0, 50, false)
	assertPlayerTurn(t, p, "player rolls and roll is not a skunk", turn.Complete, 0, 0)
}
