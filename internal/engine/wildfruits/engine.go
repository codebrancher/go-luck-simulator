package wildfruits

import (
	"fmt"
	"go-slot-machine/internal/engine"
	"go-slot-machine/internal/randomizer"
	"strconv"
	"strings"
	"time"
)

type SlotMachine struct {
	Title         string
	RNG           randomizer.Randomizer
	GameConfig    GameConfig
	GameState     GameState
	DisplayConfig DisplayConfig
	observers     []engine.Observer
}

type GameConfig struct {
	Wheels          [3][3]string
	WeightedSymbols []string
	WildSymbol      string
	BonusSymbol     string
	VisualDelay     time.Duration
	Currency        string
}

type GameState struct {
	Cash                  int
	LastBet               int
	BonusGames            int
	LockedPositions       map[int]map[int]bool
	TotalWinAmountPerSpin int
	TotalWinAmount        int
	TopWinAmount          int
	TotalBetAmount        int
	TotalBets             int
	TotalSpins            int
	TotalWins             int
	WinRate               float64
	PeakCash              int
	MaxDrawdown           int
	Wins                  []int
	StartTime             time.Time
	StartingCash          int
}

type DisplayConfig struct {
	WinningDescription  string
	BonusGamesInitiated int
	WinningPositions    map[int]map[int]bool
}

func NewSlotMachine(rng randomizer.Randomizer, visualDelay time.Duration, startingCash int, currency string) *SlotMachine {
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
			VisualDelay:     visualDelay,
			Currency:        currency,
		},
		GameState: GameState{
			Cash:                  startingCash,
			LastBet:               0,
			BonusGames:            0,
			LockedPositions:       make(map[int]map[int]bool),
			TotalWinAmountPerSpin: 0,
			TopWinAmount:          0,
			TotalSpins:            0,
			TotalWins:             0,
			WinRate:               0,
			PeakCash:              startingCash,
			StartTime:             time.Now(),
			StartingCash:          startingCash,
		},
		DisplayConfig: DisplayConfig{
			WinningPositions: winningPositions,
		},
	}

	return sm
}
func (s *SlotMachine) Spin() error {
	// Reset winning positions
	for i := range s.GameConfig.Wheels {
		for j := range s.GameConfig.Wheels[i] {
			s.DisplayConfig.WinningPositions[i][j] = false
		}
	}
	// Deduct the bet only if there are no bonus games active
	if s.GameState.BonusGames == 0 {
		if s.GameState.LastBet > s.GameState.Cash {
			return fmt.Errorf("not enough cash to make this bet")
		}
		s.GameState.Cash -= s.GameState.LastBet
		s.GameState.TotalBetAmount += s.GameState.LastBet
		s.GameConfig.BonusSymbol = "  "
	}

	// Spin the wheels regardless of bonus or normal play
	s.spinWheels()
	s.UpdateDrawdown()

	// Lock or reset positions based on bonus state
	if s.GameState.BonusGames > 0 {
		s.lockBonusSymbols()
		s.GameState.BonusGames--
		if s.GameState.BonusGames == 0 {
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
		s.GameState.Cash += winAmount
		s.RecordWin(winAmount)
		s.GameState.TotalWins++
	}
	s.GameState.TotalWinAmount += winAmount
	// Notify observers after calculating and displaying the win for this spin
	s.DisplayConfig.WinningDescription = winningDescriptions
	s.GameState.TotalWinAmountPerSpin = winAmount
	s.GameState.WinRate = s.calculateWinRate()
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
			winAmount := s.GameState.LastBet * symbolPayouts[mainSymbol]
			totalWinAmount += winAmount
			s.markWinningPositions(i)
			if winDescriptions != "" {
				winDescriptions += "\n"
			}
			winDescriptions += fmt.Sprintf("%s on line %d: %d%s", mainSymbol, i+1, winAmount, s.GameConfig.Currency)
		}
	}
	// Update TopWinAmount if the total win amount is greater than the current TopWinAmount
	if totalWinAmount > s.GameState.TopWinAmount {
		s.GameState.TopWinAmount = totalWinAmount
	}

	return winDescriptions, totalWinAmount
}

func (s *SlotMachine) RequestBet(input string) (int, error) {
	input = strings.TrimSpace(input)
	var bet int
	var err error

	if s.GameState.BonusGames > 0 {
		return 0, nil
	}

	if input == "" {
		if s.GameState.LastBet == 0 || s.GameState.LastBet > s.GameState.Cash {
			return 0, fmt.Errorf("invalid re-bet amount")
		}
		bet = s.GameState.LastBet
	} else {
		bet, err = strconv.Atoi(input)
		if err != nil || bet < 0 || bet > s.GameState.Cash {
			return 0, fmt.Errorf("invalid bet")
		}
	}

	s.GameState.LastBet = bet
	s.GameState.TotalBets++

	return bet, nil
}

func (s *SlotMachine) spinWheels() {
	s.GameState.TotalSpins++
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
	locked, ok := s.GameState.LockedPositions[row][col]
	return ok && locked
}

func (s *SlotMachine) getRandomSymbol() string {
	return s.GameConfig.WeightedSymbols[s.RNG.Intn(len(s.GameConfig.WeightedSymbols))]
}
