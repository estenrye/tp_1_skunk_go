package turntest

import (
	"fmt"
	"testing"

	"github.com/estenrye/skunk/internal/skunk/dice"
	"github.com/estenrye/skunk/internal/skunk/die"
	"github.com/estenrye/skunk/internal/skunk/turn"
	"github.com/estenrye/skunk/internal/skunk_test/dietest"
)

func assertTurn(t *testing.T, when string,
	expectedScore int, expectedPenalty int, expectedState turn.State, playerTurn turn.ISkunkTurn,
) {
	if playerTurn.GetScore() != expectedScore {
		t.Errorf("When %s, expected score to be %d, got %d", when, expectedScore, playerTurn.GetScore())
	}
	if playerTurn.GetPenalty() != expectedPenalty {
		t.Errorf("When %s, expected penalty to be %d, got %d", when, expectedPenalty, playerTurn.GetPenalty())
	}
	if playerTurn.GetState() != expectedState {
		t.Errorf("When %s, expected state to be %s, got %s", when, expectedState, playerTurn.GetState())
	}
}

func Test_when_a_single_skunk_is_rolled_score_is_zero_penalty_is_one_and_state_is_complete(t *testing.T) {
	for i := 3; i <= 6; i++ {
		die1 := dietest.NewDieFromDieState(die.Skunk)
		die2 := dietest.NewDieFromInt(i)
		skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
		playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)
		playerTurn.Roll()

		assertTurn(t, "a single skunk is rolled", 0, 1, turn.Complete, playerTurn)
	}
}

func Test_when_a_skunk_deuce_is_rolled_score_is_zero_penalty_is_two_and_state_is_complete(t *testing.T) {
	die1 := dietest.NewDieFromDieState(die.Skunk)
	die2 := dietest.NewDieFromDieState(die.Deuce)
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)
	playerTurn.Roll()

	assertTurn(t, "a skunk deuce is rolled", 0, 2, turn.Complete, playerTurn)
}

func Test_when_a_double_skunk_is_rolled_score_is_zero_penalty_is_four_and_state_is_completeResetScore(t *testing.T) {
	die1 := dietest.NewDieFromDieState(die.Skunk)
	die2 := dietest.NewDieFromDieState(die.Skunk)
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)
	playerTurn.Roll()

	assertTurn(t, "a double skunk is rolled", 0, 4, turn.CompleteResetScore, playerTurn)
}

func Test_when_no_skunk_is_rolled_score_is_sum_penalty_is_zero_and_state_is_active(t *testing.T) {
	for i := 2; i <= 6; i++ {
		for j := 2; j <= 6; j++ {
			die1 := dietest.NewDieFromInt(i)
			die2 := dietest.NewDieFromInt(j)
			skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
			playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)
			playerTurn.Roll()
			playerTurn.Roll()

			assertTurn(t,
				fmt.Sprintf("a %s and a %s are rolled twice", die.State(i), die.State(j)),
				2*(i+j), 0, turn.Active,
				playerTurn)
		}
	}
}

func Test_when_no_actions_have_been_taken_state_is_NotStarted_score_is_zero_penatly_is_zero(t *testing.T) {
	die1 := dietest.NewDieFromDieState(die.Skunk)
	die2 := dietest.NewDieFromDieState(die.Skunk)
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)

	assertTurn(t, "no actions have been taken", 0, 0, turn.NotStarted, playerTurn)
}

func Test_when_pass_is_taken_as_first_action_score_is_zero_penalty_is_zero_state_is_complete(t *testing.T) {
	die1 := dietest.NewDieFromDieState(die.Skunk)
	die2 := dietest.NewDieFromDieState(die.Skunk)
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)
	playerTurn.Pass()

	assertTurn(t, "player passes on the first turn", 0, 0, turn.Complete, playerTurn)
}

func Test_when_pass_is_taken_as_an_action_after_a_scorable_roll_subsequent_rolls_have_no_effect(t *testing.T) {
	die1 := dietest.NewDieFromArray([]int{6, 3, 4, 2, 5, 1, 1, 1})
	die2 := dietest.NewDieFromArray([]int{4, 6, 2, 3, 4, 5, 2, 1})
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)

	for i := 0; i < 9; i++ {
		if i != 4 {
			playerTurn.Roll()
		} else {
			playerTurn.Pass()
		}
	}

	assertTurn(t, "pass is taken as an action after a series of scorable rolls", 30, 0, turn.Complete, playerTurn)
}

func Test_when_a_skunk_is_rolled_pass_action_has_no_effect(t *testing.T) {
	die1 := dietest.NewDieFromDieState(die.Skunk)
	die2 := dietest.NewDieFromDieState(die.Skunk)
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)
	playerTurn.Roll()
	playerTurn.Pass()

	assertTurn(t, "double skunk is rolled and a pass action is taken", 0, 4, turn.CompleteResetScore, playerTurn)

	die1 = dietest.NewDieFromDieState(die.Skunk)
	die2 = dietest.NewDieFromDieState(die.Deuce)
	skunkRoll = dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn = turn.NewTurnFromISkunkDice(skunkRoll)
	playerTurn.Roll()
	playerTurn.Pass()

	assertTurn(t, "skunk deuce is rolled and a pass action is taken", 0, 2, turn.Complete, playerTurn)

	die1 = dietest.NewDieFromDieState(die.Skunk)
	die2 = dietest.NewDieFromDieState(die.Five)
	skunkRoll = dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn = turn.NewTurnFromISkunkDice(skunkRoll)
	playerTurn.Roll()
	playerTurn.Pass()

	assertTurn(t, "single skunk is rolled and a pass action is taken", 0, 1, turn.Complete, playerTurn)
}

func Test_when_a_skunk_is_rolled_roll_action_has_no_effect(t *testing.T) {
	die1 := dietest.NewDieFromArray([]int{1, 1, 1, 5})
	die2 := dietest.NewDieFromArray([]int{1, 5, 2, 5})
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)
	for i := 0; i < 4; i++ {
		playerTurn.Roll()
	}

	assertTurn(t, "double skunk is rolled and a roll action is taken", 0, 4, turn.CompleteResetScore, playerTurn)

	die1 = dietest.NewDieFromArray([]int{1, 1, 1, 5})
	die2 = dietest.NewDieFromArray([]int{2, 1, 5, 5})
	skunkRoll = dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn = turn.NewTurnFromISkunkDice(skunkRoll)
	for i := 0; i < 4; i++ {
		playerTurn.Roll()
	}

	assertTurn(t, "skunk deuce is rolled and a roll action is taken", 0, 2, turn.Complete, playerTurn)

	die1 = dietest.NewDieFromArray([]int{1, 1, 1, 5})
	die2 = dietest.NewDieFromArray([]int{5, 1, 2, 5})
	skunkRoll = dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn = turn.NewTurnFromISkunkDice(skunkRoll)
	for i := 0; i < 4; i++ {
		playerTurn.Roll()
	}

	assertTurn(t, "single skunk is rolled and a roll action is taken", 0, 1, turn.Complete, playerTurn)
}
