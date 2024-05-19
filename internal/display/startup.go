package display

import (
	"fmt"
	"go-luck-simulator/internal/utils"
	"strings"
)

type StartUp struct{}

func (su *StartUp) Show(message string) {
	fmt.Println(message)
}

func (su *StartUp) ShowStartupInfo() {
	su.ClearScreen()
	frameWidth := 32
	contentWidth := frameWidth - 2

	border := strings.Repeat("=", frameWidth)
	fmt.Println(border)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintCentered("\033[39;1mGoLuckSimulator\033[0m", contentWidth+11)
	utils.PrintBlankLine()
	utils.PrintBlankLine()
	utils.PrintIntermediaryLine()
	utils.PrintCentered("🎲 Randomizer 🎲", contentWidth-2)
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	utils.PrintLeftAligned("1. Default", contentWidth)
	utils.PrintLeftAligned("2. LCG", contentWidth)
	utils.PrintLeftAligned("3. XorShift*", contentWidth)
	utils.PrintIntermediaryLine()
	utils.PrintCentered("🕹️  Games 🕹️", contentWidth+2)
	utils.PrintIntermediaryLine()
	utils.PrintBlankLine()
	utils.PrintLeftAligned("1. WildFruits", contentWidth)
	utils.PrintLeftAligned("2. PeakPursuit", contentWidth)

	fmt.Println(border)
	fmt.Println()
}

func (su *StartUp) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
