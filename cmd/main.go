package main

import (
	"go-slot-machine/internal/display"
	"go-slot-machine/internal/engine/wildfruits"
	"go-slot-machine/internal/manager"
	"go-slot-machine/internal/randomizer"
	"time"
)

func main() {
	startingCash := 100000
	visualDelay := 0 * time.Millisecond
	autoDelay := 0 * time.Second
	seed := uint64(time.Now().UnixNano())

	// Seed the LCG with the current Unix timestamp
	rng := randomizer.NewXorShiftRNG(seed)
	game := wildfruits.NewSlotMachine(rng, visualDelay, startingCash)
	disp := &display.ConsoleDisplay{}
	//game.RegisterObserver(disp)
	m := manager.GameManager{
		Game:      game,
		Display:   disp,
		AutoDelay: autoDelay,
	}

	//m.Play()

	// Run simulation
	totalSpins := 100000
	betAmount := 1
	m.AutomatedPlay(totalSpins, betAmount)

	// Calculate and display results
	disp.ShowStats(game, startingCash)
}
