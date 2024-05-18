package display

import (
	"fmt"
	"go-slot-machine/internal/engine/peakpursuit"
	"go-slot-machine/internal/utils"
	"strconv"
	"strings"
	"time"
)

type PeakPursuitDisplay struct{}

func (cd *PeakPursuitDisplay) Show(message string) {
	fmt.Println(message)
}

func (cd *PeakPursuitDisplay) ShowStartupInfo(title string) {
	cd.ClearScreen()
	frameWidth := 32
	contentWidth := frameWidth - 2

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintCentered(title, contentWidth-2)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintIntermediaryLine()
	utils.PrintCentered("Commands", contentWidth)
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	utils.PrintLeftAligned("play", contentWidth)
	utils.PrintBlankLine()
	utils.PrintLeftAligned("sim [spins] [bet amount]", contentWidth)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintLeftAligned("exit", contentWidth)
	fmt.Println(border)
	fmt.Println()
}

func (cd *PeakPursuitDisplay) ShowHelp() {
	cd.ClearScreen()
	frameWidth := 32
	contentWidth := frameWidth - 2

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	utils.PrintBlankLine()
	utils.PrintCentered("Need help?", contentWidth)
	utils.PrintBlankLine()
	utils.PrintIntermediaryLine()
	utils.PrintCentered("Commands", contentWidth)
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	utils.PrintLeftAligned("cashout", contentWidth)
	utils.PrintBlankLine()
	utils.PrintLeftAligned("bet [amount]", contentWidth)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintLeftAligned("info", contentWidth)
	utils.PrintBlankLine()
	utils.PrintLeftAligned("stats", contentWidth)
	utils.PrintBlankLine()
	utils.PrintLeftAligned("help", contentWidth)
	utils.PrintBlankLine()
	utils.PrintLeftAligned("exit", contentWidth)
	fmt.Println(border)
	fmt.Println()
}

func (cd *PeakPursuitDisplay) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (cd *PeakPursuitDisplay) Update(state *peakpursuit.PeakPursuitObserverState) {
	cd.ClearScreen()

	frameWidth := 32
	contentWidth := frameWidth - 2
	contentWidthUnicode := frameWidth - 4

	rtp := float64(state.TotalWinAmount) / float64(state.TotalBetAmount) * 100

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintCentered(state.Title, contentWidthUnicode)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintIntermediaryLine()
	utils.PrintCentered("Pot", contentWidth)
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	if state.Winnings > 0 {
		utils.PrintCentered("\033[48;5;22m      "+strconv.Itoa(state.Winnings)+"      \033[0m", contentWidth+14)
	} else {
		utils.PrintCentered(strconv.Itoa(state.Winnings), contentWidth)
	}
	utils.PrintBlankLine()

	maxLevels := 0
	for _, symbol := range state.Symbols {
		if len(symbol.Payouts) > 1 && len(symbol.Payouts) > maxLevels {
			maxLevels = len(symbol.Payouts)
		}
	}

	levelDisplays := make([]string, maxLevels)
	symbolDisplays := make([]string, len(state.Symbols)-3)

	i := 0
	for _, symbol := range state.Symbols {
		if len(symbol.Payouts) > 1 { // Only display symbols with levels
			stateSymbol := state.SymbolLevels[symbol.Representation]
			symbolDisplays[i] = symbol.Representation // Prepare symbol display line

			for level := 1; level <= maxLevels; level++ {
				multiplier := fmt.Sprintf(" %3dx ", symbol.Payouts[level-1])
				bgStart := "\033[32m"
				bgEnd := "\033[0m"
				normalGreen := "\033[33m"
				normalYellow := "\033[39m"

				// Apply color based on the level and max reached condition
				if stateSymbol.MaxReached {
					multiplier = bgStart + multiplier + bgEnd
				} else if level <= stateSymbol.CurrentLevel {
					multiplier = normalGreen + multiplier + bgEnd
				} else {
					multiplier = normalYellow + multiplier + bgEnd
				}

				levelDisplays[level-1] += multiplier + " "
			}
			i++
		}
	}

	// Print each level from highest to lowest (to appear as a rising barometer)
	for j := maxLevels - 1; j >= 0; j-- {
		fmt.Printf("|    %s     |\n", levelDisplays[j])

	}
	utils.PrintBlankLine()
	fmt.Printf("|       %s       |\n", strings.Join(symbolDisplays, "     "))
	utils.PrintBlankLine()
	utils.PrintIntermediaryLine()
	utils.PrintCentered("------", contentWidth)
	fmt.Printf("|            | %s |            |\n", state.SpinSymbolRepresentation)
	utils.PrintCentered("------", contentWidth)
	// Displaying winning descriptions
	if state.WinningDescription != "" {
		utils.PrintIntermediaryLine()
		for _, line := range strings.Split(state.WinningDescription, "\n") {
			utils.PrintCentered(line, contentWidth-2)
		}
	}
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	leftValLine1 := fmt.Sprintf(" Bet: %d%s", state.LastBet, state.Currency)
	rightValLine1 := fmt.Sprintf("RTP: %.1f%% ", rtp)
	utils.PrintWithMiddlePadding(leftValLine1, rightValLine1, contentWidth+2)
	leftValLine2 := fmt.Sprintf(" Cash: %d%s", state.Amount, state.Currency)
	utils.PrintWithMiddlePadding(leftValLine2, "", contentWidth+2)
	utils.PrintBlankLine()
	fmt.Println(border)
	fmt.Println()
}

