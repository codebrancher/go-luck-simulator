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
	printBlankLine()
	printBlankLine()
	printCentered(title, contentWidth-2)
	printBlankLine()
	printBlankLine()
	printIntermediaryLine()
	printCentered("Commands", contentWidth)
	printIntermediaryLine()
	printBlankLine()
	printLeftAligned("play", contentWidth)
	printBlankLine()
	printLeftAligned("sim [spins] [bet amount]", contentWidth)
	printBlankLine()
	printBlankLine()
	printLeftAligned("exit", contentWidth)
	fmt.Println(border)
	fmt.Println()
}

func (cd *PeakPursuitDisplay) ShowHelp() {
	cd.ClearScreen()
	frameWidth := 32
	contentWidth := frameWidth - 2

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	printBlankLine()
	printCentered("Need help?", contentWidth)
	printBlankLine()
	printIntermediaryLine()
	printCentered("Commands", contentWidth)
	printIntermediaryLine()
	printBlankLine()
	printLeftAligned("cashout", contentWidth)
	printBlankLine()
	printLeftAligned("bet [amount]", contentWidth)
	printBlankLine()
	printBlankLine()
	printLeftAligned("info", contentWidth)
	printBlankLine()
	printLeftAligned("stats", contentWidth)
	printBlankLine()
	printLeftAligned("help", contentWidth)
	printBlankLine()
	printLeftAligned("exit", contentWidth)
	fmt.Println(border)
	fmt.Println()
}

