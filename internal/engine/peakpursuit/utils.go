package peakpursuit

import (
	"fmt"
	"go-luck-simulator/internal/utils"
	"time"
)

func (s *SlotMachine) SetVisualDelay(delay time.Duration) {
	s.GameConfig.VisualDelay = delay
}

// handleWildSymbol processes the effect of a wild symbol spin.
func (s *SlotMachine) handleWildSymbol() {
	for symbolKey, _ := range s.State.SymbolLevels {
		if s.isSpecialSymbol(symbolKey) {
			continue
		}
		s.processSymbolLevelIncrease(symbolKey)
	}
}

// isSpecialSymbol checks if the symbol is special and should be ignored in certain operations.
func (s *SlotMachine) isSpecialSymbol(symbol string) bool {
	return symbol == s.GameConfig.ResetSymbolRepresentation || symbol == s.GameConfig.WildSymbolRepresentation
}

// processSymbolLevelIncrease increases the level of a symbol if not at maximum and processes payouts.
func (s *SlotMachine) processSymbolLevelIncrease(symbolKey string) {
	state := s.State.SymbolLevels[symbolKey]
	if state.CurrentLevel >= s.GameConfig.MaxLevel {
		s.State.CurrentMaxLevelStreak++
		s.State.MaxLevelReached++
		s.processPayouts(symbolKey, state)
		return
	}

	state.CurrentLevel++
	s.State.SymbolLevels[symbolKey] = state

	symbol, exists := findSymbol(symbolKey)
	if !exists || len(symbol.Payouts) <= state.CurrentLevel-1 {
		return
	}

	s.processPayouts(symbolKey, state)
}

// processPayouts calculates and applies payouts based on the symbol level and bet.
func (s *SlotMachine) processPayouts(symbolKey string, state SymbolState) {
	symbol, exists := findSymbol(symbolKey)
	if !exists || len(symbol.Payouts) <= state.CurrentLevel-1 {
		return
	}

	payoutMultiplier := symbol.Payouts[state.CurrentLevel-1]
	payout := s.State.LastBet * payoutMultiplier

	// Check if the symbol has reached the maximum level
	if state.CurrentLevel == s.GameConfig.MaxLevel {
		if state.MaxReached {
			// If max level was already reached, handle instant payout
			s.State.InstantWinCounter[symbolKey]++
			s.applyPayout(symbolKey, payout, true) // Instant payout
		} else {
			// First time reaching max level, mark as reached and add to pot
			state.MaxReached = true
			s.State.SymbolLevels[symbolKey] = state
			s.applyPayout(symbolKey, payout, false) // Add to pot
		}
	} else {
		// Not reached max level yet, add to pot
		s.applyPayout(symbolKey, payout, false) // Add to pot
	}
}

// applyPayout applies the calculated payout and updates the game state based on the context.
func (s *SlotMachine) applyPayout(symbolKey string, payout int, instant bool) {
	if instant {
		s.State.Amount += payout
		s.State.Wins = utils.RecordWin(payout, s.State.Wins)
		s.State.TotalWins++
		s.State.TotalWinAmount += payout
		if s.DisplayConfig.WinningDescription != "" {
			s.DisplayConfig.WinningDescription += "\n"
		}
		s.DisplayConfig.WinningDescription += fmt.Sprintf("%s instant payout: %d%s %s", symbolKey, payout, s.GameConfig.Currency, symbolKey)
	} else {
		s.State.Winnings += payout
	}
	s.updateTopWinAmount(payout)
}

// updateTopWinAmount updates the top win amount if the current payout exceeds it.
func (s *SlotMachine) updateTopWinAmount(payout int) {
	if payout > s.State.TopWinAmount {
		s.State.TopWinAmount = payout
	}
}

func (s *SlotMachine) CashOut() {
	s.State.Amount += s.State.Winnings
	s.State.TotalWinAmount += s.State.Winnings
	s.State.Wins = utils.RecordWin(s.State.Winnings, s.State.Wins)
	s.updateTopWinAmount(s.State.Winnings)
	s.State.TotalWins++
	s.State.Winnings = 0
	s.resetAllSymbols()
	s.DisplayConfig.WinningDescription = "💰 Cash out 💰"
}

// hasActivePot returns true if any symbol has a level above the initial level.
func (s *SlotMachine) hasActivePot() bool {
	for _, state := range s.State.SymbolLevels {
		if state.CurrentLevel > 0 {
			return true
		}
	}
	return false
}

func (s *SlotMachine) spinWheel() {
	weightedSymbols := generateWeightedSymbols()
	var index int

	for spin := 0; spin < 30; spin++ {
		index = s.RNG.Intn(len(weightedSymbols))
		s.GameConfig.SpinSymbolRepresentation = weightedSymbols[index]
		s.NotifyObservers()
		time.Sleep(s.GameConfig.VisualDelay)
	}
	s.GameConfig.SpinSymbolRepresentation = weightedSymbols[index]
	s.NotifyObservers()
	s.State.TotalSpins++
	s.State.CurrentSpinStreak++
}

func generateWeightedSymbols() []string {
	var weightedSymbols []string
	for _, sym := range Symbols {
		for i := 0; i < sym.Frequency; i++ {
			weightedSymbols = append(weightedSymbols, sym.Representation)
		}
	}
	return weightedSymbols
}

func findSymbol(representation string) (Symbol, bool) {
	for _, sym := range Symbols {
		if sym.Representation == representation {
			return sym, true
		}
	}
	return Symbol{}, false
}

func (s *SlotMachine) resetAllSymbols() {
	for key := range s.State.SymbolLevels {
		s.State.SymbolLevels[key] = SymbolState{0, false}
	}
	s.State.LongestSpinStreak = utils.UpdateLongestSpinStreak(s.State.CurrentSpinStreak, s.State.LongestSpinStreak)
	s.State.LongestMaxLevelStreak = utils.UpdateMaxLevelStreak(s.State.CurrentMaxLevelStreak, s.State.LongestMaxLevelStreak)
	s.State.CurrentMaxLevelStreak = 0
	s.State.CurrentSpinStreak = 0
	s.State.Winnings = 0
	s.DisplayConfig.WinningDescription = "👽 Game state reset 👽"
}
