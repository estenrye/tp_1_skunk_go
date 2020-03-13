package turntest

import (
	"testing"

	"github.com/estenrye/skunk/internal/skunk/dice"
	"github.com/estenrye/skunk/internal/skunk/die"
	"github.com/estenrye/skunk/internal/skunk/turn"
	"github.com/estenrye/skunk/internal/skunk_test/dietest"
)

func Test_when_a_single_skunk_is_rolled_score_is_zero_and_state_is_complete(t *testing.T) {
	for i := 3; i <= 6; i++ {
		die1 := dietest.NewDieFromDieState(die.Skunk)
		die2 := dietest.NewDieFromInt(i)
		skunkRoll := dice.NewDiceFromISkunkDice(die1, die2)
		playerTurn := turn.NewTurnFromISkunkDice(skunkRoll)
		playerTurn.Roll()

		if playerTurn.GetScore() != 0 {
			t.Errorf("When a single skunk is rolled, expected score to be zero and state to be Complete, got %s, %d", playerTurn.GetState(), playerTurn.GetScore())
		}
	}
}
