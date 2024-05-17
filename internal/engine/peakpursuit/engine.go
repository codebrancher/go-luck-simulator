package peakpursuit

import (
	"fmt"
	"go-slot-machine/internal/engine"
	"go-slot-machine/internal/randomizer"
	"go-slot-machine/internal/utils"
	"strconv"
	"strings"
	"time"
)

type SlotMachine struct {
	Title         string
	RNG           randomizer.Randomizer
	GameConfig    GameConfig
	GameState     GameState
	DisplayConfig DisplayConfig
	observers     []engine.Observer[*PeakPursuitState]
	display       engine.Observer[*PeakPursuitState]
}

type GameConfig struct {
	Symbols                   []Symbol
	MaxLevel                  int
	VisualDelay               time.Duration
	Currency                  string
	ResetSymbolRepresentation string
	WildSymbolRepresentation  string
	DummySymbolRepresentation string
	SpinSymbolRepresentation  string
}

type GameState struct {
	Cash                  int
	SymbolLevels          map[string]SymbolState
	TotalSpins            int
	TotalWins             int
	PeakCash              int
	LastBet               int
	TotalBetAmount        int
	TotalWinAmount        int
	Winnings              int
	TotalBets             int
	WinRate               float64
	MaxDrawdown           int
	Wins                  []int
	StartTime             time.Time
	StartingCash          int
	TopWinAmount          int
	CurrentSpinStreak     int
	LongestSpinStreak     int
	MaxLevelReached       int
	LongestMaxLevelStreak int
	CurrentMaxLevelStreak int
}

type DisplayConfig struct {
	WinningDescription string
}

type SymbolState struct {
	CurrentLevel int
	MaxReached   bool
}

// NewSlotMachine creates a new instance of SlotMachine with given configurations.
func NewSlotMachine(rng randomizer.Randomizer, startingCash int, currency string, display engine.Observer[*PeakPursuitState]) *SlotMachine {

	slotMachine := &SlotMachine{
		Title: "💫 Peak Pursuit 💫",
		RNG:   rng,
		GameConfig: GameConfig{
			Symbols:                   Symbols,
			MaxLevel:                  5,
			Currency:                  currency,
			ResetSymbolRepresentation: "👽",
			WildSymbolRepresentation:  "🌠",
			DummySymbolRepresentation: "  ",
		},
		GameState: GameState{
			Cash:         startingCash,
			SymbolLevels: make(map[string]SymbolState),
			TopWinAmount: 0,
			TotalSpins:   0,
			TotalWins:    0,
			WinRate:      0,
			PeakCash:     startingCash,
			StartTime:    time.Now(),
			StartingCash: startingCash,
		},
		display: display,
	}

	weightedSymbols := generateWeightedSymbols()

	// Initialize symbol levels for each symbol in the config
	for _, symbol := range weightedSymbols {
		slotMachine.GameState.SymbolLevels[symbol] = SymbolState{
			CurrentLevel: 0,
			MaxReached:   false,
		}
	}

	return slotMachine
}

func (s *SlotMachine) Spin() error {
	// Reset the winning description
	s.DisplayConfig.WinningDescription = ""

	// Deduct the bet if applicable
	if s.GameState.Cash < s.GameState.LastBet {
		return fmt.Errorf("not enough cash to make this bet")
	}
	s.GameState.Cash -= s.GameState.LastBet
	s.GameState.TotalBetAmount += s.GameState.LastBet
	s.GameState.TotalBets++

	// Perform the spin
	s.spinWheel()
	s.GameState.PeakCash, s.GameState.MaxDrawdown = utils.UpdateDrawDown(s.GameState.Cash, s.GameState.PeakCash, s.GameState.MaxDrawdown)

	// Check for special symbols and adjust levels
	if s.GameConfig.SpinSymbolRepresentation == s.GameConfig.ResetSymbolRepresentation {
		s.resetAllSymbols()
	} else if s.GameConfig.SpinSymbolRepresentation == s.GameConfig.WildSymbolRepresentation {
		s.handleWildSymbol()
	} else if s.GameConfig.SpinSymbolRepresentation == s.GameConfig.DummySymbolRepresentation {

	} else {
		s.processSymbolLevelIncrease(s.GameConfig.SpinSymbolRepresentation)

	}

	// Calculate and apply winnings
	s.CalculateWinnings()

	// Display the results
	s.DisplayResults()

	return nil
}

