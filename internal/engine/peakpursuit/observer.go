package peakpursuit

type PeakPursuitObserverState struct {
	Title string
	PeakPursuitGameState
	WinningDescription       string
	Currency                 string
	SpinSymbolRepresentation string
	Winnings                 int
	SymbolLevels             map[string]SymbolState
	Symbols                  []Symbol
	LongestSpinStreak        int
	MaxLevelReached          int
	InstantWinCounter        map[string]int
}

func (s *SlotMachine) NotifyObservers() {
	state := &PeakPursuitObserverState{
		Title:                    s.Title,
		PeakPursuitGameState:     s.State,
		WinningDescription:       s.DisplayConfig.WinningDescription,
		Currency:                 s.GameConfig.Currency,
		SpinSymbolRepresentation: s.GameConfig.SpinSymbolRepresentation,
		Winnings:                 s.State.Winnings,
		SymbolLevels:             s.State.SymbolLevels,
		Symbols:                  s.GameConfig.Symbols,
		LongestSpinStreak:        s.State.LongestSpinStreak,
		MaxLevelReached:          s.State.MaxLevelReached,
		InstantWinCounter:        s.State.InstantWinCounter,
	}
	s.observerManager.NotifyObservers(state)
}

func (s *SlotMachine) EnableObserver() {
	s.observerManager.RegisterObserver(s.display)
}

func (s *SlotMachine) DisableObserver() {
	s.observerManager.UnregisterObserver(s.display)
}
