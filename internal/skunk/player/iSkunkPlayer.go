package player

// ISkunkPlayer represents a
type ISkunkPlayer interface {
	Roll()
	Pass()
	GetScore() int
	// GetState() State
	GetLastRoll() string
}
