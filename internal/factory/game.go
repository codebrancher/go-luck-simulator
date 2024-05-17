package factory

import (
	"fmt"
	"go-slot-machine/internal/display"
	"go-slot-machine/internal/engine/peakpursuit"
	"go-slot-machine/internal/engine/wildfruits"
	"go-slot-machine/internal/manager"
	"go-slot-machine/internal/randomizer"
)

func NewGame(gameType string, rng randomizer.Randomizer, startingCash int, currency string) (manager.Game, error) {
	switch gameType {
	case "1":
		disp := &display.WildFruitsDisplay{}
		g := wildfruits.NewSlotMachine(rng, startingCash, currency, disp)
		m := manager.NewWildFruitsManager(g, disp)
		return m, nil
	case "2":
		disp := &display.PeakPursuitDisplay{}
		g := peakpursuit.NewSlotMachine(rng, startingCash, currency, disp)
		m := manager.NewPeakPursuitManager(g, disp)
		return m, nil
	default:
		return nil, fmt.Errorf("unknown game type: %s", gameType)
	}
}
