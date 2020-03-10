package die

// ISkunkDie implements an interface for performing a roll action and getting its result.
type ISkunkDie interface {
	Roll()
	GetLastRoll() State
}