func (cd *PeakPursuitDisplay) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (cd *PeakPursuitDisplay) Update(state *peakpursuit.PeakPursuitState) {
	cd.ClearScreen()

	frameWidth := 32               // Width includes the border "|"
	contentWidth := frameWidth - 2 // Width available for content between the borders
	contentWidthUnicode := frameWidth - 4

	rtp := float64(state.TotalWinAmount) / float64(state.TotalBetAmount) * 100

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	printBlankLine()
	printBlankLine()
	printCentered(state.Title, contentWidthUnicode)
	printBlankLine()
	printBlankLine()
	printIntermediaryLine()
	printCentered("Pot", contentWidth)
	printIntermediaryLine()
	printBlankLine()
	if state.Winnings > 0 {
		printCentered("\033[48;5;22m      "+strconv.Itoa(state.Winnings)+"      \033[0m", contentWidth+14)
	} else {
		printCentered(strconv.Itoa(state.Winnings), contentWidth)
	}
	printBlankLine()

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
				bgStart := "\033[32m" // Green with blink
				bgEnd := "\033[0m"
				normalGreen := "\033[33m"  // Green without blink
				normalYellow := "\033[39m" // Yellow

				// Apply color based on the level and max reached condition
				if stateSymbol.MaxReached {
					// Blink all levels with green background when max level is reached
					multiplier = bgStart + multiplier + bgEnd
				} else if level <= stateSymbol.CurrentLevel {
					// Green without blink up to the current level
					multiplier = normalGreen + multiplier + bgEnd
				} else {
					// Yellow for levels above the current one
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
	printBlankLine()
	fmt.Printf("|       %s       |\n", strings.Join(symbolDisplays, "     "))
	printBlankLine()
	printIntermediaryLine()
	printCentered("------", contentWidth)
	fmt.Printf("|            | %s |            |\n", state.SpinSymbolRepresentation)
	printCentered("------", contentWidth)
	// Displaying winning descriptions
	if state.WinningDescription != "" {
		printIntermediaryLine()
		for _, line := range strings.Split(state.WinningDescription, "\n") {
			printCentered(line, contentWidth-2)
		}
	}
	printIntermediaryLine()
	printBlankLine()
	leftValLine1 := fmt.Sprintf(" Bet: %d%s", state.LastBet, state.Currency)
	rightValLine1 := fmt.Sprintf("RTP: %.1f%% ", rtp)
	printWithMiddlePadding(leftValLine1, rightValLine1, contentWidth+2)
	leftValLine2 := fmt.Sprintf(" Cash: %d%s", state.Cash, state.Currency)
	printWithMiddlePadding(leftValLine2, "", contentWidth+2)
	printBlankLine()
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
	printBlankLine()
	printBlankLine()
	printCentered(game.Title, contentWidthUnicode)
	printBlankLine()
	printBlankLine()
	printIntermediaryLine()

	// Explanation of Wild Symbol and Bonus Games
	resetSymbolDescription := fmt.Sprintf("%s resets game state", game.GameConfig.ResetSymbolRepresentation)
	wildSymbolDescription := fmt.Sprintf("%s increases every symbol", game.GameConfig.WildSymbolRepresentation)
	printCentered("Features", contentWidth)
	printIntermediaryLine()
	printCentered(resetSymbolDescription, contentWidth-1)
	printCentered(wildSymbolDescription, contentWidth-1)

	printIntermediaryLine()
	printBlankLine()
	levelDisplays, symbolDisplays := cd.prepareSymbolDisplays(symbols, game)

	// Print each level from highest to lowest (to appear as a rising barometer)
	for j := len(levelDisplays) - 1; j >= 0; j-- {
		fmt.Printf("|    %s     |\n", levelDisplays[j])

	}
	printBlankLine()
	fmt.Printf("|       %s       |\n", strings.Join(symbolDisplays, "     "))
	printBlankLine()

	printIntermediaryLine()
	printCentered("Commands", contentWidth)
	printIntermediaryLine()
	printLeftAligned("cashout", contentWidth)
	printLeftAligned("bet [amount]", contentWidth)
	printLeftAligned("help", contentWidth)
	printIntermediaryLine()
	printBlankLine()
	printCentered(fmt.Sprintf("Your cash: %d%s", game.GameState.Cash, game.GameConfig.Currency), contentWidth)
	printBlankLine()
	fmt.Println(border)
	fmt.Println()
}

func (cd *PeakPursuitDisplay) ShowStats(game *peakpursuit.SlotMachine) {
	cd.ClearScreen()
	frameWidth := 32
	contentWidth := frameWidth - 2

	rtp := utils.CalculateRTP(game.GameState.TotalWinAmount, game.GameState.TotalBetAmount)
	profitability := utils.CalculateProfitability(game.GameState.Cash, game.GameState.StartingCash)
	avgBetAmount := utils.CalculateAverageBet(game.GameState.TotalBetAmount, game.GameState.TotalBets)
	avgPayout := utils.CalculateAveragePayout(game.GameState.TotalWinAmount, game.GameState.TotalWins)
	deviation := utils.CalculateWinDeviation(game.GameState.Wins)

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	printBlankLine()
	printCentered("Who doesn't like stats?", contentWidth)
	printBlankLine()
	printIntermediaryLine()
	printBlankLine()
	printWithMiddlePadding(" RTP", fmt.Sprintf("%.2f%% ", rtp), frameWidth)
	printWithMiddlePadding(" Win rate", fmt.Sprintf("%.2f%% ", game.GameState.WinRate), frameWidth)
	printWithMiddlePadding(" Playtime", fmt.Sprintf("%s ", formatDuration(time.Since(game.GameState.StartTime))), frameWidth)
	printIntermediaryLine()
	printBlankLine()
	printWithMiddlePadding(" Total spins", fmt.Sprintf("%d ", game.GameState.TotalSpins), frameWidth)
	printWithMiddlePadding(" Total wins", fmt.Sprintf("%d ", game.GameState.TotalWins), frameWidth)
	printIntermediaryLine()
	printBlankLine()
	printWithMiddlePadding(" Starting cash", fmt.Sprintf("%d%s ", game.GameState.StartingCash, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Current cash", fmt.Sprintf("%d%s ", game.GameState.Cash, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Total win amount", fmt.Sprintf("%d%s ", game.GameState.TotalWinAmount, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Total bet amount", fmt.Sprintf("%d%s ", game.GameState.TotalBetAmount, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Max drawdown", fmt.Sprintf("%d%s ", game.GameState.MaxDrawdown, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Net profit/loss", fmt.Sprintf("%d ", profitability), frameWidth)
	printWithMiddlePadding(" Average bet", fmt.Sprintf("%.2f%s ", avgBetAmount, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Average win", fmt.Sprintf("%.2f%s ", avgPayout, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Volatility", fmt.Sprintf("%.2f ", deviation), frameWidth)
	printWithMiddlePadding(" Top win", fmt.Sprintf("%d%s ", game.GameState.TopWinAmount, game.GameConfig.Currency), frameWidth)
	printIntermediaryLine()
	printBlankLine()
	printWithMiddlePadding(" Best Spin Streak", fmt.Sprintf("%d ", game.GameState.LongestSpinStreak), frameWidth)
	printWithMiddlePadding(" Max Level Reached", fmt.Sprintf("%d ", game.GameState.MaxLevelReached), frameWidth)
	printWithMiddlePadding(" Best Max Level Streak", fmt.Sprintf("%d ", game.GameState.LongestMaxLevelStreak), frameWidth)
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
			stateSymbol := game.GameState.SymbolLevels[symbol.Representation]
			symbolDisplays[i] = symbol.Representation // Prepare symbol display line

			for level := 1; level <= maxLevels; level++ {
				multiplier := fmt.Sprintf(" %3dx ", symbol.Payouts[level-1])
				bgStart := "\033[32m" // Green with blink
				bgEnd := "\033[0m"
				normalGreen := "\033[33m"  // Green without blink
				normalYellow := "\033[39m" // Yellow

				// Apply color based on the level and max reached condition
				if stateSymbol.MaxReached {
					// Blink all levels with green background when max level is reached
					multiplier = bgStart + multiplier + bgEnd
				} else if level <= stateSymbol.CurrentLevel {
					// Green without blink up to the current level
					multiplier = normalGreen + multiplier + bgEnd
				} else {
					// Yellow for levels above the current one
					multiplier = normalYellow + multiplier + bgEnd
				}

				levelDisplays[level-1] += multiplier + " "
			}
			i++
		}
	}

	return levelDisplays, symbolDisplays
}
