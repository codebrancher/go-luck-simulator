package manager

import (
	"bufio"
	"fmt"
	"go-slot-machine/internal/display"
	"os"
	"strconv"
	"strings"
)

type GameManager struct {
	Display *display.StartUp
}

func NewGameManager(display *display.StartUp) *GameManager {
	return &GameManager{
		Display: display,
	}
}

func (gm *GameManager) Select() (string, string, int) {
	reader := bufio.NewReader(os.Stdin)
	gm.Display.ShowStartupInfo()

	// Step 1: Choose the randomizer
	fmt.Print("Select randomizer (number): ")
	rng, _ := reader.ReadString('\n')
	rng = strings.TrimSpace(rng)
	rng = strings.ToLower(rng)

	// Validate game choice
	if rng != "1" && rng != "2" && rng != "3" {
		gm.Display.Show("Invalid randomizer choice.")
		return "", "", 0
	}

	// Step 2: Choose the game
	fmt.Print("Select game (number): ")
	gameType, _ := reader.ReadString('\n')
	gameType = strings.TrimSpace(gameType)
	gameType = strings.ToLower(gameType)

	// Validate game choice
	if gameType != "1" && gameType != "2" {
		gm.Display.Show("Invalid game choice.")
		return "", "", 0
	}

	// Step 3: Enter starting cash
	fmt.Print("Enter starting cash: ")
	cashInput, _ := reader.ReadString('\n')
	startingCash, err := strconv.Atoi(strings.TrimSpace(cashInput))
	if err != nil || startingCash < 0 {
		gm.Display.Show("Invalid cash amount.")
		return "", "", 0
	}

	return rng, gameType, startingCash
}
