package wildfruits

import (
	"fmt"
	"time"
)

func (s *SlotMachine) chooseBonusSymbol() {
	var filteredSymbols []string
	for _, sym := range Symbols {
		if sym.Representation != s.GameConfig.WildSymbol { // Exclude the wild symbol
			for i := 0; i < sym.Frequency; i++ {
				filteredSymbols = append(filteredSymbols, sym.Representation)
			}
		}
	}

	if len(filteredSymbols) > 0 {
		var index int
		// Simulate a more dynamic "spinning wheel" effect
		for spin := 0; spin < 30; spin++ {
			index = s.RNG.Intn(len(filteredSymbols))
			s.GameConfig.BonusSymbol = filteredSymbols[index]
			s.NotifyObservers() // Notify observers with each change to update the display dynamically
			time.Sleep(s.GameConfig.VisualDelay)
		}
		// Final selection, ensure it's updated one last time if needed
		s.GameConfig.BonusSymbol = filteredSymbols[index]
		s.NotifyObservers()
	} else {
		fmt.Println("No eligible Symbols found for bonus.")
	}
}

func (s *SlotMachine) resetLockedPositions() {
	s.GameState.LockedPositions = make(map[int]map[int]bool)
}

func (s *SlotMachine) lockBonusSymbols() {
	if s.GameState.BonusGames > 0 {
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				if s.GameConfig.Wheels[row][col] == s.GameConfig.BonusSymbol {
					if s.GameState.LockedPositions[row] == nil {
						s.GameState.LockedPositions[row] = make(map[int]bool)
					}
					s.GameState.LockedPositions[row][col] = true
				}
			}
		}
	}
}

func (s *SlotMachine) handleBonusGameInitiation() {
	if s.GameState.BonusGames == 0 { // Only initiate new bonus games if none are active
		s.chooseBonusSymbol()       // Choose a new bonus symbol
		s.GameState.BonusGames += 7 // Start 7 free games
		s.DisplayConfig.BonusGamesInitiated++
	} else {
		s.GameState.BonusGames += 2 // Extend existing free games
	}
}
