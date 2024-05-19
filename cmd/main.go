package main

import (
	"fmt"
	"go-luck-simulator/internal/display"
	"go-luck-simulator/internal/factory"
	"go-luck-simulator/internal/manager"
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
