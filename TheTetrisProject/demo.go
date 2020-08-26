package main

import (
	"fmt"
	"time"
)

func main() {

	var (
		temp        Tetromino
		square      = temp.CreateSquareShape()
		I           = temp.CreateIShape()
		skew        = temp.CreateSkewShape()
		skewInverse = temp.CreateSkewInverseShape()
		L           = temp.CreateLShape()
		LInverse    = temp.CreateLInverseShape()
		T           = temp.CreateTShape()
	)

	var tetrominos [7]Tetromino = [7]Tetromino{square, I,
		L, LInverse,
		T, skew, skewInverse}

	var tet Tetromino
	i := 0
	for {
		if i >= 7 {
			i = 0
		}

		tet = tetrominos[i]

		fmt.Println("Original:")
		PrintMatrixIn(tet.Shape)

		fmt.Println("1st rotate:")
		tet.Rotate90Degs()
		PrintMatrixIn(tet.Shape)

		fmt.Println("2nd rotate:")
		tet.Rotate90Degs()
		PrintMatrixIn(tet.Shape)

		fmt.Println("3rd rotate:")
		tet.Rotate90Degs()
		PrintMatrixIn(tet.Shape)

		time.Sleep(time.Second * 2)
		i++
	}

}

func PrintMatrixIn(matrix [4][4]rune) {

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%c ", matrix[i][j])
		}
		fmt.Println("")
	}

}
