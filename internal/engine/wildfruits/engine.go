package wildfruits

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
	State           WildFruitsGameState
	DisplayConfig   DisplayConfig
	observerManager *engine.ObserverManager[*WildFruitObserverState]
	display         engine.Observer[*WildFruitObserverState]
}

type GameConfig struct {
	Wheels          [3][3]string
	WeightedSymbols []string
	WildSymbol      string
	BonusSymbol     string
	VisualDelay     time.Duration
	Currency        string
	MaxWin          int
}

type WildFruitsGameState struct {
	engine.GameState
	BonusGames            int
	LockedPositions       map[int]map[int]bool
	TotalWinAmountPerSpin int
}

type DisplayConfig struct {
	WinningDescription  string
	BonusGamesInitiated int
	WinningPositions    map[int]map[int]bool
}

func NewSlotMachine(rng randomizer.Randomizer, startingAmount int, currency string, display engine.Observer[*WildFruitObserverState]) *SlotMachine {
	// Generate the weighted symbols list once
	weightedSymbols := generateWeightedSymbols()

	// Initialize the nested maps for WinningPositions
	winningPositions := make(map[int]map[int]bool)
	for i := 0; i < 3; i++ {
		winningPositions[i] = make(map[int]bool)
	}

	// Create a new SlotMachine with all necessary fields initialized
	sm := &SlotMachine{
		Title: "🍒 Wild Fruits 🍒",
		RNG:   rng,
		GameConfig: GameConfig{
			Wheels:          generateInitialWheels(weightedSymbols),
			WeightedSymbols: weightedSymbols,
			WildSymbol:      "🌟",
			BonusSymbol:     "  ",
			Currency:        currency,
			MaxWin:          50, // max multiplier * winning lines
		},
		State: WildFruitsGameState{
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
			BonusGames:            0,
			LockedPositions:       make(map[int]map[int]bool),
			TotalWinAmountPerSpin: 0,
		},
		DisplayConfig: DisplayConfig{
			WinningPositions: winningPositions,
		},
		display:         display,
		observerManager: engine.NewObserverManager[*WildFruitObserverState](),
	}

	return sm
}
func (s *SlotMachine) Spin() error {
	s.State.StartTime = time.Now()
	// Reset winning positions
	for i := range s.GameConfig.Wheels {
		for j := range s.GameConfig.Wheels[i] {
			s.DisplayConfig.WinningPositions[i][j] = false
		}
	}
	// Deduct the bet only if there are no bonus games active
	if s.State.BonusGames == 0 {
		if s.State.LastBet > s.State.Amount {
			return fmt.Errorf("not enough cash to make this bet")
		}
		s.State.Amount -= s.State.LastBet
		s.State.TotalBetAmount += s.State.LastBet
		s.GameConfig.BonusSymbol = "  "
	}

	// Spin the wheels regardless of bonus or normal play
	s.spinWheels()
	s.State.PeakAmount, s.State.MaxDrawDown = utils.UpdateDrawDown(s.State.Amount, s.State.PeakAmount, s.State.MaxDrawDown)

	// Lock or reset positions based on bonus state
	if s.State.BonusGames > 0 {
		s.lockBonusSymbols()
		s.State.BonusGames--
		if s.State.BonusGames == 0 {
			s.resetLockedPositions()
		}
	} else {
		s.resetLockedPositions()
	}

	// Check for 3 wild symbols to potentially initiate or extend bonus games
	wildCount := s.countWildSymbols()
	if wildCount >= 3 {
		s.handleBonusGameInitiation()
	}

	return nil
}

func (s *SlotMachine) DisplayResults() {
	// Display winnings or losses
	winningDescriptions, winAmount := s.CalculateWinnings()
	if winningDescriptions != "" {
		s.State.Amount += winAmount
		s.State.Wins = utils.RecordWin(winAmount, s.State.Wins)
		s.State.TotalWins++
	}
	s.State.TotalWinAmount += winAmount
	// Notify observers after calculating and displaying the win for this spin
	s.DisplayConfig.WinningDescription = winningDescriptions
	s.State.TotalWinAmountPerSpin = winAmount
	s.State.WinRate = utils.CalculateWinRate(s.State.TotalSpins, s.State.TotalWins)
	s.NotifyObservers()
}

func (s *SlotMachine) CalculateWinnings() (string, int) {
	var totalWinAmount int
	var winDescriptions string

	// Create a map from symbol representation to payout
	symbolPayouts := make(map[string]int)
	for _, sym := range Symbols {
		symbolPayouts[sym.Representation] = sym.Payout
	}

	// Check horizontal and diagonal lines
	lines := append(s.GameConfig.Wheels[:], [3]string{s.GameConfig.Wheels[0][0], s.GameConfig.Wheels[1][1], s.GameConfig.Wheels[2][2]}, [3]string{s.GameConfig.Wheels[0][2], s.GameConfig.Wheels[1][1], s.GameConfig.Wheels[2][0]})

	for i, line := range lines {
		if isWinningLine(line) {
			mainSymbol := getMainSymbol(line)
			if mainSymbol == "" {
				mainSymbol = s.GameConfig.WildSymbol
			}
			winAmount := s.State.LastBet * symbolPayouts[mainSymbol]
			totalWinAmount += winAmount
			s.markWinningPositions(i)
			if winDescriptions != "" {
				winDescriptions += "\n"
			}
			winDescriptions += fmt.Sprintf("%s on line %d: %d%s", mainSymbol, i+1, winAmount, s.GameConfig.Currency)
		}
	}
	// Update TopWinAmount if the total win amount is greater than the current TopWinAmount
	if totalWinAmount > s.State.TopWinAmount {
		s.State.TopWinAmount = totalWinAmount
	}

	return winDescriptions, totalWinAmount
}

func (s *SlotMachine) RequestCommand(input string) (int, error) {
	input = strings.TrimSpace(input)
	var bet int
	var err error

	if s.State.BonusGames > 0 {
		return 0, nil
	}

	if input == "" {
		if s.State.LastBet == 0 || s.State.LastBet > s.State.Amount {
			return 0, fmt.Errorf("invalid re-bet amount")
		}
		bet = s.State.LastBet
	} else {
		bet, err = strconv.Atoi(input)
		if err != nil || bet < 0 || bet > s.State.Amount {
			return 0, fmt.Errorf("invalid bet")
		}
	}

	s.State.LastBet = bet
	s.State.TotalBets++

	return bet, nil
}

func (s *SlotMachine) spinWheels() {
	s.State.TotalSpins++
	for col := 0; col < 3; col++ {
		s.spinColumn(col)
	}
}

func (s *SlotMachine) spinColumn(col int) {
	for spin := 0; spin < 20; spin++ {
		s.updateColumnSymbols(col)
		s.NotifyObservers()
		time.Sleep(s.GameConfig.VisualDelay)
	}
}

func (s *SlotMachine) updateColumnSymbols(col int) {
	for row := 0; row < 3; row++ {
		if !s.isLocked(row, col) {
			s.GameConfig.Wheels[row][col] = s.getRandomSymbol()
		}
	}
}

func (s *SlotMachine) isLocked(row, col int) bool {
	locked, ok := s.State.LockedPositions[row][col]
	return ok && locked
}

func (s *SlotMachine) getRandomSymbol() string {
	return s.GameConfig.WeightedSymbols[s.RNG.Intn(len(s.GameConfig.WeightedSymbols))]
}
