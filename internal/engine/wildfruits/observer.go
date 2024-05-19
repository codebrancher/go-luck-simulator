package wildfruits

type WildFruitObserverState struct {
	Title string
	WildFruitsGameState
	Wheels             [3][3]string
	BonusSymbol        string
	WinningDescription string
	BonusGames         int
	WinningPositions   map[int]map[int]bool
	Currency           string
}

func (s *SlotMachine) NotifyObservers() {
	state := &WildFruitObserverState{
		Title:               s.Title,
		Wheels:              s.GameConfig.Wheels,
		BonusSymbol:         s.GameConfig.BonusSymbol,
		WinningDescription:  s.DisplayConfig.WinningDescription,
		WinningPositions:    s.DisplayConfig.WinningPositions,
		Currency:            s.GameConfig.Currency,
		WildFruitsGameState: s.State,
		BonusGames:          s.State.BonusGames,
	}
	s.observerManager.NotifyObservers(state)
}

func (s *SlotMachine) EnableObserver() {
	s.observerManager.RegisterObserver(s.display)
}

func (s *SlotMachine) DisableObserver() {
	s.observerManager.UnregisterObserver(s.display)
}