func (cd *PeakPursuitDisplay) ShowInfo(symbols []peakpursuit.Symbol, game *peakpursuit.SlotMachine) {
	cd.ClearScreen()

	frameWidth := 32
	contentWidth := frameWidth - 2
	contentWidthUnicode := frameWidth - 4

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintCentered(game.Title, contentWidthUnicode)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintIntermediaryLine()

	// Explanation of Wild Symbol and Bonus Games
	resetSymbolDescription := fmt.Sprintf("%s resets game state", game.GameConfig.ResetSymbolRepresentation)
	wildSymbolDescription := fmt.Sprintf("%s increases every symbol", game.GameConfig.WildSymbolRepresentation)
	utils.PrintCentered("Features", contentWidth)
	utils.PrintIntermediaryLine()
	utils.PrintCentered(resetSymbolDescription, contentWidth-1)
	utils.PrintCentered(wildSymbolDescription, contentWidth-1)
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	levelDisplays, symbolDisplays := cd.prepareSymbolDisplays(symbols, game)

	// Print each level from highest to lowest (to appear as a rising barometer)
	for j := len(levelDisplays) - 1; j >= 0; j-- {
		fmt.Printf("|    %s     |\n", levelDisplays[j])

	}
	utils.PrintBlankLine()
	fmt.Printf("|       %s       |\n", strings.Join(symbolDisplays, "     "))
	utils.PrintBlankLine()

	utils.PrintIntermediaryLine()
	utils.PrintCentered("Commands", contentWidth)
	utils.PrintIntermediaryLine()
	utils.PrintLeftAligned("cashout", contentWidth)
	utils.PrintLeftAligned("bet [amount]", contentWidth)
	utils.PrintLeftAligned("help", contentWidth)
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	utils.PrintCentered(fmt.Sprintf("Your cash: %d%s", game.State.Amount, game.GameConfig.Currency), contentWidth)
	utils.PrintBlankLine()
	fmt.Println(border)
	fmt.Println()
}

