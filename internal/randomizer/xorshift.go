package randomizer

import "time"

type XorShiftRNG struct {
	state uint64
}

func NewXorShiftRNG(seed uint64) *XorShiftRNG {
	if seed == 0 {
		seed = uint64(time.Now().UnixNano())
	}
	return &XorShiftRNG{state: seed}
}

func (x *XorShiftRNG) Intn(n int) int {
	// Basic xorshift* algorithm
	x.state ^= x.state >> 12
	x.state ^= x.state << 25
	x.state ^= x.state >> 27
	randVal := x.state * 2685821657736338717
	return int(randVal % uint64(n))
}
