package main

import (
	"fmt"
	"go-slot-machine/internal/display"
	"go-slot-machine/internal/factory"
	"go-slot-machine/internal/manager"
)

func main() {
	currency := ""

	disp := &display.StartUp{}
	m := manager.NewGameManager(disp)
	rngType, gameType, startingCash := m.Select()

	r, err := factory.NewRandomizer(rngType)

	g, err := factory.NewGame(gameType, r, startingCash, currency)
	if err != nil {
		fmt.Println("Error creating game:", err)
		return
	}

	g.Start()
}
