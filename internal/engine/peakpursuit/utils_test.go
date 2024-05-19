package peakpursuit

import (
	"go-luck-simulator/internal/randomizer"
	"testing"
)

// NoOpObserver is a no-operation observer used for testing.
type NoOpObserver struct{}

func (n *NoOpObserver) Update(state *PeakPursuitObserverState) {
	// Do nothing
}

// Helper function to create a SlotMachine with initial conditions
func createTestSlotMachine() *SlotMachine {
	rng := randomizer.NewDefaultRNG()
	noopObserver := &NoOpObserver{}
	machine := NewSlotMachine(rng, 100, "USD", noopObserver)
	return machine
}

// Test when the symbol is not at max level
func TestProcessSymbolLevelIncrease_NotAtMax(t *testing.T) {
	machine := createTestSlotMachine()
	symbolKey := "testSymbol"
	machine.GameConfig.Symbols = []Symbol{{Representation: symbolKey, Payouts: []int{10, 20, 30, 40, 50}}}
	machine.State.SymbolLevels[symbolKey] = SymbolState{CurrentLevel: 2, MaxReached: false}

	machine.processSymbolLevelIncrease(symbolKey)
	newState := machine.State.SymbolLevels[symbolKey]

	if newState.CurrentLevel != 3 {
		t.Errorf("Expected level 3, got %d", newState.CurrentLevel)
	}
}

// Test when the symbol reaches max level
func TestProcessSymbolLevelIncrease_AtMax(t *testing.T) {
	machine := createTestSlotMachine()
	symbolKey := "testSymbol"
	machine.GameConfig.Symbols = []Symbol{{Representation: symbolKey, Payouts: []int{10, 20, 30, 40, 50}}}
	machine.GameConfig.MaxLevel = 5
	machine.State.SymbolLevels[symbolKey] = SymbolState{CurrentLevel: 4, MaxReached: false}

	machine.processSymbolLevelIncrease(symbolKey)
	newState := machine.State.SymbolLevels[symbolKey]

	if !newState.MaxReached || newState.CurrentLevel != 5 {
		t.Errorf("Expected max reached at level 5, got level %d and MaxReached %t", newState.CurrentLevel, newState.MaxReached)
	}
}

// Test when the symbol is already at max level
func TestProcessSymbolLevelIncrease_AlreadyAtMax(t *testing.T) {
	machine := createTestSlotMachine()
	symbolKey := "testSymbol"
	machine.GameConfig.Symbols = []Symbol{{Representation: symbolKey, Payouts: []int{10, 20, 30, 40, 50}}}
	machine.GameConfig.MaxLevel = 5
	machine.State.SymbolLevels[symbolKey] = SymbolState{CurrentLevel: 5, MaxReached: true}

	machine.processSymbolLevelIncrease(symbolKey)
	newState := machine.State.SymbolLevels[symbolKey]

	if newState.CurrentLevel != 5 {
		t.Errorf("Expected to remain at level 5, got %d", newState.CurrentLevel)
	}
}

// Test when the symbol does not exist
func TestProcessSymbolLevelIncrease_SymbolNotExists(t *testing.T) {
	machine := createTestSlotMachine()
	symbolKey := "nonexistentSymbol"

	// This should not panic or alter state since the symbol does not exist
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The function panicked for non-existent symbol")
		}
	}()
	machine.processSymbolLevelIncrease(symbolKey)

	_, exists := machine.State.SymbolLevels[symbolKey]
	if exists {
		t.Errorf("State was altered for a non-existent symbol")
	}
}
