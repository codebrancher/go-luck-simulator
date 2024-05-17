package peakpursuit

import (
	"go-slot-machine/internal/engine"
)

type PeakPursuitState struct {
	Title                    string
	Cash                     int
	LastBet                  int
	WinningDescription       string
	Currency                 string
	SpinSymbolRepresentation string
	Winnings                 int
	SymbolLevels             map[string]SymbolState
	Symbols                  []Symbol
	TotalBetAmount           int
	TotalWinAmount           int
	LongestSpinStreak        int
	MaxLevelReached          int
}

func (s *SlotMachine) RegisterObserver(o engine.Observer[*PeakPursuitState]) {
	s.observers = append(s.observers, o)
}

func (s *SlotMachine) UnregisterObserver(o engine.Observer[*PeakPursuitState]) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *SlotMachine) NotifyObservers() {
	state := &PeakPursuitState{
		Title:                    s.Title,
		Cash:                     s.GameState.Cash,
		LastBet:                  s.GameState.LastBet,
		WinningDescription:       s.DisplayConfig.WinningDescription,
		Currency:                 s.GameConfig.Currency,
		SpinSymbolRepresentation: s.GameConfig.SpinSymbolRepresentation,
		Winnings:                 s.GameState.Winnings,
		SymbolLevels:             s.GameState.SymbolLevels,
		Symbols:                  s.GameConfig.Symbols,
		TotalBetAmount:           s.GameState.TotalBetAmount,
		TotalWinAmount:           s.GameState.TotalWinAmount,
		LongestSpinStreak:        s.GameState.LongestSpinStreak,
		MaxLevelReached:          s.GameState.MaxLevelReached,
	}
	for _, observer := range s.observers {
		observer.Update(state)
	}
}

func (s *SlotMachine) EnableObserver() {
	s.RegisterObserver(s.display)
}

func (s *SlotMachine) DisableObserver() {
	s.UnregisterObserver(s.display)
}
