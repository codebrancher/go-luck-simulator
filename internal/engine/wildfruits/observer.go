package wildfruits

import (
	"go-slot-machine/internal/engine"
	"time"
)

type WildFruitState struct {
	Title              string
	Cash               int
	Wheels             [3][3]string
	BonusSymbol        string
	LastBet            int
	TotalWinAmount     int
	WinningDescription string
	FreeGames          int
	TopWinAmount       int
	WinRate            float64
	WinningPositions   map[int]map[int]bool
	StartTime          time.Time
	StartingCash       int
	Currency           string
}

func (s *SlotMachine) RegisterObserver(o engine.Observer[*WildFruitState]) {
	s.observers = append(s.observers, o)
}

func (s *SlotMachine) UnregisterObserver(o engine.Observer[*WildFruitState]) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *SlotMachine) NotifyObservers() {
	state := &WildFruitState{
		Title:              s.Title,
		Cash:               s.GameState.Cash,
		Wheels:             s.GameConfig.Wheels,
		BonusSymbol:        s.GameConfig.BonusSymbol,
		LastBet:            s.GameState.LastBet,
		TotalWinAmount:     s.GameState.TotalWinAmountPerSpin,
		WinningDescription: s.DisplayConfig.WinningDescription,
		FreeGames:          s.GameState.BonusGames,
		TopWinAmount:       s.GameState.TopWinAmount,
		WinRate:            s.GameState.WinRate,
		WinningPositions:   s.DisplayConfig.WinningPositions,
		StartTime:          s.GameState.StartTime,
		StartingCash:       s.GameState.StartingCash,
		Currency:           s.GameConfig.Currency,
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
