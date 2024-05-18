package engine

import (
	"time"
)

type GameState struct {
	Amount         int
	TotalSpins     int
	TotalWins      int
	WinRate        float64
	PeakAmount     int
	LastBet        int
	TotalBetAmount int
	TotalBets      int
	TotalWinAmount int
	TopWinAmount   int
	MaxWinCounter  int
	StartTime      time.Time
	StartingAmount int
	MaxDrawDown    int
	Wins           []int
}
