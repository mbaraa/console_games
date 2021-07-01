package main

import "fmt"

func main() {
	data := NewGameData()
	input := NewGameInput(data)
	logic := NewGameLogic(data)
	output := NewGameOutput(data)

	for {
		output.PrintGame()
		input.GetInput()

		if state := logic.CheckWinner(); state != GAME_ON {
			output.PrintGame()
			switch state {
			case X_WON:
				fmt.Println("X Won!")
				break
			case O_WON:
				fmt.Println("O Won!")
				break
			}

			fmt.Print("press enter to continue...")
			fmt.Scanln()
			data.ResetCells()
			continue
		}
	}
}
