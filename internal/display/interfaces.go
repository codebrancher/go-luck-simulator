package display

import "go-luck-simulator/internal/engine/wildfruits"

type Display interface {
	Show(message string)
	ClearScreen()
	ShowInfo(symbols []wildfruits.Symbol, game *wildfruits.SlotMachine)
	ShowStats(game *wildfruits.SlotMachine)
	ShowStartupInfo(title string)
	ShowHelp()
}