func (s *SlotMachine) DisplayResults() {
	s.GameState.WinRate = utils.CalculateWinRate(s.GameState.TotalSpins, s.GameState.TotalWins)
	s.NotifyObservers()
}

func (s *SlotMachine) RequestCommand(input string) (int, error) {
	input = strings.TrimSpace(input)
	var bet int
	var err error

	if input == "" {
		if s.GameState.LastBet == 0 || s.GameState.LastBet > s.GameState.Cash {
			return 0, fmt.Errorf("invalid re-bet amount")
		}
		bet = s.GameState.LastBet
	} else if input == "cashout" {
		s.CashOut()
		s.NotifyObservers()
		return 0, nil
	} else {
		// Attempt to parse the new bet input
		bet, err = strconv.Atoi(input)
		if err != nil || bet < 0 {
			return 0, fmt.Errorf("invalid bet")
		}

		// Check if the new bet is different from the last bet and if there's an active pot
		if bet != s.GameState.LastBet && s.hasActivePot() {
			return 0, fmt.Errorf("cashout first")
		}

		// Check if the bet exceeds available cash
		if bet > s.GameState.Cash {
			return 0, fmt.Errorf("bet exceeds available cash")
		}
	}

	s.GameState.LastBet = bet
	return bet, nil
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
	s.GameState.TotalSpins++
	s.GameState.CurrentSpinStreak++
}

func (s *SlotMachine) CalculateWinnings() (description string, amount int) {
	return "", 0
}

// handleWildSymbol processes the effect of a wild symbol spin.
func (s *SlotMachine) handleWildSymbol() {
	for symbolKey, _ := range s.GameState.SymbolLevels {
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
	state := s.GameState.SymbolLevels[symbolKey]
	if state.CurrentLevel >= s.GameConfig.MaxLevel {
		s.GameState.CurrentMaxLevelStreak++
		s.GameState.MaxLevelReached++
		s.processPayouts(symbolKey, state)
		return
	}

	state.CurrentLevel++
	s.GameState.SymbolLevels[symbolKey] = state

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
	payout := s.GameState.LastBet * payoutMultiplier

	// Check if the symbol has reached the maximum level
	if state.CurrentLevel == s.GameConfig.MaxLevel {
		if state.MaxReached {
			// If max level was already reached, handle instant payout
			s.applyPayout(symbolKey, payout, true) // Instant payout
		} else {
			// First time reaching max level, mark as reached and add to pot
			state.MaxReached = true
			s.GameState.SymbolLevels[symbolKey] = state
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
		s.GameState.Cash += payout
		s.GameState.Wins = utils.RecordWin(payout, s.GameState.Wins)
		s.GameState.TotalWins++
		s.GameState.TotalWinAmount += payout
		if s.DisplayConfig.WinningDescription != "" {
			s.DisplayConfig.WinningDescription += "\n"
		}
		s.DisplayConfig.WinningDescription += fmt.Sprintf("%s instant payout: %d%s %s", symbolKey, payout, s.GameConfig.Currency, symbolKey)
	} else {
		s.GameState.Winnings += payout
	}
	s.updateTopWinAmount(payout)
}

// updateTopWinAmount updates the top win amount if the current payout exceeds it.
func (s *SlotMachine) updateTopWinAmount(payout int) {
	if payout > s.GameState.TopWinAmount {
		s.GameState.TopWinAmount = payout
	}
}

func (s *SlotMachine) CashOut() {
	s.GameState.Cash += s.GameState.Winnings
	s.GameState.TotalWinAmount += s.GameState.Winnings
	s.GameState.Wins = utils.RecordWin(s.GameState.Winnings, s.GameState.Wins)
	s.updateTopWinAmount(s.GameState.Winnings)
	s.GameState.TotalWins++
	s.GameState.Winnings = 0
	s.resetAllSymbols()
	s.DisplayConfig.WinningDescription = "💰 Cash out 💰"
}

// hasActivePot returns true if any symbol has a level above the initial level.
func (s *SlotMachine) hasActivePot() bool {
	for _, state := range s.GameState.SymbolLevels {
		if state.CurrentLevel > 0 {
			return true
		}
	}
	return false
}
