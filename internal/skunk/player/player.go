package player

import "github.com/estenrye/skunk/internal/skunk/turn"

// Player represents the concept of a human player playing the game of skunk
type Player struct {
	name  string
	score int
	chips int
	turn  turn.ISkunkTurn
	state State
}

// NewPlayer initializes a human player.
func NewPlayer(name string) ISkunkPlayer {
	return &Player{
		name:  name,
		score: 0,
		chips: 50,
		state: TurnNotStarted,
	}
}

// NewTurn creates a new turn for the player.
func (p *Player) NewTurn() {
	p.NewTurnFromISkunkTurn(turn.NewTurn())
}

// NewTurnFromISkunkTurn creates a new turn for a player from a known turn.
func (p *Player) NewTurnFromISkunkTurn(turn turn.ISkunkTurn) {
	p.turn = turn
	p.state = TurnNotStarted
}

// Roll performs the roll action for the player.
func (p *Player) Roll() {
	if p.turn == nil {
		return
	}
	p.turn.Roll()

	if p.turn.GetState() == turn.Active {
		p.state = ActiveTurn
	}

	if p.turn.GetState() == turn.Complete {
		p.state = CompleteTurn
		p.score += p.turn.GetScore()
		p.chips -= p.turn.GetPenalty()
	}

	if p.turn.GetState() == turn.CompleteResetScore {
		p.state = CompleteTurn
		p.score = 0
		p.chips -= p.turn.GetPenalty()
	}
}

// Pass performs the pass action for the player.
func (p *Player) Pass() {
	if p.turn == nil {
		return
	}
	p.turn.Pass()
	p.score += p.turn.GetScore()

	if p.score >= 100 {
		p.state = CompleteEndgame
	} else {
		p.state = CompleteTurn
	}
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

// GetLastState returns tthe player state as of the last action taken.
func (p Player) GetLastState() State {
	return p.state
}
