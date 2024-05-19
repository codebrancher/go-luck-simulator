package display

import (
	"fmt"
	"go-luck-simulator/internal/engine/wildfruits"
	"go-luck-simulator/internal/utils"
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
	utils.PrintLeftAligned("sim [spins] [bet]", contentWidth)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintLeftAligned("exit", contentWidth)
	fmt.Println(border)
	fmt.Println()
}

func (cd *WildFruitsDisplay) ShowHelp() {
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
	utils.PrintLeftAligned("bet [amount]", contentWidth)
	utils.PrintBlankLine()
	utils.PrintLeftAligned("auto [spins] [bet amount]", contentWidth)
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

func (cd *WildFruitsDisplay) ShowInfo(symbols []wildfruits.Symbol, game *wildfruits.SlotMachine) {
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
	wildSymbolDescription := fmt.Sprintf("3x %s for 7 Bonus Games!", game.GameConfig.WildSymbol)
	utils.PrintCentered("Features", contentWidth)
	utils.PrintIntermediaryLine()
	utils.PrintCentered(wildSymbolDescription, contentWidth-1)
	utils.PrintCentered(fmt.Sprintf("%s is wild!", game.GameConfig.WildSymbol), contentWidth-1)
	utils.PrintIntermediaryLine()

	// High Value and Special Symbols
	utils.PrintCentered("High & Special Value", contentWidth)
	utils.PrintIntermediaryLine()
	for _, sym := range symbols {
		if sym.Payout >= 5 || sym.Representation == game.GameConfig.WildSymbol {
			repeatedSymbol := strings.Repeat(sym.Representation, 3)
			payoutInfo := fmt.Sprintf("%s %dx", repeatedSymbol, sym.Payout)
			utils.PrintCentered(payoutInfo, contentWidthUnicode-1)
		}
	}

	utils.PrintIntermediaryLine()

	// Regular Symbols
	utils.PrintCentered("Regular Value", contentWidth)
	utils.PrintIntermediaryLine()
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
			utils.PrintWithMiddlePadding("  "+eligibleSymbols[i], eligibleSymbols[i+1]+"  ", contentWidthUnicode-2)
		} else {
			utils.PrintWithMiddlePadding("  "+eligibleSymbols[i], "", contentWidthUnicode-2)
		}
	}
	utils.PrintIntermediaryLine()
	utils.PrintCentered("Commands", contentWidth)
	utils.PrintIntermediaryLine()
	utils.PrintLeftAligned("bet [amount]", contentWidth)
	utils.PrintLeftAligned("auto [spins] [bet amount]", contentWidth)
	utils.PrintLeftAligned("help", contentWidth)
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	utils.PrintCentered(fmt.Sprintf("Your cash: %d%s", game.State.Amount, game.GameConfig.Currency), contentWidth)
	utils.PrintBlankLine()
	fmt.Println(border)
	fmt.Println()
}
func (cd *WildFruitsDisplay) ShowStats(game *wildfruits.SlotMachine) {
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
	utils.PrintWithMiddlePadding(" Bonus games initiated", fmt.Sprintf("%d ", game.DisplayConfig.BonusGamesInitiated), frameWidth)
	fmt.Println(border)
	fmt.Println()

}

func (cd *WildFruitsDisplay) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (cd *WildFruitsDisplay) Update(state *wildfruits.WildFruitObserverState) {
	cd.ClearScreen()

	frameWidth := 32               // Width includes the border "|"
	contentWidth := frameWidth - 2 // Width available for content between the borders
	contentWidthUnicode := frameWidth - 4

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintCentered(state.Title, contentWidthUnicode)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	utils.PrintCentered("Bonus Symbol", contentWidth)
	utils.PrintBlankLine()
	fmt.Printf("|            [ %s ]            |\n", state.BonusSymbol)

	if state.BonusGames > 0 {
		utils.PrintBlankLine()
		utils.PrintCentered(fmt.Sprintf("%d Bonus Games left!", state.BonusGames), contentWidth)
	}
	utils.PrintBlankLine()
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
	utils.PrintBlankLine()
	if state.TotalWinAmountPerSpin > 0 {
		utils.PrintCentered(fmt.Sprintf("\033[38;5;22m*** %d%s ***\033[0m", state.TotalWinAmountPerSpin, state.Currency), contentWidth+14)
	} else {
		utils.PrintCentered(fmt.Sprintf("*** %d%s ***", state.TotalWinAmountPerSpin, state.Currency), contentWidth)
	}
	utils.PrintBlankLine()

	// Displaying winning descriptions
	if state.WinningDescription != "" {
		utils.PrintIntermediaryLine()
		for _, line := range strings.Split(state.WinningDescription, "\n") {
			utils.PrintCentered(line, contentWidth-1)
		}
	}
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	leftValLine1 := fmt.Sprintf(" Bet: %d%s", state.LastBet, state.Currency)
	rightValLine1 := fmt.Sprintf("Top Win: %d%s ", state.TopWinAmount, state.Currency)
	utils.PrintWithMiddlePadding(leftValLine1, rightValLine1, contentWidth+2)
	leftValLine2 := fmt.Sprintf(" Cash: %d%s", state.Amount, state.Currency)
	rightValLine2 := fmt.Sprintf("Win Rate: %d%% ", int(state.WinRate))
	utils.PrintWithMiddlePadding(leftValLine2, rightValLine2, contentWidth+2)
	utils.PrintBlankLine()
	fmt.Println(border)
	fmt.Println()
}
