package factory

import (
	"fmt"
	"go-luck-simulator/internal/display"
	"go-luck-simulator/internal/engine/peakpursuit"
	"go-luck-simulator/internal/engine/wildfruits"
	"go-luck-simulator/internal/manager"
	"go-luck-simulator/internal/randomizer"
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
