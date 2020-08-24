package SharedFuncs

import (
	"fmt"

	Const "../Constants"
)

func Clear() {

	fmt.Printf("\033[H\033[2J")

}

func PrintMatrix(matrix [Const.ROWS][Const.COLUMNS]rune) {

	for i := 0; i < Const.ROWS; i++ {
		for j := 0; j < Const.COLUMNS; j++ {
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
