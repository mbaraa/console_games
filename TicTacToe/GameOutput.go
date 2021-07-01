package main

import "fmt"

type GameOutput struct {
	data *GameData
}

const (
	GREENBOLD_OUTPUT = "\033[1m\033[32m"
	RESET_COLOR      = "\033[0m"
)

func NewGameOutput(data *GameData) *GameOutput {
	return &GameOutput{data}
}

func (gop *GameOutput) PrintGame() {
	fmt.Print("\033[H\033[2J") // clear screen
	gop.printScores()
	gop.printCells()
}

func (gop *GameOutput) printScores() {
	fmt.Printf("Scores:\nX: %d\tO: %d\n",
		gop.data.GetXScore(), gop.data.GetOScore())
}

func (gop *GameOutput) printCells() {
	for i := 0; i < len(gop.data.GetCells())-2; i += 3 {
		fmt.Println("_________________________________________________")
		fmt.Println("|\t \t|\t \t|\t \t|")
		fmt.Printf("|\t%s%s%s\t|\t%s%s%s\t|\t%s%s%s\t|\n",
			GREENBOLD_OUTPUT,
			GetCellCode(gop.data.GetCells()[i], i),
			RESET_COLOR,
			GREENBOLD_OUTPUT,
			GetCellCode(gop.data.GetCells()[i+1], i+1),
			RESET_COLOR,
			GREENBOLD_OUTPUT,
			GetCellCode(gop.data.GetCells()[i+2], i+2),
			RESET_COLOR,
		)
		fmt.Println("|\t \t|\t \t|\t \t|")
	}

	fmt.Println("_________________________________________________")
}
