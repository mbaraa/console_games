package main

import (
	"fmt"

	"./Types"
)

func main() {

	var (
		temp        Types.Tetromino
		square      = temp.CreateSquareShape()
		I           = temp.CreateIShape()
		skew        = temp.CreateSkewShape()
		skewInverse = temp.CreateSkewInverseShape()
		L           = temp.CreateLShape()
		LInverse    = temp.CreateLInverseShape()
		T           = temp.CreateTShape()
	)

	var tetrominos [7]Types.Tetromino = [7]Types.Tetromino{square, I,
		L, LInverse,
		T, skew, skewInverse}

	var tet Types.Tetromino
	tet = tetrominos[3]

	fmt.Println("Original:")
	PrintMatrix(tet.Shape)

	fmt.Println("1st rotate:")
	tet.Rotate90Degs()
	PrintMatrix(tet.Shape)

	fmt.Println("2nd rotate:")
	tet.Rotate90Degs()
	PrintMatrix(tet.Shape)

	fmt.Println("3rd rotate:")
	tet.Rotate90Degs()
	PrintMatrix(tet.Shape)

}

func PrintMatrix(matrix [4][4]rune) {

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%c ", matrix[i][j])
		}
		fmt.Println("")
	}

}
