package display

import (
	"fmt"
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
	printBlankLine()
	printBlankLine()
	printCentered("\033[39;1mThe Go Slot Machine!\033[0m", contentWidth+11)
	printBlankLine()
	printBlankLine()
	printIntermediaryLine()
	printCentered("🎲 Randomizer 🎲", contentWidth-2)
	printIntermediaryLine()
	printBlankLine()
	printLeftAligned("1. Default", contentWidth)
	printLeftAligned("2. LCG", contentWidth)
	printLeftAligned("3. XorShift*", contentWidth)
	printIntermediaryLine()
	printCentered("🕹️  Games 🕹️", contentWidth+2)
	printIntermediaryLine()
	printBlankLine()
	printLeftAligned("1. WildFruits", contentWidth)
	printLeftAligned("2. PeakPursuit", contentWidth)

	fmt.Println(border)
	fmt.Println()
}

func (su *StartUp) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
