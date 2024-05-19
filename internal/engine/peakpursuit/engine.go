package peakpursuit

import (
	"fmt"
	"go-luck-simulator/internal/engine"
	"go-luck-simulator/internal/randomizer"
	"go-luck-simulator/internal/utils"
	"strconv"
	"strings"
	"time"
)

type SlotMachine struct {
	Title           string
	RNG             randomizer.Randomizer
	GameConfig      GameConfig
	State           PeakPursuitGameState
	DisplayConfig   DisplayConfig
	observerManager *engine.ObserverManager[*PeakPursuitObserverState]
	display         engine.Observer[*PeakPursuitObserverState]
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

type PeakPursuitGameState struct {
	engine.GameState
	SymbolLevels          map[string]SymbolState
	Winnings              int
	InstantWinCounter     map[string]int
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
	CurrentLevel      int
	MaxAlreadyReached bool
}

// NewSlotMachine creates a new instance of SlotMachine with given configurations.
func NewSlotMachine(rng randomizer.Randomizer, startingAmount int, currency string, display engine.Observer[*PeakPursuitObserverState]) *SlotMachine {

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
		State: PeakPursuitGameState{
			GameState: engine.GameState{
				Amount:         startingAmount,
				TotalSpins:     0,
				TotalWins:      0,
				WinRate:        0,
				PeakAmount:     startingAmount,
				LastBet:        0,
				TotalBetAmount: 0,
				TotalBets:      0,
				TotalWinAmount: 0,
				TopWinAmount:   0,
				StartTime:      time.Time{},
				StartingAmount: startingAmount,
				MaxDrawDown:    0,
				Wins:           nil,
			},
			SymbolLevels:      make(map[string]SymbolState),
			InstantWinCounter: make(map[string]int),
		},
		display:         display,
		observerManager: engine.NewObserverManager[*PeakPursuitObserverState](),
	}

	weightedSymbols := generateWeightedSymbols()

	// Initialize symbol levels for each symbol in the config
	for _, symbol := range weightedSymbols {
		slotMachine.State.SymbolLevels[symbol] = SymbolState{
			CurrentLevel:      0,
			MaxAlreadyReached: false,
		}
	}

	return slotMachine
}

func (s *SlotMachine) Spin() error {
	s.State.StartTime = time.Now()

	// Reset the winning description
	s.DisplayConfig.WinningDescription = ""

	// Deduct the bet if applicable
	if s.State.Amount < s.State.LastBet {
		return fmt.Errorf("not enough cash to make this bet")
	}
	s.State.Amount -= s.State.LastBet
	s.State.TotalBetAmount += s.State.LastBet
	s.State.TotalBets++

	// Perform the spin
	s.spinWheel()
	s.State.PeakAmount, s.State.MaxDrawDown = utils.UpdateDrawDown(s.State.Amount, s.State.PeakAmount, s.State.MaxDrawDown)

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
	s.State.WinRate = utils.CalculateWinRate(s.State.TotalSpins, s.State.TotalWins)
	s.NotifyObservers()
}

func (s *SlotMachine) RequestCommand(input string) (int, error) {
	input = strings.TrimSpace(input)
	var bet int
	var err error

	if input == "" {
		if s.State.LastBet == 0 || s.State.LastBet > s.State.Amount {
			return 0, fmt.Errorf("invalid re-bet amount")
		}
		bet = s.State.LastBet
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
		if bet != s.State.LastBet && s.hasActivePot() {
			return 0, fmt.Errorf("cashout first")
		}

		// Check if the bet exceeds available cash
		if bet > s.State.Amount {
			return 0, fmt.Errorf("bet exceeds available cash")
		}
	}

	s.State.LastBet = bet
	return bet, nil
}

func (s *SlotMachine) CalculateWinnings() (description string, amount int) {
	return "", 0
}
