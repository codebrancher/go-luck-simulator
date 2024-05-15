package display

import "go-slot-machine/internal/engine/wildfruits"

type Display interface {
	Show(message string)
	ShowWheels(wheels [3][3]string)
	ClearScreen()
	ShowInfo(symbols []wildfruits.Symbol, initialCash int, wildSymbol string)
	ShowStats(game *wildfruits.SlotMachine, startingCash int)
}
