package display

import (
	"fmt"
	"go-slot-machine/internal/engine/wildfruits"
	"strings"
	"time"
)

type WildFruitsDisplay struct{}

func (cd *WildFruitsDisplay) Show(message string) {
	fmt.Println(message)
}

func (cd *WildFruitsDisplay) ShowStartupInfo(title string) {
	cd.ClearScreen()
	frameWidth := 32
	contentWidth := frameWidth - 2

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	printBlankLine()
	printCentered(title, contentWidth-2)
	printBlankLine()
	printIntermediaryLine()
	printCentered("Commands", contentWidth)
	printIntermediaryLine()
	printBlankLine()
	printLeftAligned("play", contentWidth)
	printBlankLine()
	printLeftAligned("sim [spins] [bet]", contentWidth)
	printBlankLine()
	printBlankLine()
	printLeftAligned("exit", contentWidth)
	fmt.Println(border)
	fmt.Println()
}

func (cd *WildFruitsDisplay) ShowHelp() {
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
	printLeftAligned("bet [amount]", contentWidth)
	printBlankLine()
	printLeftAligned("auto [spins] [bet amount]", contentWidth)
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

func (cd *WildFruitsDisplay) ShowInfo(symbols []wildfruits.Symbol, game *wildfruits.SlotMachine) {
	cd.ClearScreen()

	frameWidth := 32
	contentWidth := frameWidth - 2
	contentWidthUnicode := frameWidth - 4

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	printBlankLine()
	printCentered("The Go Slot Machine!", contentWidth)
	printCentered(game.Title, contentWidthUnicode)
	printBlankLine()
	printIntermediaryLine()

	// Explanation of Wild Symbol and Bonus Games
	wildSymbolDescription := fmt.Sprintf("3x %s for 7 Bonus Games!", game.GameConfig.WildSymbol)
	printCentered("Features", contentWidth)
	printIntermediaryLine()
	printCentered(wildSymbolDescription, contentWidth-1)
	printCentered(fmt.Sprintf("%s is wild!", game.GameConfig.WildSymbol), contentWidth-1)
	printIntermediaryLine()

	// High Value and Special Symbols
	printCentered("High & Special Value", contentWidth)
	printIntermediaryLine()
	for _, sym := range symbols {
		if sym.Payout >= 5 || sym.Representation == game.GameConfig.WildSymbol {
			repeatedSymbol := strings.Repeat(sym.Representation, 3)
			payoutInfo := fmt.Sprintf("%s %dx", repeatedSymbol, sym.Payout)
			printCentered(payoutInfo, contentWidthUnicode-1)
		}
	}

	printIntermediaryLine()

	// Regular Symbols
	printCentered("Regular Value", contentWidth)
	printIntermediaryLine()
	eligibleSymbols := make([]string, 0)

	for _, sym := range symbols {
		if sym.Payout < 5 && sym.Representation != game.GameConfig.WildSymbol {
			repeatedSymbol := strings.Repeat(sym.Representation, 3)
			payoutInfo := fmt.Sprintf("%s %dx", repeatedSymbol, sym.Payout)
			eligibleSymbols = append(eligibleSymbols, payoutInfo)
		}
	}
	// Print eligible symbols in pairs
	for i := 0; i < len(eligibleSymbols); i += 2 {
		if i+1 < len(eligibleSymbols) {
			printWithMiddlePadding("  "+eligibleSymbols[i], eligibleSymbols[i+1]+"  ", contentWidthUnicode-2)
		} else {
			printWithMiddlePadding("  "+eligibleSymbols[i], "", contentWidthUnicode-2)
		}
	}
	printIntermediaryLine()
	printCentered("Commands", contentWidth)
	printIntermediaryLine()
	printLeftAligned("bet [amount]", contentWidth)
	printLeftAligned("auto [spins] [bet amount]", contentWidth)
	printLeftAligned("help", contentWidth)
	printIntermediaryLine()
	printBlankLine()
	printCentered(fmt.Sprintf("Your cash: %d%s", game.GameState.Cash, game.GameConfig.Currency), contentWidth)
	printBlankLine()
	fmt.Println(border)
	fmt.Println()
}
func (cd *WildFruitsDisplay) ShowStats(game *wildfruits.SlotMachine) {
	cd.ClearScreen()
	frameWidth := 32
	contentWidth := frameWidth - 2

	var rtp, avgBetAmount float64
	totalBets := game.GameState.TotalSpins * game.GameState.LastBet

	if totalBets > 0 {
		rtp = float64(game.GameState.TotalWinAmount) / float64(game.GameState.TotalBetAmount) * 100
		avgBetAmount = float64(game.GameState.TotalBetAmount) / float64(game.GameState.TotalBets)
	} else {
		rtp = 0
	}
	averagePayout := float64(game.GameState.TotalWinAmount) / float64(game.GameState.TotalWins)
	hitFrequency := float64(game.GameState.TotalWins) / float64(game.GameState.TotalSpins) * 100
	profitability := game.GameState.Cash - game.GameState.StartingCash

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	printBlankLine()
	printCentered("Who doesn't like stats?", contentWidth)
	printBlankLine()
	printIntermediaryLine()
	printBlankLine()
	printWithMiddlePadding(" RTP", fmt.Sprintf("%.2f%% ", rtp), frameWidth)
	printWithMiddlePadding(" Win rate", fmt.Sprintf("%.2f%% ", hitFrequency), frameWidth)
	printWithMiddlePadding(" Playtime", fmt.Sprintf("%s ", formatDuration(time.Since(game.GameState.StartTime))), frameWidth)
	printIntermediaryLine()
	printBlankLine()
	printWithMiddlePadding(" Total spins", fmt.Sprintf("%d ", game.GameState.TotalSpins), frameWidth)
	printWithMiddlePadding(" Total wins", fmt.Sprintf("%d ", game.GameState.TotalWins), frameWidth)
	printWithMiddlePadding(" Bonus games initiated", fmt.Sprintf("%d ", game.DisplayConfig.BonusGamesInitiated), frameWidth)
	printIntermediaryLine()
	printBlankLine()
	printWithMiddlePadding(" Starting cash", fmt.Sprintf("%d%s ", game.GameState.StartingCash, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Current cash", fmt.Sprintf("%d%s ", game.GameState.Cash, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Total win amount", fmt.Sprintf("%d%s ", game.GameState.TotalWinAmount, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Total bet amount", fmt.Sprintf("%d%s ", game.GameState.TotalBetAmount, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Max drawdown", fmt.Sprintf("%d%s ", game.GameState.MaxDrawdown, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Net profit/loss", fmt.Sprintf("%d ", profitability), frameWidth)
	printWithMiddlePadding(" Average bet", fmt.Sprintf("%.0f%s ", avgBetAmount, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Average win", fmt.Sprintf("%.2f%s ", averagePayout, game.GameConfig.Currency), frameWidth)
	printWithMiddlePadding(" Volatility", fmt.Sprintf("%.2f ", game.CalculateWinDeviation()), frameWidth)
	printWithMiddlePadding(" Top win", fmt.Sprintf("%d%s ", game.GameState.TopWinAmount, game.GameConfig.Currency), frameWidth)
	fmt.Println(border)
	fmt.Println()

}

func (cd *WildFruitsDisplay) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (cd *WildFruitsDisplay) Update(state *wildfruits.WildFruitState) {
	cd.ClearScreen()

	frameWidth := 32               // Width includes the border "|"
	contentWidth := frameWidth - 2 // Width available for content between the borders
	contentWidthUnicode := frameWidth - 4

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	printBlankLine()
	printCentered("The Go Slot Machine!", contentWidth)
	printCentered(state.Title, contentWidthUnicode)
	printBlankLine()
	printIntermediaryLine()
	printBlankLine()
	printCentered("Bonus Symbol", contentWidth)
	printBlankLine()
	fmt.Printf("|            [ %s ]            |\n", state.BonusSymbol)

	if state.FreeGames > 0 {
		printBlankLine()
		printCentered(fmt.Sprintf("%d Bonus Games left!", state.FreeGames), contentWidth)
	}
	printBlankLine()
	fmt.Println("|       ================       |")
	for row := 0; row < 3; row++ {
		lineDisplay := "|       "
		for col := 0; col < 3; col++ {
			symbol := state.Wheels[row][col]
			if state.WinningPositions[row][col] {
				symbol = fmt.Sprintf("\033[48;5;22m %s \033[0m", symbol)
			} else {
				symbol = fmt.Sprintf(" %s ", symbol)
			}
			lineDisplay += fmt.Sprintf("|%s", symbol)
		}
		lineDisplay += "|       |"
		fmt.Println(lineDisplay)
	}
	fmt.Println("|       ================       |")
	printBlankLine()
	if state.TotalWinAmount > 0 {
		printCentered(fmt.Sprintf("\033[38;5;22m*** %d%s ***\033[0m", state.TotalWinAmount, state.Currency), contentWidth+14)
	} else {
		printCentered(fmt.Sprintf("*** %d%s ***", state.TotalWinAmount, state.Currency), contentWidth)
	}
	printBlankLine()

	// Displaying winning descriptions
	if state.WinningDescription != "" {
		printIntermediaryLine()
		for _, line := range strings.Split(state.WinningDescription, "\n") {
			printCentered(line, contentWidth-1)
		}
	}
	printIntermediaryLine()
	printBlankLine()
	leftValLine1 := fmt.Sprintf(" Bet: %d%s", state.LastBet, state.Currency)
	rightValLine1 := fmt.Sprintf("Top Win: %d%s ", state.TopWinAmount, state.Currency)
	printWithMiddlePadding(leftValLine1, rightValLine1, contentWidth+2)
	leftValLine2 := fmt.Sprintf(" Cash: %d%s", state.Cash, state.Currency)
	rightValLine2 := fmt.Sprintf("Win Rate: %d%% ", int(state.WinRate))
	printWithMiddlePadding(leftValLine2, rightValLine2, contentWidth+2)
	printBlankLine()
	fmt.Println(border)
	fmt.Println()
}
