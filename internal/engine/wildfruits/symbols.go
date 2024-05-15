package wildfruits

type Symbol struct {
	Representation string
	Payout         int
	Frequency      int
}

var Symbols = []Symbol{
	{"🍉", 10, 5},
	{"🥥", 5, 7},
	{"🌟", 5, 7}, // wild triggers free games
	{"🫐", 3, 8},
	{"🍇", 3, 8},
	{"🍏", 2, 10},
	{"🍓", 2, 10},
	{"🍊", 2, 10},
	{"🍒", 1, 35},
}

func (s *SlotMachine) countWildSymbols() int {
	wildCount := 0
	for _, row := range s.GameConfig.Wheels {
		for _, symbol := range row {
			if symbol == s.GameConfig.WildSymbol {
				wildCount++
			}
		}
	}
	return wildCount
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
