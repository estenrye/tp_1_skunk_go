package turntest

import (
	"testing"

	"github.com/estenrye/skunk/internal/skunk/dice"
	"github.com/estenrye/skunk/internal/skunk/die"
	"github.com/estenrye/skunk/internal/skunk/turn"
	"github.com/estenrye/skunk/internal/skunk_test/dietest"
)

func Test_when_a_single_skunk_is_rolled_score_is_zero_penalty_is_one_and_state_is_complete(t *testing.T) {
	for i := 3; i <= 6; i++ {
		die1 := dietest.NewDieFromDieState(die.Skunk)
		die2 := dietest.NewDieFromInt(i)
		skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
		playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)
		playerTurn.Roll()

		if playerTurn.GetScore() != 0 {
			t.Errorf("When a single skunk is rolled, expected score to be 0, got %d", playerTurn.GetScore())
		}

		if playerTurn.GetPenalty() != 1 {
			t.Errorf("When a single skunk is rolled, expected penalty to be 1, got %d", playerTurn.GetPenalty())
		}

		if playerTurn.GetState() != turn.Complete {
			t.Errorf("When a single skunk is rolled, expected state to be Complete, got %s", playerTurn.GetState())
		}
	}
}

func Test_when_a_skunk_deuce_is_rolled_score_is_zero_penalty_is_two_and_state_is_complete(t *testing.T) {
	die1 := dietest.NewDieFromDieState(die.Skunk)
	die2 := dietest.NewDieFromDieState(die.Deuce)
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)
	playerTurn.Roll()

	if playerTurn.GetScore() != 0 {
		t.Errorf("When a skunk deuce is rolled, expected score to be 0, got %d", playerTurn.GetScore())
	}

	if playerTurn.GetPenalty() != 2 {
		t.Errorf("When a skunk deuce is rolled, expected penalty to be 2, got %d", playerTurn.GetPenalty())
	}

	if playerTurn.GetState() != turn.Complete {
		t.Errorf("When a skunk deuce is rolled, expected state to be Complete, got %s", playerTurn.GetState())
	}
}

func Test_when_a_double_skunk_is_rolled_score_is_zero_penalty_is_four_and_state_is_completeResetScore(t *testing.T) {
	die1 := dietest.NewDieFromDieState(die.Skunk)
	die2 := dietest.NewDieFromDieState(die.Skunk)
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
	playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)
	playerTurn.Roll()

	if playerTurn.GetScore() != 0 {
		t.Errorf("When a double skunk is rolled, expected score to be 0, got %d", playerTurn.GetScore())
	}

	if playerTurn.GetPenalty() != 4 {
		t.Errorf("When a double skunk is rolled, expected penalty to be 4, got %d", playerTurn.GetPenalty())
	}

	if playerTurn.GetState() != turn.CompleteResetScore {
		t.Errorf("When a double skunk is rolled, expected state to be Complete Reset Score, got %s", playerTurn.GetState())
	}
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

			if playerTurn.GetScore() != 2*(i+j) {
				t.Errorf("When a %s and a %s are rolled twice, expected score to be %d, got %d", die.State(i), die.State(j), 2*(i+j), playerTurn.GetScore())
			}

			if playerTurn.GetPenalty() != 0 {
				t.Errorf("When a %s and a %s are rolled twice, expected penalty to be 0, got %d", die.State(i), die.State(j), playerTurn.GetPenalty())
			}

			if playerTurn.GetState() != turn.Active {
				t.Errorf("When a %s and a %s are is rolled twice, expected state to be Active, got %s", die.State(i), die.State(j), playerTurn.GetState())
			}
		}
	}
}
