package utils

import (
	"math"
)

func CalculateRTP(totalWinAmount, totalBetAmount int) float64 {
	return float64(totalWinAmount) / float64(totalBetAmount) * 100
}

func CalculateWinRate(totalSpins, totalWins int) float64 {
	if totalSpins == 0 {
		return 0
	}
	return (float64(totalWins) / float64(totalSpins)) * 100
}

func CalculateAverageBet(totalBetAmount, totalBets int) float64 {
	return float64(totalBetAmount) / float64(totalBets)
}

func CalculateAveragePayout(totalWinAmount, totalWins int) float64 {
	return float64(totalWinAmount) / float64(totalWins)
}

func CalculateProfitability(cash, startingCash int) int {
	return cash - startingCash
}

func UpdateDrawDown(cash, peakCash, maxDrawDown int) (int, int) {
	if cash > peakCash {
		peakCash = cash
	} else {
		currentDrawDown := peakCash - cash
		if currentDrawDown > maxDrawDown {
			maxDrawDown = currentDrawDown
		}
	}
	return peakCash, maxDrawDown
}

func UpdateLongestSpinStreak(currentSpinStreak, longestSpinStreak int) int {
	if currentSpinStreak > longestSpinStreak {
		return currentSpinStreak
	}
	return longestSpinStreak
}

func UpdateMaxLevelStreak(currentMaxLevelStreak, longestMaxLevelStreak int) int {
	if currentMaxLevelStreak > longestMaxLevelStreak {
		return currentMaxLevelStreak
	}
	return longestMaxLevelStreak
}

func CalculateWinDeviation(wins []int) float64 {
	if len(wins) == 0 {
		return 0.0
	}
	mean, variance := 0.0, 0.0
	for _, win := range wins {
		mean += float64(win)
	}
	mean /= float64(len(wins))

	for _, win := range wins {
		variance += math.Pow(float64(win)-mean, 2)
	}
	variance /= float64(len(wins))
	return math.Sqrt(variance)
}

func RecordWin(winAmount int, wins []int) []int {
	wins = append(wins, winAmount)
	return wins
}
