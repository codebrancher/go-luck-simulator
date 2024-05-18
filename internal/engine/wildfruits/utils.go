package wildfruits

import (
	"math/rand"
	"time"
)

func isWinningLine(line [3]string) bool {
	uniqueSymbols := make(map[string]bool)
	for _, symbol := range line {
		if symbol != "🌟" {
			uniqueSymbols[symbol] = true
		}
	}
	return len(uniqueSymbols) <= 1
}

func getMainSymbol(line [3]string) string {
	symbolCount := make(map[string]int)
	for _, symbol := range line {
		if symbol != "🌟" {
			symbolCount[symbol]++
		}
	}

	maxCount := 0
	mainSymbol := ""
	for symbol, count := range symbolCount {
		if count > maxCount {
			maxCount = count
			mainSymbol = symbol
		}
	}
	return mainSymbol
}

func generateInitialWheels(weightedSymbols []string) [3][3]string {
	var wheels [3][3]string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			wheels[i][j] = weightedSymbols[rand.Intn(len(weightedSymbols))]
		}
	}
	return wheels
}

func (s *SlotMachine) getLinePositions(lineIndex int) [][2]int {
	var positions [][2]int
	switch lineIndex {
	case 0: // Top horizontal line
		positions = append(positions, [2]int{0, 0}, [2]int{0, 1}, [2]int{0, 2})
	case 1: // Middle horizontal line
		positions = append(positions, [2]int{1, 0}, [2]int{1, 1}, [2]int{1, 2})
	case 2: // Bottom horizontal line
		positions = append(positions, [2]int{2, 0}, [2]int{2, 1}, [2]int{2, 2})
	case 3: // Left diagonal
		positions = append(positions, [2]int{0, 0}, [2]int{1, 1}, [2]int{2, 2})
	case 4: // Right diagonal
		positions = append(positions, [2]int{0, 2}, [2]int{1, 1}, [2]int{2, 0})
	}
	return positions
}

func (s *SlotMachine) markWinningPositions(lineIndex int) {
	positions := s.getLinePositions(lineIndex)
	for _, pos := range positions {
		s.DisplayConfig.WinningPositions[pos[0]][pos[1]] = true
	}
}

func (s *SlotMachine) SetVisualDelay(delay time.Duration) {
	s.GameConfig.VisualDelay = delay
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
