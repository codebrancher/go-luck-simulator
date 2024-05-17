package factory

import (
	"fmt"
	"go-slot-machine/internal/randomizer"
	"time"
)

func NewRandomizer(r string) (randomizer.Randomizer, error) {
	seed := uint64(time.Now().UnixNano())
	switch r {
	case "1":
		rng := randomizer.NewDefaultRNG()
		return rng, nil
	case "2":
		rng := randomizer.NewLCG(seed)
		return rng, nil
	case "3":
		rng := randomizer.NewXorShiftRNG(seed) // switch the randomizer for observing different outcomes
		return rng, nil
	default:
		return nil, fmt.Errorf("unknown randomizer type: %s", r)
	}
}
