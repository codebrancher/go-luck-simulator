package peakpursuit

import (
	"fmt"
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
	machine := NewSlotMachine(rng, 100, "", noopObserver)
	return machine
}

// Test when the symbol is not at max level
func TestProcessSymbolLevelIncrease_NotAtMax(t *testing.T) {
	machine := createTestSlotMachine()
	symbolKey := "testSymbol"
	machine.GameConfig.Symbols = []Symbol{{Representation: symbolKey, Payouts: []int{10, 20, 30, 40, 50}}}
	machine.State.SymbolLevels[symbolKey] = SymbolState{CurrentLevel: 2, MaxAlreadyReached: false}

	machine.processSymbolLevelIncrease(symbolKey)
	newState := machine.State.SymbolLevels[symbolKey]

	if newState.CurrentLevel != 3 {
		t.Errorf("Expected level 3, got %d", newState.CurrentLevel)
	}
}

func TestProcessSymbolLevelIncrease_FirstAtMax(t *testing.T) {
	machine := createTestSlotMachine()
	symbolKey := "testSymbol"
	machine.GameConfig.Symbols = []Symbol{{Representation: symbolKey, Payouts: []int{10, 20, 30, 40, 50}}}
	machine.GameConfig.MaxLevel = 5
	machine.State.SymbolLevels[symbolKey] = SymbolState{CurrentLevel: 4, MaxAlreadyReached: false}

	machine.processSymbolLevelIncrease(symbolKey)
	newState := machine.State.SymbolLevels[symbolKey]

	if newState.MaxAlreadyReached || newState.CurrentLevel != 5 {
		t.Errorf("Expected max reached at level 5, got level %d and MaxAlreadyReached %t", newState.CurrentLevel, newState.MaxAlreadyReached)
	}
}

