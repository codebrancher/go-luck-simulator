package randomizer

// Randomizer defines the interface for a random number generator
type Randomizer interface {
	Intn(n int) int
}
