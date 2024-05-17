package manager

import (
	"bufio"
	"fmt"
	"go-slot-machine/internal/display"
	"go-slot-machine/internal/engine/peakpursuit"
	"os"
	"strconv"
	"strings"
	"time"
)

type PeakPursuitManager struct {
	Game        *peakpursuit.SlotMachine
	Display     *display.PeakPursuitDisplay
	AutoDelay   time.Duration
	VisualDelay time.Duration
}

func NewPeakPursuitManager(game *peakpursuit.SlotMachine, display *display.PeakPursuitDisplay) *PeakPursuitManager {
	return &PeakPursuitManager{
		Game:        game,
		Display:     display,
		AutoDelay:   0,
		VisualDelay: 0,
	}
}

func (pm *PeakPursuitManager) Start() {
	reader := bufio.NewReader(os.Stdin)
	pm.Display.ShowStartupInfo(pm.Game.Title)
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
			pm.VisualDelay = 50 * time.Millisecond
			pm.AutoDelay = 1 * time.Second
			pm.Game.SetVisualDelay(pm.VisualDelay)
			pm.Game.EnableObserver()
			pm.Play()
			return
		case "sim":
			if len(args) != 3 {
				pm.Display.Show("Invalid command. Usage: sim [spins] [amount]")
				continue
			}
			spins, err1 := strconv.Atoi(args[1])
			amount, err2 := strconv.Atoi(args[2])
			if err1 != nil || err2 != nil {
				pm.Display.Show("Invalid parameters. Ensure both spins and amount are numbers.")
				continue
			}
			pm.VisualDelay = 0
			pm.AutoDelay = 0
			pm.Game.SetVisualDelay(pm.VisualDelay)
			pm.Game.DisableObserver()
			pm.RunSimulation(spins, amount)
			return
		case "exit":
			pm.Display.Show("Exiting game.")
			return
		default:
			pm.Display.Show("Invalid command. Use 'play' to start the game.")
		}
	}
}

func (pm *PeakPursuitManager) Play() {
	inputReader := bufio.NewReader(os.Stdin)
	lastValidBet := ""
	pm.Display.ShowInfo(peakpursuit.Symbols, pm.Game)
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
			pm.Display.Show("Exiting game.")
			return
		case "cashout":
			_, err := pm.Game.RequestCommand("cashout")
			if err != nil {
				pm.Display.Show("Error: " + err.Error())
			} else {
				pm.Display.Show("")
			}
		case "bet":
			if len(commandParts) != 2 {
				pm.Display.Show("Invalid bet command. Usage: bet <amount>")
				continue
			}
			if err := pm.handleManualPlay(commandParts[1]); err != nil {
				pm.Display.Show("Error: " + err.Error())
			} else {
				lastValidBet = input
			}
		case "stats":
			pm.Display.ShowStats(pm.Game)
		case "info":
			pm.Display.ShowInfo(peakpursuit.Symbols, pm.Game)
		default:
			pm.Display.ShowHelp()
		}
	}
}

func (pm *PeakPursuitManager) handleManualPlay(betInput string) error {
	bet, err := pm.Game.RequestCommand(betInput)
	if err != nil {
		return err
	}
	if bet == 0 {
		pm.Display.Show("No bet placed or invalid input; please try again.")
		return fmt.Errorf("no bet placed")
	}
	if err = pm.Game.Spin(); err != nil {
		return err
	}
	pm.Game.DisplayResults()
	return nil
}

func (pm *PeakPursuitManager) AutomatedPlay(totalSpins int, betAmount int) {
	betStr := strconv.Itoa(betAmount) + "\n"
	for i := 0; i < totalSpins; i++ {
		if _, err := pm.Game.RequestCommand(betStr); err != nil {
			pm.Display.Show("Error: " + err.Error())
			break
		}
		if err := pm.Game.Spin(); err != nil {
			pm.Display.Show("Error: " + err.Error())
			break
		}
		pm.Game.DisplayResults()
		time.Sleep(pm.AutoDelay)
	}
}

func (pm *PeakPursuitManager) RunSimulation(totalSpins, betAmount int) {
	pm.Display.Show("Running simulation...")
	pm.AutomatedPlay(totalSpins, betAmount)
	pm.Display.ShowStats(pm.Game)
}
