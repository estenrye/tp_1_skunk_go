package turn

// ISkunkTurn provides an interface to interact with a turn object
type ISkunkTurn interface {
	Roll()
	Pass()
	GetScore() int
	GetPenalty() int
	GetState() State
	GetLastRoll() string
}
