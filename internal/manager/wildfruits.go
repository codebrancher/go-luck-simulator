package manager

import (
	"bufio"
	"fmt"
	"go-luck-simulator/internal/display"
	"go-luck-simulator/internal/engine/wildfruits"
	"os"
	"strconv"
	"strings"
	"time"
)

type WildFruitsManager struct {
	Game        *wildfruits.SlotMachine
	Display     display.Display
	AutoDelay   time.Duration
	VisualDelay time.Duration
}

func NewWildFruitsManager(game *wildfruits.SlotMachine, display display.Display) *WildFruitsManager {
	return &WildFruitsManager{
		Game:        game,
		Display:     display,
		AutoDelay:   0,
		VisualDelay: 0,
	}
}

func (gm *WildFruitsManager) Start() {
	reader := bufio.NewReader(os.Stdin)
	gm.Display.ShowStartupInfo(gm.Game.Title)
	for {
		fmt.Print("Enter command: ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		args := strings.Fields(command)

		if len(args) == 0 {
			continue
		}

		switch strings.ToLower(args[0]) {
		case "play":
			gm.VisualDelay = 100 * time.Millisecond
			gm.AutoDelay = 1 * time.Second
			gm.Game.SetVisualDelay(gm.VisualDelay)
			gm.Game.EnableObserver()
			gm.Play()
			return
		case "sim":
			if len(args) != 3 {
				fmt.Println("Invalid command. Usage: sim [spins] [amount]")
				continue
			}
			spins, err1 := strconv.Atoi(args[1])
			amount, err2 := strconv.Atoi(args[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid parameters. Ensure both spins and amount are numbers.")
				continue
			}
			gm.VisualDelay = 0
			gm.AutoDelay = 0
			gm.Game.SetVisualDelay(gm.VisualDelay)
			gm.Game.DisableObserver()
			gm.RunSimulation(spins, amount)
			return
		case "exit":
			gm.Display.Show("Exiting game.")
			return
		default:
			fmt.Println("Invalid command. Use 'play' or 'sim [spins] [amount]'.")
		}
	}
}

func (gm *WildFruitsManager) Play() {
	inputReader := bufio.NewReader(os.Stdin)
	lastValidBet := ""
	gm.Display.ShowInfo(wildfruits.Symbols, gm.Game)
	for {
		fmt.Print("Enter command: ")
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
		case "help":
			gm.Display.ShowHelp()
		default:
			gm.Display.ShowHelp()
		}
	}
}

func (gm *WildFruitsManager) AutomatedPlay(totalSpins int, betAmount int) {
	betStr := strconv.Itoa(betAmount) + "\n"
	for i := 0; i < totalSpins; i++ {
		if _, err := gm.Game.RequestCommand(betStr); err != nil {
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

func (gm *WildFruitsManager) RunSimulation(totalSpins, betAmount int) {
	gm.Display.Show("Running simulation...")
	gm.AutomatedPlay(totalSpins, betAmount)
	gm.Display.ShowStats(gm.Game)
}

func (gm *WildFruitsManager) handleManualPlay(betInput string) error {
	bet, err := gm.Game.RequestCommand(betInput)
	if err != nil {
		return err
	}
	if bet == 0 && gm.Game.State.BonusGames == 0 {
		gm.Display.Show("No bet placed or invalid input; please try again.")
		return fmt.Errorf("no bet placed")
	}
	if err = gm.Game.Spin(); err != nil {
		return err
	}
	gm.Game.DisplayResults()
	return nil
}
