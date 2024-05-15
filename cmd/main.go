package main

import (
	"go-slot-machine/internal/display"
	"go-slot-machine/internal/engine/wildfruits"
	"go-slot-machine/internal/manager"
	"go-slot-machine/internal/randomizer"
	"time"
)

func main() {
	startingCash := 250                   // modify the starting cash, for simulations with 100000 spins you can set it 100000 to get more granularity
	visualDelay := 100 * time.Millisecond // refreshing rate of the symbols, for testing/simulations you might want to set it to 0
	autoDelay := 1 * time.Second          // delay between the auto spins, for testing/simulations you might want to set it to 0
	runSim := false                       // set to true to run the simulation

	seed := uint64(time.Now().UnixNano())
	rng := randomizer.NewXorShiftRNG(seed) // switch the randomizer for observing different outcomes

	game := wildfruits.NewSlotMachine(rng, visualDelay, startingCash)

	disp := &display.ConsoleDisplay{}
	game.RegisterObserver(disp) // I suggest to comment this line out for simulations since it produces the console output and slows down simulations

	m := manager.GameManager{
		Game:      game,
		Display:   disp,
		AutoDelay: autoDelay,
	}

	if runSim {
		totalSpins := 100000
		betAmount := 1
		m.RunSimulation(totalSpins, betAmount)
	} else {
		m.Play()
	}
}
