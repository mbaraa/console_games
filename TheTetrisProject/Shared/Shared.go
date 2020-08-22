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

type Tetromino struct {
	Height int // space taken on Y axis

	Width int // space taken on X axis

	Shape [4][2]rune // 4x2

	EqNone [4][2]rune // 4x2

	X int // X axis

	Y int // Y axis

	Rotation int // rotation * pi/2 rad

	RotationsCounter int //
}
