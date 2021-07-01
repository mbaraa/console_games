package main

import (
	"fmt"
	"os"
	"strconv"
)

type GameInput struct {
	data       *GameData
	cellNumber int
	input      string
}

func NewGameInput(data *GameData) *GameInput {
	return &GameInput{data: data}
}

func (gi *GameInput) GetInput() {
	fmt.Printf("Enter cell number for player %d (enter 0 to exit): ",
		gi.data.GetCurrentPlayer())

	fmt.Scan(&gi.input)

	if !gi.checkInput() {
		return
	}

	gi.data.MarkCell(uint(gi.cellNumber) - 1)
	gi.data.ChangePlayer()
}

func (gi *GameInput) checkInput() bool {
	if gi.input == "0" {
		fmt.Println("\ngood bye!")
		os.Exit(0)
	}

	var err error
	gi.cellNumber, err = strconv.Atoi(gi.input)
	if err != nil || gi.cellNumber < 1 || gi.cellNumber > 9 {
		fmt.Println("Enter a valid cell number!\n")
		return false
	}

	if gi.data.GetCells()[gi.cellNumber-1] != EMPTY_CELL {
		fmt.Println("This cell is already occupied!\n")
		return false
	}

	return true
}
