package wildfruits

import "go-slot-machine/internal/engine"

func (s *SlotMachine) RegisterObserver(o engine.Observer) {
	s.observers = append(s.observers, o)
}

func (s *SlotMachine) NotifyObservers() {
	state := engine.DisplayState{
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
	}
	for _, observer := range s.observers {
		observer.Update(state)
	}
}
