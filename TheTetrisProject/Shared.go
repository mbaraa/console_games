package main

import (
	"fmt"
)

func Clear() {

	fmt.Printf("\033[H\033[2J")

}

// added additional stuff to make the map look better
func PrintMatrix(matrix [ROWS][COLUMNS]rune) {
	fmt.Print("\n ")
	for i := 0; i < ROWS+3; i++ {
		fmt.Print("7")
	}
	fmt.Println("")
	for i := 0; i < ROWS; i++ {
		fmt.Print(" 7 ")

		for j := 0; j < COLUMNS; j++ {

			fmt.Printf("%c ", matrix[i][j])
		}
		fmt.Println("7")
	}
	fmt.Print(" ")

	for i := 0; i < ROWS+3; i++ {
		fmt.Print("7")
	}
	fmt.Println("")

}
func SwapInt(x, y *int) {

	var temp int = *x
	*x = *y
	*y = temp

}

func GetMinArrayElement(array []int) int {
	var min = array[0]

	for i := 0; i < len(array); i++ {

		if array[i] < min {
			min = array[i]
		}

	}

	return min

}
