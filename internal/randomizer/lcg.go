package randomizer

import "time"

type LCG struct {
	state uint64
}

func NewLCG(seed uint64) *LCG {
	if seed == 0 {
		seed = uint64(time.Now().UnixNano())
	}
	return &LCG{state: seed}
}

func (lcg *LCG) Intn(n int) int {
	// Constants for the LCG (example values, you might need to choose more suitable ones)
	const a = 1664525    // Multiplier
	const c = 1013904223 // Increment
	const m = 4294967296 // Modulus (2^32)
	lcg.state = (a*lcg.state + c) % m
	return int(lcg.state % uint64(n))
}
