package main

import (
	"fmt"
)

func Clear() {

	fmt.Printf("\033[H\033[2J")

}

func PrintMatrix(matrix [ROWS][COLUMNS]rune) {

	for i := 0; i < ROWS; i++ {
		fmt.Print(i, "  ")
		if i < 10 {
			fmt.Print(" ")
		}
		for j := 0; j < COLUMNS; j++ {
			fmt.Printf("%c ", matrix[i][j])
		}
		fmt.Println("")
	}

}

func SwapInt(x, y *int) {

	var temp int = *x
	*x = *y
	*y = temp

}
