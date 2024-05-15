package display

import (
	"fmt"
	"go-slot-machine/internal/engine"
	"go-slot-machine/internal/engine/wildfruits"
	"strings"
	"time"
)

type ConsoleDisplay struct{}

func (cd *ConsoleDisplay) Show(message string) {
	fmt.Println(message)
}

func (cd *ConsoleDisplay) ShowStartupInfo() {
	cd.ClearScreen()
	frameWidth := 32
	contentWidth := frameWidth - 2

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	cd.printBlankLine()
	cd.printCentered("The Go Slot Machine!", contentWidth)
	cd.printBlankLine()
	cd.printIntermediaryLine()
	cd.printCentered("Commands", contentWidth)
	cd.printIntermediaryLine()
	cd.printBlankLine()
	cd.printLeftAligned("play", contentWidth)
	cd.printBlankLine()
	cd.printLeftAligned("sim [spins] [bet]", contentWidth)
	fmt.Println(border)
	fmt.Println()
}

func (cd *ConsoleDisplay) ShowWheels(wheels [3][3]string) {}

func (cd *ConsoleDisplay) ShowInfo(symbols []wildfruits.Symbol, game *wildfruits.SlotMachine) {
	cd.ClearScreen()

	frameWidth := 32
	contentWidth := frameWidth - 2
	contentWidthUnicode := frameWidth - 4

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	cd.printBlankLine()
	cd.printCentered("The Go Slot Machine!", contentWidth)
	cd.printCentered(game.Title, contentWidthUnicode)
	cd.printBlankLine()
	cd.printIntermediaryLine()

	// Explanation of Wild Symbol and Bonus Games
	wildSymbolDescription := fmt.Sprintf("3x %s for 7 Bonus Games!", game.GameConfig.WildSymbol)
	cd.printCentered("Features", contentWidth)
	cd.printIntermediaryLine()
	cd.printCentered(wildSymbolDescription, contentWidth-1)
	cd.printCentered(fmt.Sprintf("%s is wild!", game.GameConfig.WildSymbol), contentWidth-1)
	cd.printIntermediaryLine()

	// High Value and Special Symbols
	cd.printCentered("High & Special Value", contentWidth)
	cd.printIntermediaryLine()
	for _, sym := range symbols {
		if sym.Payout >= 5 || sym.Representation == game.GameConfig.WildSymbol {
			repeatedSymbol := strings.Repeat(sym.Representation, 3)
			payoutInfo := fmt.Sprintf("%s %dx", repeatedSymbol, sym.Payout)
			cd.printCentered(payoutInfo, contentWidthUnicode-1)
		}
	}

	cd.printIntermediaryLine()

	// Regular Symbols
	cd.printCentered("Regular Value", contentWidth)
	cd.printIntermediaryLine()
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
			cd.printWithMiddlePadding("  "+eligibleSymbols[i], eligibleSymbols[i+1]+"  ", contentWidthUnicode-2)
		} else {
			cd.printWithMiddlePadding("  "+eligibleSymbols[i], "", contentWidthUnicode-2)
		}
	}

	cd.printIntermediaryLine()
	cd.printBlankLine()
	cd.printCentered(fmt.Sprintf("Your cash: %d%s", game.GameState.Cash, game.GameConfig.Currency), contentWidth)
	cd.printBlankLine()
	fmt.Println(border)
	fmt.Println()
}
func (cd *ConsoleDisplay) ShowStats(game *wildfruits.SlotMachine) {
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
	cd.printBlankLine()
	cd.printCentered("Who doesn't like stats?", contentWidth)
	cd.printBlankLine()
	cd.printIntermediaryLine()
	cd.printBlankLine()
	cd.printWithMiddlePadding(" RTP", fmt.Sprintf("%.2f%% ", rtp), frameWidth)
	cd.printWithMiddlePadding(" Win rate", fmt.Sprintf("%.2f%% ", hitFrequency), frameWidth)
	cd.printWithMiddlePadding(" Playtime", fmt.Sprintf("%s ", formatDuration(time.Since(game.GameState.StartTime))), frameWidth)
	cd.printIntermediaryLine()
	cd.printBlankLine()
	cd.printWithMiddlePadding(" Total spins", fmt.Sprintf("%d ", game.GameState.TotalSpins), frameWidth)
	cd.printWithMiddlePadding(" Total wins", fmt.Sprintf("%d ", game.GameState.TotalWins), frameWidth)
	cd.printWithMiddlePadding(" Bonus games initiated", fmt.Sprintf("%d ", game.DisplayConfig.BonusGamesInitiated), frameWidth)
	cd.printIntermediaryLine()
	cd.printBlankLine()
	cd.printWithMiddlePadding(" Starting cash", fmt.Sprintf("%d%s ", game.GameState.StartingCash, game.GameConfig.Currency), frameWidth)
	cd.printWithMiddlePadding(" Current cash", fmt.Sprintf("%d%s ", game.GameState.Cash, game.GameConfig.Currency), frameWidth)
	cd.printWithMiddlePadding(" Total win amount", fmt.Sprintf("%d%s ", game.GameState.TotalWinAmount, game.GameConfig.Currency), frameWidth)
	cd.printWithMiddlePadding(" Total bet amount", fmt.Sprintf("%d%s ", game.GameState.TotalBetAmount, game.GameConfig.Currency), frameWidth)
	cd.printWithMiddlePadding(" Max drawdown", fmt.Sprintf("%d%s ", game.GameState.MaxDrawdown, game.GameConfig.Currency), frameWidth)
	cd.printWithMiddlePadding(" Net profit/loss", fmt.Sprintf("%d ", profitability), frameWidth)
	cd.printWithMiddlePadding(" Average bet", fmt.Sprintf("%.0f%s ", avgBetAmount, game.GameConfig.Currency), frameWidth)
	cd.printWithMiddlePadding(" Average win", fmt.Sprintf("%.2f%s ", averagePayout, game.GameConfig.Currency), frameWidth)
	cd.printWithMiddlePadding(" Volatility", fmt.Sprintf("%.2f ", game.CalculateWinDeviation()), frameWidth)
	cd.printWithMiddlePadding(" Top win", fmt.Sprintf("%d%s ", game.GameState.TopWinAmount, game.GameConfig.Currency), frameWidth)
	fmt.Println(border)
	fmt.Println()

}

