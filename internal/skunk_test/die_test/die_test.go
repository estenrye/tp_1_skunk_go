package dietest

import (
	"testing"

	"github.com/estenrye/skunk/internal/skunk/die"
)

func Test_mockdie_is_predictable_for_a_single_roll(t *testing.T) {
	for i := 1; i <= 6; i++ {
		d1 := NewDieFromInt(i)
		d1.Roll()
		if d1.GetLastRoll() != i {
			t.Errorf("MockDie is not reliable, expected %d but got %d", i, d1.GetLastRoll())
		}
	}
}

func Test_mockdie_is_predictable_for_multiple_rolls(t *testing.T) {
	for i := 1; i <= 6; i++ {
		d1 := NewDieFromInt(i)
		for j := 0; j < 10; j++ {
			d1.Roll()
			if d1.GetLastRoll() != i {
				t.Errorf("MockDie is not reliable, expected %d but got %d", i, d1.GetLastRoll())
			}
		}
	}
}

func Test_mockdie_wraps_around_a_sequence_predictably_for_multiple_rolls(t *testing.T) {
	d1 := NewDieFromArray([]int{1, 2, 3, 4, 5, 6})
	for j := 1; j <= 3; j++ {
		for i := 1; i <= 6; i++ {
			d1.Roll()
			if d1.GetLastRoll() != i {
				t.Errorf("MockDie is not reliable, expected %d but got %d", i, d1.GetLastRoll())
			}
		}
	}
}

func Test_die_roll_is_always_greater_than_zero_and_less_than_seven(t *testing.T) {
	d1 := die.NewDie()
	var numRolls = 100000
	var counts [7]int
	for i := 0; i < 7; i++ {
		counts[i] = 0
	}

	for i := 0; i < numRolls; i++ {
		d1.Roll()
		switch actual := d1.GetLastRoll(); actual {
		case 1:
			counts[1]++
			break
		case 2:
			counts[2]++
			break
		case 3:
			counts[3]++
			break
		case 4:
			counts[4]++
			break
		case 5:
			counts[5]++
			break
		case 6:
			counts[6]++
			break
		default:
			counts[0]++
			break
		}
	}

	if counts[0] > 0 {
		t.Errorf("Die fails to return a value between 1 and 6.  Occurs %d out of %d rolls", counts[0], numRolls)
	}

	for i := 1; i < 7; i++ {
		if counts[i] < numRolls/7 || counts[i] > numRolls/5 {
			t.Errorf("Die fails to meet pseudorandom results for outcome %d.  Outcome occurs %d out of %d rolls.",
				i, counts[i], numRolls)
		}
	}
}