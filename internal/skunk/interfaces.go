package skunk

// IRollable implements an interface for performing a roll action and getting its result.
type IRollable interface {
	Roll()
	GetLastRoll() int
}