func (cd *ConsoleDisplay) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (cd *ConsoleDisplay) Update(state engine.DisplayState) {
	cd.ClearScreen()

	frameWidth := 32               // Width includes the border "|"
	contentWidth := frameWidth - 2 // Width available for content between the borders
	contentWidthUnicode := frameWidth - 4

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	cd.printBlankLine()
	cd.printCentered("The Go Slot Machine!", contentWidth)
	cd.printCentered(state.Title, contentWidthUnicode)
	cd.printBlankLine()

	cd.printIntermediaryLine()
	cd.printBlankLine()
	cd.printCentered("Bonus Symbol", contentWidth)
	cd.printBlankLine()
	fmt.Printf("|            [ %s ]            |\n", state.BonusSymbol)

	if state.FreeGames > 0 {
		cd.printBlankLine()
		cd.printCentered(fmt.Sprintf("%d Bonus Games left!", state.FreeGames), contentWidth)
	}
	cd.printBlankLine()
	fmt.Println("|       ================       |")
	for row := 0; row < 3; row++ {
		lineDisplay := "|       "
		for col := 0; col < 3; col++ {
			symbol := state.Wheels[row][col]
			if state.WinningPositions[row][col] {
				// Highlight winning symbols, here using ANSI color codes for green background
				symbol = fmt.Sprintf("\033[48;5;22m %s \033[0m", symbol) // 156 is a light green in 256-color palette
			} else {
				symbol = fmt.Sprintf(" %s ", symbol)
			}
			lineDisplay += fmt.Sprintf("|%s", symbol)
		}
		lineDisplay += "|       |"
		fmt.Println(lineDisplay)
	}
	fmt.Println("|       ================       |")
	cd.printBlankLine()
	if state.TotalWinAmount > 0 {
		cd.printCentered(fmt.Sprintf("\033[38;5;22m*** %d%s ***\033[0m", state.TotalWinAmount, state.Currency), contentWidth+14)
	} else {
		cd.printCentered(fmt.Sprintf("*** %d%s ***", state.TotalWinAmount, state.Currency), contentWidth)
	}
	cd.printBlankLine()

	// Displaying winning descriptions
	if state.WinningDescription != "" {
		cd.printIntermediaryLine()
		for _, line := range strings.Split(state.WinningDescription, "\n") {
			cd.printCentered(line, contentWidth-1)
		}
	}
	cd.printIntermediaryLine()
	cd.printBlankLine()
	leftValLine1 := fmt.Sprintf(" Bet: %d%s", state.LastBet, state.Currency)
	rightValLine1 := fmt.Sprintf("Top Win: %d%s ", state.TopWinAmount, state.Currency)
	cd.printWithMiddlePadding(leftValLine1, rightValLine1, contentWidth+2)
	leftValLine2 := fmt.Sprintf(" Cash: %d%s", state.Cash, state.Currency)
	rightValLine2 := fmt.Sprintf("Win Rate: %d%% ", int(state.WinRate))
	cd.printWithMiddlePadding(leftValLine2, rightValLine2, contentWidth+2)
	cd.printBlankLine()
	fmt.Println(border)
	fmt.Println()
}
