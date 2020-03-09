package skunk

import (
	"testing"
)

func Test_die_roll_is_always_greater_than_zero_and_less_than_seven(t *testing.T) {
	var die IRollable = Die{}
	var numRolls = 10000
	var counts [7]int
	for i := 0; i < 7; i++ {
		counts[i] = 0
	}

	for i := 0; i < numRolls; i++ {
		die.Roll()
		switch actual := die.GetLastRoll(); actual {
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
