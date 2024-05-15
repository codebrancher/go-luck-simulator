package randomizer

import (
	"math/rand"
	"time"
)

type DefaultRNG struct {
	r *rand.Rand
}

func NewDefaultRNG() *DefaultRNG {
	src := rand.NewSource(time.Now().UnixNano())
	return &DefaultRNG{r: rand.New(src)}
}

func (d *DefaultRNG) Intn(n int) int {
	return d.r.Intn(n)
}