// Test when the symbol is already at max level
func TestProcessSymbolLevelIncrease_AlreadyAtMax(t *testing.T) {
	machine := createTestSlotMachine()
	symbolKey := "testSymbol"
	machine.GameConfig.Symbols = []Symbol{{Representation: symbolKey, Payouts: []int{10, 20, 30, 40, 50}}}
	machine.GameConfig.MaxLevel = 5
	machine.State.SymbolLevels[symbolKey] = SymbolState{CurrentLevel: 5, MaxAlreadyReached: true}

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

// Helper function to prepare a slot machine with initial conditions for testing
func prepareSlotMachineForPayoutTest() *SlotMachine {
	rng := randomizer.NewDefaultRNG()
	machine := NewSlotMachine(rng, 100, "", &NoOpObserver{})
	machine.State.LastBet = 10

	// Define some symbols with payouts
	machine.GameConfig.Symbols = []Symbol{
		{"💫", []int{1, 2, 4, 7, 10}, 90},
		{"🪐", []int{3, 8, 16, 29, 50}, 12},
	}
	return machine
}

// TestProcessPayouts_NormalLevel
func TestProcessPayouts_NormalLevel(t *testing.T) {
	machine := prepareSlotMachineForPayoutTest()
	symbolKey := "💫"
	state := SymbolState{CurrentLevel: 1, MaxAlreadyReached: false}

	machine.processPayouts(symbolKey, state)

	// Expected payout calculation
	expectedPayout := 10 * 1
	if machine.State.Winnings != expectedPayout {
		t.Errorf("Expected winnings of %d, got %d", expectedPayout, machine.State.Winnings)
	}
}

// TestProcessPayouts_MaxLevelReached
func TestProcessPayouts_FirstMaxLevelReached(t *testing.T) {
	machine := prepareSlotMachineForPayoutTest()
	symbolKey := "🪐"
	state := SymbolState{CurrentLevel: 5, MaxAlreadyReached: false} // Max level for this symbol

	machine.processPayouts(symbolKey, state)

	if !machine.State.SymbolLevels[symbolKey].MaxAlreadyReached {
		t.Errorf("MaxAlreadyReached should be set to true")
	}
	expectedPayout := 10 * 50
	if machine.State.Winnings != expectedPayout {
		t.Errorf("Expected winnings of %d, got %d", expectedPayout, machine.State.Winnings)
	}
}

// TestProcessPayouts_SymbolNotFound
func TestProcessPayouts_SymbolNotFound(t *testing.T) {
	machine := prepareSlotMachineForPayoutTest()
	symbolKey := "nonexistentSymbol"
	state := SymbolState{CurrentLevel: 1, MaxAlreadyReached: false}

	machine.processPayouts(symbolKey, state)

	// Expecting no change since the symbol does not exist
	if machine.State.Winnings != 0 {
		t.Errorf("Winnings should be 0 for nonexistent symbol, got %d", machine.State.Winnings)
	}
}

func prepareSlotMachineForInstantPayoutTest() *SlotMachine {
	rng := randomizer.NewDefaultRNG()
	machine := NewSlotMachine(rng, 100, "USD", &NoOpObserver{})
	machine.State.TopWinAmount = 40 // Initial top win amount for comparison
	return machine
}

// TestApplyPayoutInstant
func TestApplyPayoutInstant(t *testing.T) {
	machine := prepareSlotMachineForInstantPayoutTest()
	payout := 50

	machine.applyPayout("testSymbol", payout, true)

	if machine.State.Amount != 150 { // Initial 100 + 50 payout
		t.Errorf("Expected amount after instant payout to be 150, got %d", machine.State.Amount)
	}
	if len(machine.State.Wins) != 1 || machine.State.Wins[0] != payout {
		t.Errorf("Expected wins to record %d, got %v", payout, machine.State.Wins)
	}
	if machine.State.TotalWins != 1 {
		t.Errorf("Expected total wins to be 1, got %d", machine.State.TotalWins)
	}
	if machine.State.TotalWinAmount != 50 {
		t.Errorf("Expected total win amount to be 50, got %d", machine.State.TotalWinAmount)
	}
	expectedDescription := fmt.Sprintf("testSymbol instant payout: %dUSD testSymbol", payout)
	if machine.DisplayConfig.WinningDescription != expectedDescription {
		t.Errorf("Expected display description '%s', got '%s'", expectedDescription, machine.DisplayConfig.WinningDescription)
	}
}

// TestApplyPayoutNonInstant
func TestApplyPayoutNonInstant(t *testing.T) {
	machine := prepareSlotMachineForInstantPayoutTest()
	payout := 30

	machine.applyPayout("testSymbol", payout, false)

	if machine.State.Winnings != 30 {
		t.Errorf("Expected deferred winnings to be 30, got %d", machine.State.Winnings)
	}
}

// TestApplyPayoutTopWinUpdate
func TestApplyPayoutTopWinUpdate(t *testing.T) {
	machine := prepareSlotMachineForInstantPayoutTest()
	payout := 50 // Higher than the initial top win amount of 40

	machine.applyPayout("testSymbol", payout, false) // Testing with non-instant to focus on top win update

	if machine.State.TopWinAmount != payout {
		t.Errorf("Expected top win amount to be updated to %d, got %d", payout, machine.State.TopWinAmount)
	}
}

func prepareSlotMachineWithWinnings() *SlotMachine {
	rng := randomizer.NewDefaultRNG()
	machine := NewSlotMachine(rng, 100, "", &NoOpObserver{})
	machine.State.Winnings = 50
	machine.State.Amount = 100
	machine.State.TopWinAmount = 40 // less than current winnings to test if top win updates
	return machine
}

// TestCashOut
func TestCashOut(t *testing.T) {
	machine := prepareSlotMachineWithWinnings()

	machine.CashOut()

	if machine.State.Amount != 150 { // 100 initial + 50 winnings
		t.Errorf("Expected amount after cashout to be 150, got %d", machine.State.Amount)
	}

	if machine.State.Winnings != 0 {
		t.Errorf("Expected winnings to be reset to 0, got %d", machine.State.Winnings)
	}

	if machine.State.TopWinAmount != 50 { // New top win should be the recent winnings of 50
		t.Errorf("Expected top win amount to be updated to 50, got %d", machine.State.TopWinAmount)
	}

	if machine.State.TotalWins != 1 { // Should increment total wins
		t.Errorf("Expected total wins to be 1, got %d", machine.State.TotalWins)
	}

	// Check if symbols are reset
	for _, symbolState := range machine.State.SymbolLevels {
		if symbolState.CurrentLevel != 0 || symbolState.MaxAlreadyReached {
			t.Errorf("Expected all symbols to be reset, found one that wasn't")
		}
	}

	expectedDescription := "💰 Cash out 💰"
	if machine.DisplayConfig.WinningDescription != expectedDescription {
		t.Errorf("Expected display description to be '%s', got '%s'", expectedDescription, machine.DisplayConfig.WinningDescription)
	}
}
