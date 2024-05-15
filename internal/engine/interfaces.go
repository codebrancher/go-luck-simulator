package engine

import "time"

type GameEngine interface {
	Spin() error                                         // Conducts the spin logic without displaying results
	DisplayResults()                                     // Handles all output logic separate from game mechanics
	CalculateWinnings() (description string, amount int) // Calculates winnings based on the last spin
	RequestBet(input string) (int, error)                // Handles the bet input and validation logic
}

type Observer interface {
	Update(state DisplayState)
}

type DisplayState struct {
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
