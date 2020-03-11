package dicetest

import (
	"testing"

	"github.com/estenrye/skunk/internal/skunk/dice"
	"github.com/estenrye/skunk/internal/skunk/die"
	"github.com/estenrye/skunk/internal/skunk_test/dietest"
)

func Test_when_two_skunks_are_rolled_score_is_zero_and_state_is_double_skunk(t *testing.T) {
	die1 := dietest.NewDieFromDieState(die.Skunk)
	die2 := dietest.NewDieFromDieState(die.Skunk)
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
	skunkRoll.Roll()
	actualRoll := skunkRoll.GetLastRoll()
	if actualRoll != 0 {
		t.Errorf("When two skunks are rolled, expected a score of 0, got %d.", actualRoll)
	}

	actualState := skunkRoll.GetLastState()
	if actualState != dice.DoubleSkunk {
		t.Errorf("When two skunks are rolled, expected a state of %s, got %s.", dice.DoubleSkunk, actualState)
	}

	if actualState.GetPenalty() != 4 {
		t.Errorf("When two skunks are rolled, expected a penalty of 4, got %d.", actualState.GetPenalty())
	}
}

func Test_when_a_skunk_and_deuce_are_rolled_score_is_zero_and_state_is_skunk_deuce(t *testing.T) {
	die1 := dietest.NewDieFromArray([]int{1, 2})
	die2 := dietest.NewDieFromArray([]int{2, 1})
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)

	for i := 0; i < 2; i++ {
		skunkRoll.Roll()
		actualRoll := skunkRoll.GetLastRoll()
		if actualRoll != 0 {
			t.Errorf("When a %s and a %s are rolled, expected a score of 0, got %d.", skunkRoll.GetLastDie1(), skunkRoll.GetLastDie2(), actualRoll)
		}

		actualState := skunkRoll.GetLastState()
		if actualState != dice.SkunkDeuce {
			t.Errorf("When a %s and a %s are rolled, expected a state of %s, got %s.", skunkRoll.GetLastDie1(), skunkRoll.GetLastDie2(), dice.SkunkDeuce, actualState)
		}

		if actualState.GetPenalty() != 2 {
			t.Errorf("When a %s and a %s are rolled, expected a penalty of 2, got %d", skunkRoll.GetLastDie1(), skunkRoll.GetLastDie2(), actualState.GetPenalty())
		}

	}
}

func Test_when_a_skunk_and_a_three_or_greater_is_rolled(t *testing.T) {
	die1 := dietest.NewDieFromArray([]int{1, 1, 1, 1, 3, 4, 5, 6})
	die2 := dietest.NewDieFromArray([]int{3, 4, 5, 6, 1, 1, 1, 1})
	for i := 0; i < 8; i++ {
		die1.Roll()
		die2.Roll()
		skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
		actualRoll := skunkRoll.GetLastRoll()
		actualState := skunkRoll.GetLastState()

		if actualRoll != 0 {
			t.Errorf("When a %s and a %s are rolled, expected a score of 0, got %d.", skunkRoll.GetLastDie1(), skunkRoll.GetLastDie2(), actualRoll)
		}

		if actualState != dice.SingleSkunk {
			t.Errorf("When a %s and a %s are rolled, expected a state of %s, got %s.", skunkRoll.GetLastDie1(), skunkRoll.GetLastDie2(), dice.SingleSkunk, actualState)
		}

		if actualState.GetPenalty() != 1 {
			t.Errorf("When a %s and a %s are rolled, expected a penalty of 1, got %d", skunkRoll.GetLastDie1(), skunkRoll.GetLastDie2(), actualState.GetPenalty())
		}
	}
}

func Test_when_a_skunk_is_not_rolled_state_is_scorable_roll_and_last_roll_is_sum_of_dice(t *testing.T) {

	for i := 2; i <= 6; i++ {
		for j := 2; j <= 6; j++ {
			die1 := dietest.NewDieFromArray([]int{i, j})
			die2 := dietest.NewDieFromArray([]int{j, i})
			expectedRoll := i + j
			skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)

			for k := 0; k < 2; k++ {
				skunkRoll.Roll()
				actualRoll := skunkRoll.GetLastRoll()
				actualState := skunkRoll.GetLastState()

				if actualRoll != expectedRoll {
					t.Errorf("When a %s and a %s are rolled, expected a score of %d, got %d.", skunkRoll.GetLastDie1(), skunkRoll.GetLastDie2(), expectedRoll, actualRoll)
				}

				if actualState != dice.ScorableRoll {
					t.Errorf("When a %s and a %s are rolled, expected a state of %s, got %s.", skunkRoll.GetLastDie1(), skunkRoll.GetLastDie2(), dice.ScorableRoll, actualState)
				}

				if actualState.GetPenalty() != 0 {
					t.Errorf("When a %s and a %s are rolled, expected a penalty of 0, got %d", skunkRoll.GetLastDie1(), skunkRoll.GetLastDie2(), actualState.GetPenalty())
				}
			}
		}
	}
}

func Test_when_GetLastDieN_is_called_it_returns_expected_value(t *testing.T) {
	die1 := dietest.NewDieFromDieState(die.Five)
	die2 := dietest.NewDieFromDieState(die.Six)
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)

	skunkRoll.Roll()
	if skunkRoll.GetLastDie1() != die.Five {
		t.Errorf("When GetLastDie1 is called, expected %s, got %s", die.Five, skunkRoll.GetLastDie1())
	}

	if skunkRoll.GetLastDie2() != die.Six {
		t.Errorf("When GetLastDie2 is called, expected %s, got %s.", die.Six, skunkRoll.GetLastDie2())
	}
}

func Test_when_GetLastRoll_is_called_before_Roll_value_is_zero_and_states_are_unknown(t *testing.T) {
	die1 := dietest.NewDieFromDieState(die.Three)
	die2 := dietest.NewDieFromDieState(die.Four)
	skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)

	if skunkRoll.GetLastRoll() != 0 {
		t.Errorf("When GetLastRoll is called before Roll, expected 0, got %d", skunkRoll.GetLastRoll())
	}

	if skunkRoll.GetLastState() != dice.UnknownState {
		t.Errorf("When GetLastState is called before Roll, expected %s, got %s", dice.UnknownState, skunkRoll.GetLastState())
	}

	if skunkRoll.GetLastDie1() != die.UnknownState {
		t.Errorf("When GetLastDie1 is called before Roll, expected %s, got %s", die.UnknownState, skunkRoll.GetLastDie1())
	}

	if skunkRoll.GetLastDie2() != die.UnknownState {
		t.Errorf("When GetLastDie2 is called before Roll, expected %s, got %s", die.UnknownState, skunkRoll.GetLastDie2())
	}
}
