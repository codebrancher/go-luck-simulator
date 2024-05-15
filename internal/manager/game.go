package manager

import (
	"bufio"
	"fmt"
	"go-slot-machine/internal/display"
	"go-slot-machine/internal/engine/wildfruits"
	"os"
	"strconv"
	"strings"
	"time"
)

type GameManager struct {
	Game      *wildfruits.SlotMachine
	Display   display.Display
	AutoDelay time.Duration
}

func (gm *GameManager) Play() {
	inputReader := bufio.NewReader(os.Stdin)
	lastValidBet := ""
	gm.Display.ShowInfo(wildfruits.Symbols, gm.Game)
	for {
		gm.Display.Show("Ready? Enter command:")
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)

		// If input is empty and there's a last valid bet, re-use that bet
		if input == "" && lastValidBet != "" {
			input = lastValidBet
		}

		// Continue if still no input (no last bet available)
		if input == "" {
			continue
		}

		commandParts := strings.Fields(input)
		if len(commandParts) == 0 {
			continue
		}

		switch strings.ToLower(commandParts[0]) {
		case "exit":
			gm.Display.Show("Exiting game.")
			return
		case "auto":
			if len(commandParts) != 3 {
				gm.Display.Show("Invalid auto command. Usage: auto <spins> <bet amount>")
				continue
			}
			spins, err1 := strconv.Atoi(commandParts[1])
			bet, err2 := strconv.Atoi(commandParts[2])
			if err1 != nil || err2 != nil || spins <= 0 || bet <= 0 {
				gm.Display.Show("Invalid parameters for auto command.")
				continue
			}
			gm.AutomatedPlay(spins, bet)
		case "bet":
			if len(commandParts) != 2 {
				gm.Display.Show("Invalid bet command. Usage: bet <amount>")
				continue
			}
			if err := gm.handleManualPlay(commandParts[1]); err != nil {
				gm.Display.Show("Error: " + err.Error())
			} else {
				lastValidBet = input
			}
		case "info":
			gm.Display.ShowInfo(wildfruits.Symbols, gm.Game)
		case "stats":
			gm.Display.ShowStats(gm.Game)
		default:
			gm.Display.Show("Invalid command. Please use 'info', 'stats', 'bet', 'auto', or 'exit'.")
		}
	}
}

func (gm *GameManager) AutomatedPlay(totalSpins int, betAmount int) {
	betStr := strconv.Itoa(betAmount) + "\n"
	for i := 0; i < totalSpins; i++ {
		if _, err := gm.Game.RequestBet(betStr); err != nil {
			gm.Display.Show("Error: " + err.Error())
			break
		}
		if err := gm.Game.Spin(); err != nil {
			gm.Display.Show("Error: " + err.Error())
			break
		}
		gm.Game.DisplayResults()
		time.Sleep(gm.AutoDelay)
	}
}

func (gm *GameManager) RunSimulation(totalSpins, betAmount int) {
	gm.Display.Show("Running simulation...")
	gm.AutomatedPlay(totalSpins, betAmount)
	gm.Display.ShowStats(gm.Game)
}

func (gm *GameManager) handleManualPlay(betInput string) error {
	bet, err := gm.Game.RequestBet(betInput)
	if err != nil {
		return err
	}
	if bet == 0 && gm.Game.GameState.BonusGames == 0 {
		gm.Display.Show("No bet placed or invalid input; please try again.")
		return fmt.Errorf("no bet placed")
	}
	if err = gm.Game.Spin(); err != nil {
		return err
	}
	gm.Game.DisplayResults()
	return nil
}
