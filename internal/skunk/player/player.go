package player

import "github.com/estenrye/skunk/internal/skunk/turn"

// Player represents the concept of a human player playing the game of skunk
type Player struct {
	name  string
	score int
	chips int
	turn  turn.ISkunkTurn
}

// NewPlayer initializes a human player.
func NewPlayer(name string) ISkunkPlayer {
	return &Player{
		name:  name,
		score: 0,
		chips: 50,
	}
}

// NewTurn creates a new turn for the player.
func (p *Player) NewTurn() {
	p.NewTurnFromISkunkTurn(turn.NewTurn())
}

// NewTurnFromISkunkTurn creates a new turn for a player from a known turn.
func (p *Player) NewTurnFromISkunkTurn(turn turn.ISkunkTurn) {
	p.turn = turn
}

// Roll performs the roll action for the player.
func (p *Player) Roll() {
	p.turn.Roll()
}

// Pass performs the pass action for the player.
func (p *Player) Pass() {
	p.turn.Pass()
}

// GetName returns the player's name
func (p Player) GetName() string {
	return p.name
}

// GetLastScore returns the player's score as of the last action taken.
func (p Player) GetLastScore() int {
	return p.score
}

// GetLastChips returns the player's chips as of the last action taken.
func (p Player) GetLastChips() int {
	return p.chips
}

// GetLastTurn returns a data representation of the player's turn object.
func (p Player) GetLastTurn() turn.ISkunkTurnResult {
	return p.turn
}
