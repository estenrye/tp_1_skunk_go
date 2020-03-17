package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/estenrye/skunk/internal/skunk/player"
)

func getGreeting() string {
	return "Hello World!"
}

func getTurnInput(p player.ISkunkPlayer, reader *bufio.Reader) {
	fmt.Printf("%s: Current Turn Score: %d.  Would you like to roll or pass? (r|p): ", p.GetName(), p.GetLastTurn().GetScore())
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		switch text {
		case "r":
			p.Roll()
			r := p.GetLastTurn().GetLastRoll()
			fmt.Printf("%s: Rolled a %s and a %s.\n", p.GetName(), r.GetLastDie1(), r.GetLastDie2())
			break
		case "p":
			p.Pass()
			break
		default:
			break
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	p := player.NewPlayer("Player 1")

	for p.GetLastState() != player.CompleteEndgame && p.GetLastChips() > 0 {
		fmt.Printf("%s: %d points, %d chips\n", p.GetName(), p.GetLastScore(), p.GetLastChips())
		p.NewTurn()
		fmt.Printf("It's %s's turn.\n", p.GetName())

		for p.GetLastState() != player.CompleteTurn && p.GetLastState() != player.CompleteEndgame {
			getTurnInput(p, reader)
		}
	}
	fmt.Printf("%s: %d points, %d chips\n", p.GetName(), p.GetLastScore(), p.GetLastChips())

}
