package main

import (
	"go-slot-machine/internal/display"
	"go-slot-machine/internal/engine/wildfruits"
	"go-slot-machine/internal/manager"
	"go-slot-machine/internal/randomizer"
	"time"
)

func main() {
	currency := ""
	startingCash := 250 // modify the starting cash, for simulations with 100000 spins you can set it 100000 to get more granularity

	seed := uint64(time.Now().UnixNano())
	rng := randomizer.NewXorShiftRNG(seed) // switch the randomizer for observing different outcomes

	disp := &display.ConsoleDisplay{}
	game := wildfruits.NewSlotMachine(rng, startingCash, currency, disp)

	m := manager.NewGameManager(game, disp)

	m.Start()
}
