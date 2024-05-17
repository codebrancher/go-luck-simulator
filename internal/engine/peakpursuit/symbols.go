package peakpursuit

import "go-slot-machine/internal/utils"

type Symbol struct {
	Representation string
	Payouts        []int
	Frequency      int
}

var Symbols = []Symbol{
	{"💫", []int{1, 2, 4, 7, 10}, 20},
	{"🪐", []int{3, 8, 16, 29, 50}, 12},
	{"🚀", []int{10, 40, 100, 200, 350}, 4},
	{"👽", []int{0}, 12},
	{"🌠", []int{0}, 2},
	{"  ", []int{0}, 50},
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
	for key := range s.GameState.SymbolLevels {
		s.GameState.SymbolLevels[key] = SymbolState{0, false}
	}
	s.GameState.LongestSpinStreak = utils.UpdateLongestSpinStreak(s.GameState.CurrentSpinStreak, s.GameState.LongestSpinStreak)
	s.GameState.LongestMaxLevelStreak = utils.UpdateMaxLevelStreak(s.GameState.CurrentMaxLevelStreak, s.GameState.LongestMaxLevelStreak)
	s.GameState.CurrentMaxLevelStreak = 0
	s.GameState.CurrentSpinStreak = 0
	s.GameState.Winnings = 0
	s.DisplayConfig.WinningDescription = "👽 Game state reset 👽"
}
