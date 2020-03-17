package playertest

import (
	"testing"

	"github.com/estenrye/skunk/internal/skunk/dice"
	"github.com/estenrye/skunk/internal/skunk/die"
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
		t.Errorf("When %s, expected player state to be %s, not %s", when, expectedPlayerState, p.GetLastState())
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

func Test_when_player_first_action_is_pass_player_state_is_complete_turn_state_active(t *testing.T) {
	d1 := dietest.NewDieFromInt(5)
	d2 := dietest.NewDieFromInt(3)
	when := "player passes on first action"
	playerName := "playerOne"
	p := player.NewPlayer(playerName)
	assertPlayer(t, p, when, playerName, player.TurnNotStarted, 0, 50, true)
	p.NewTurnFromISkunkTurn(turn.NewTurnFromISkunkDice(dice.NewDiceFromISkunkDice(d1, d2)))
	assertPlayer(t, p, when, playerName, player.TurnNotStarted, 0, 50, false)
	assertPlayerTurn(t, p, when, turn.NotStarted, 0, 0)
	simulatePass(t, when, p, playerName, 0, 50, player.CompleteTurn, 0, 0, turn.Complete)
}

func simulateRolls(t *testing.T, when string, p player.ISkunkPlayer, playerName string,
	d1 die.ISkunkDie, d2 die.ISkunkDie,
	turnScores []int, turnPenalties []int, turnStates []turn.State,
	expectedPlayerScores []int, expectedPlayerChips []int, expectedPlayerStates []player.State) {

	numRolls := len(turnScores)
	p.NewTurnFromISkunkTurn(turn.NewTurnFromISkunkDice(dice.NewDiceFromISkunkDice(d1, d2)))
	assertPlayer(t, p, when, playerName, expectedPlayerStates[0], expectedPlayerScores[0], expectedPlayerChips[0], false)
	assertPlayerTurn(t, p, when, turnStates[0], turnScores[0], turnPenalties[0])
	for i := 1; i < numRolls; i++ {
		p.Roll()
		assertPlayer(t, p, when, playerName, expectedPlayerStates[i], expectedPlayerScores[i], expectedPlayerChips[i], false)
		assertPlayerTurn(t, p, when, turnStates[i], turnScores[i], turnPenalties[i])
	}
}

func simulatePass(t *testing.T, when string, p player.ISkunkPlayer, playerName string,
	expectedPlayerScore int, expectedPlayerChips int, expectedPlayerState player.State,
	expectedTurnScore int, expectedTurnPenalty int, expectedTurnState turn.State) {

	p.Pass()
	assertPlayer(t, p, when, playerName, expectedPlayerState, expectedPlayerScore, expectedPlayerChips, false)
	assertPlayerTurn(t, p, when, expectedTurnState, expectedTurnScore, expectedTurnPenalty)
}

func Test_when_player_passes_turn_score_is_banked_into_player_score(t *testing.T) {
	d1 := dietest.NewDieFromArray([]int{4, 5, 2, 3, 2})
	d2 := dietest.NewDieFromArray([]int{5, 5, 3, 6, 2})
	turnScores := []int{0, 9, 19, 24, 33, 37}
	turnPenalties := []int{0, 0, 0, 0, 0, 0}
	turnStates := []turn.State{turn.NotStarted, turn.Active, turn.Active, turn.Active, turn.Active, turn.Active}
	playerScores := []int{0, 0, 0, 0, 0, 0}
	playerChips := []int{50, 50, 50, 50, 50, 50}
	playerStates := []player.State{
		player.TurnNotStarted,
		player.ActiveTurn,
		player.ActiveTurn,
		player.ActiveTurn, player.ActiveTurn, player.ActiveTurn}
	when := "player passes turn score is banked into player score"
	playerName := "playerOne"

	p := player.NewPlayer(playerName)
	simulateRolls(t, when, p, playerName, d1, d2, turnScores, turnPenalties, turnStates, playerScores, playerChips, playerStates)
	simulatePass(t, when, p, playerName, 37, 50, player.CompleteTurn, 37, 0, turn.Complete)
	playerScores = []int{37, 37, 37, 37, 37, 37}
	simulateRolls(t, when, p, playerName, d1, d2, turnScores, turnPenalties, turnStates, playerScores, playerChips, playerStates)
	simulatePass(t, when, p, playerName, 74, 50, player.CompleteTurn, 37, 0, turn.Complete)
}

func Test_when_player_rolls_single_skunk_penalties_are_applied(t *testing.T) {
	d1 := dietest.NewDieFromArray([]int{4, 5, 2, 3, 2})
	d2 := dietest.NewDieFromArray([]int{5, 5, 3, 6, 2})
	turnScores := []int{0, 9, 19, 24, 33, 37}
	turnPenalties := []int{0, 0, 0, 0, 0, 0}
	turnStates := []turn.State{turn.NotStarted, turn.Active, turn.Active, turn.Active, turn.Active, turn.Active}
	playerScores := []int{0, 0, 0, 0, 0, 0}
	playerChips := []int{50, 50, 50, 50, 50, 50}
	playerStates := []player.State{player.TurnNotStarted, player.ActiveTurn, player.ActiveTurn, player.ActiveTurn, player.ActiveTurn, player.ActiveTurn}
	when := "player rolls a single skunk"
	playerName := "playerOne"

	p := player.NewPlayer(playerName)
	simulateRolls(t, when, p, playerName, d1, d2, turnScores, turnPenalties, turnStates, playerScores, playerChips, playerStates)
	simulatePass(t, when, p, playerName, 37, 50, player.CompleteTurn, 37, 0, turn.Complete)

	d1 = dietest.NewDieFromArray([]int{5, 4, 1})
	d2 = dietest.NewDieFromArray([]int{6, 2, 4})
	turnScores = []int{0, 11, 17, 0}
	turnPenalties = []int{0, 0, 0, 1}
	turnStates = []turn.State{turn.NotStarted, turn.Active, turn.Active, turn.Complete}
	playerStates = []player.State{player.TurnNotStarted, player.ActiveTurn, player.ActiveTurn, player.CompleteTurn}
	playerScores = []int{37, 37, 37, 37}
	playerChips = []int{50, 50, 50, 49}
	simulateRolls(t, when, p, playerName, d1, d2, turnScores, turnPenalties, turnStates, playerScores, playerChips, playerStates)
}

func Test_when_player_rolls_skunk_deuce_penalties_are_applied(t *testing.T) {
	d1 := dietest.NewDieFromArray([]int{4, 5, 2, 3, 2})
	d2 := dietest.NewDieFromArray([]int{5, 5, 3, 6, 2})
	turnScores := []int{0, 9, 19, 24, 33, 37}
	turnPenalties := []int{0, 0, 0, 0, 0, 0}
	turnStates := []turn.State{turn.NotStarted, turn.Active, turn.Active, turn.Active, turn.Active, turn.Active}
	playerScores := []int{0, 0, 0, 0, 0, 0}
	playerChips := []int{50, 50, 50, 50, 50, 50}
	playerStates := []player.State{player.TurnNotStarted, player.ActiveTurn, player.ActiveTurn, player.ActiveTurn, player.ActiveTurn, player.ActiveTurn}
	when := "player rolls a skunk deuce"
	playerName := "playerOne"

	p := player.NewPlayer(playerName)
	simulateRolls(t, when, p, playerName, d1, d2, turnScores, turnPenalties, turnStates, playerScores, playerChips, playerStates)
	simulatePass(t, when, p, playerName, 37, 50, player.CompleteTurn, 37, 0, turn.Complete)

	d1 = dietest.NewDieFromArray([]int{5, 4, 1})
	d2 = dietest.NewDieFromArray([]int{6, 2, 2})
	turnScores = []int{0, 11, 17, 0}
	turnPenalties = []int{0, 0, 0, 2}
	turnStates = []turn.State{turn.NotStarted, turn.Active, turn.Active, turn.Complete}
	playerStates = []player.State{player.TurnNotStarted, player.ActiveTurn, player.ActiveTurn, player.CompleteTurn}
	playerScores = []int{37, 37, 37, 37}
	playerChips = []int{50, 50, 50, 48}
	simulateRolls(t, when, p, playerName, d1, d2, turnScores, turnPenalties, turnStates, playerScores, playerChips, playerStates)
}

func Test_when_player_rolls_double_skunk_penalties_are_applied(t *testing.T) {
	d1 := dietest.NewDieFromArray([]int{4, 5, 2, 3, 2})
	d2 := dietest.NewDieFromArray([]int{5, 5, 3, 6, 2})
	turnScores := []int{0, 9, 19, 24, 33, 37}
	turnPenalties := []int{0, 0, 0, 0, 0, 0}
	turnStates := []turn.State{turn.NotStarted, turn.Active, turn.Active, turn.Active, turn.Active, turn.Active}
	playerScores := []int{0, 0, 0, 0, 0, 0}
	playerChips := []int{50, 50, 50, 50, 50, 50}
	playerStates := []player.State{player.TurnNotStarted, player.ActiveTurn, player.ActiveTurn, player.ActiveTurn, player.ActiveTurn, player.ActiveTurn}
	when := "player rolls a double skunk"
	playerName := "playerOne"

	p := player.NewPlayer(playerName)
	simulateRolls(t, when, p, playerName, d1, d2, turnScores, turnPenalties, turnStates, playerScores, playerChips, playerStates)
	simulatePass(t, when, p, playerName, 37, 50, player.CompleteTurn, 37, 0, turn.Complete)

	d1 = dietest.NewDieFromArray([]int{5, 4, 1})
	d2 = dietest.NewDieFromArray([]int{6, 2, 1})
	turnScores = []int{0, 11, 17, 0}
	turnPenalties = []int{0, 0, 0, 4}
	turnStates = []turn.State{turn.NotStarted, turn.Active, turn.Active, turn.CompleteResetScore}
	playerStates = []player.State{player.TurnNotStarted, player.ActiveTurn, player.ActiveTurn, player.CompleteTurn}
	playerScores = []int{37, 37, 37, 0}
	playerChips = []int{50, 50, 50, 46}
	simulateRolls(t, when, p, playerName, d1, d2, turnScores, turnPenalties, turnStates, playerScores, playerChips, playerStates)
}

func Test_when_player_passes_with_score_greater_than_or_equal_to_100(t *testing.T) {
	d1 := dietest.NewDieFromArray([]int{4, 5, 2, 3, 2})
	d2 := dietest.NewDieFromArray([]int{5, 5, 3, 6, 2})
	turnScores := []int{0, 9, 19, 24, 33, 37}
	turnPenalties := []int{0, 0, 0, 0, 0, 0}
	turnStates := []turn.State{turn.NotStarted, turn.Active, turn.Active, turn.Active, turn.Active, turn.Active}
	playerScores := []int{0, 0, 0, 0, 0, 0}
	playerChips := []int{50, 50, 50, 50, 50, 50}
	playerStates := []player.State{
		player.TurnNotStarted,
		player.ActiveTurn,
		player.ActiveTurn,
		player.ActiveTurn, player.ActiveTurn, player.ActiveTurn}
	when := "when player passes with score >= 100"
	playerName := "playerOne"

	p := player.NewPlayer(playerName)
	simulateRolls(t, when, p, playerName, d1, d2, turnScores, turnPenalties, turnStates, playerScores, playerChips, playerStates)
	simulatePass(t, when, p, playerName, 37, 50, player.CompleteTurn, 37, 0, turn.Complete)
	playerScores = []int{37, 37, 37, 37, 37, 37}
	simulateRolls(t, when, p, playerName, d1, d2, turnScores, turnPenalties, turnStates, playerScores, playerChips, playerStates)
	simulatePass(t, when, p, playerName, 74, 50, player.CompleteTurn, 37, 0, turn.Complete)
	playerScores = []int{74, 74, 74, 74, 74, 74}
	simulateRolls(t, when, p, playerName, d1, d2, turnScores, turnPenalties, turnStates, playerScores, playerChips, playerStates)
	simulatePass(t, when, p, playerName, 111, 50, player.CompleteEndgame, 37, 0, turn.Complete)
}