func (cd *PeakPursuitDisplay) ShowStats(game *peakpursuit.SlotMachine) {
	cd.ClearScreen()
	frameWidth := 32
	contentWidth := frameWidth - 2

	rtp := utils.CalculateRTP(game.State.TotalWinAmount, game.State.TotalBetAmount)
	profitability := utils.CalculateProfitability(game.State.Amount, game.State.StartingAmount)
	avgBetAmount := utils.CalculateAverageBet(game.State.TotalBetAmount, game.State.TotalBets)
	avgPayout := utils.CalculateAveragePayout(game.State.TotalWinAmount, game.State.TotalWins)
	deviation := utils.CalculateWinDeviation(game.State.Wins)

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	utils.PrintBlankLine()
	utils.PrintCentered("Who doesn't like stats?", contentWidth)
	utils.PrintBlankLine()
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	utils.PrintWithMiddlePadding(" RTP", fmt.Sprintf("%.2f%% ", rtp), frameWidth)
	utils.PrintWithMiddlePadding(" Win rate", fmt.Sprintf("%.2f%% ", game.State.WinRate), frameWidth)
	utils.PrintWithMiddlePadding(" Playtime", fmt.Sprintf("%s ", utils.FormatDuration(time.Since(game.State.StartTime))), frameWidth)
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	utils.PrintWithMiddlePadding(" Total spins", fmt.Sprintf("%d ", game.State.TotalSpins), frameWidth)
	utils.PrintWithMiddlePadding(" Total wins", fmt.Sprintf("%d ", game.State.TotalWins), frameWidth)
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	utils.PrintWithMiddlePadding(" Starting cash", fmt.Sprintf("%d%s ", game.State.StartingAmount, game.GameConfig.Currency), frameWidth)
	utils.PrintWithMiddlePadding(" Current cash", fmt.Sprintf("%d%s ", game.State.Amount, game.GameConfig.Currency), frameWidth)
	utils.PrintWithMiddlePadding(" Total win amount", fmt.Sprintf("%d%s ", game.State.TotalWinAmount, game.GameConfig.Currency), frameWidth)
	utils.PrintWithMiddlePadding(" Total bet amount", fmt.Sprintf("%d%s ", game.State.TotalBetAmount, game.GameConfig.Currency), frameWidth)
	utils.PrintWithMiddlePadding(" Max drawdown", fmt.Sprintf("%d%s ", game.State.MaxDrawDown, game.GameConfig.Currency), frameWidth)
	utils.PrintWithMiddlePadding(" Net profit/loss", fmt.Sprintf("%d ", profitability), frameWidth)
	utils.PrintWithMiddlePadding(" Average bet", fmt.Sprintf("%.2f%s ", avgBetAmount, game.GameConfig.Currency), frameWidth)
	utils.PrintWithMiddlePadding(" Average win", fmt.Sprintf("%.2f%s ", avgPayout, game.GameConfig.Currency), frameWidth)
	utils.PrintWithMiddlePadding(" Volatility", fmt.Sprintf("%.2f ", deviation), frameWidth)
	utils.PrintWithMiddlePadding(" Top win", fmt.Sprintf("%d%s ", game.State.TopWinAmount, game.GameConfig.Currency), frameWidth)
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	utils.PrintWithMiddlePadding(" Best Spin Streak", fmt.Sprintf("%d ", game.State.LongestSpinStreak), frameWidth)
	utils.PrintWithMiddlePadding(" Max Level Reached", fmt.Sprintf("%d ", game.State.MaxLevelReached), frameWidth)
	utils.PrintWithMiddlePadding(" Best Max Level Streak", fmt.Sprintf("%d ", game.State.LongestMaxLevelStreak), frameWidth)
	for s, c := range game.State.InstantWinCounter {
		utils.PrintWithMiddlePadding(fmt.Sprintf(" Instant payouts %s", s), fmt.Sprintf("%d ", c), frameWidth-1)
	}
	fmt.Println(border)
	fmt.Println()
}

// prepareSymbolDisplays generates display strings for symbols and their levels.
func (cd *PeakPursuitDisplay) prepareSymbolDisplays(symbols []peakpursuit.Symbol, game *peakpursuit.SlotMachine) ([]string, []string) {
	maxLevels := 0
	for _, symbol := range symbols {
		if len(symbol.Payouts) > 1 && len(symbol.Payouts) > maxLevels {
			maxLevels = len(symbol.Payouts)
		}
	}

	levelDisplays := make([]string, maxLevels)
	symbolDisplays := make([]string, len(symbols)-3)

	i := 0
	for _, symbol := range symbols {
		if len(symbol.Payouts) > 1 { // Only display symbols with levels
			stateSymbol := game.State.SymbolLevels[symbol.Representation]
			symbolDisplays[i] = symbol.Representation

			for level := 1; level <= maxLevels; level++ {
				multiplier := fmt.Sprintf(" %3dx ", symbol.Payouts[level-1])
				bgStart := "\033[32m"
				bgEnd := "\033[0m"
				normalGreen := "\033[33m"
				normalYellow := "\033[39m"

				// Apply color based on the level and max reached condition
				if stateSymbol.MaxReached {
					multiplier = bgStart + multiplier + bgEnd
				} else if level <= stateSymbol.CurrentLevel {
					multiplier = normalGreen + multiplier + bgEnd
				} else {
					multiplier = normalYellow + multiplier + bgEnd
				}

				levelDisplays[level-1] += multiplier + " "
			}
			i++
		}
	}

	return levelDisplays, symbolDisplays
}
