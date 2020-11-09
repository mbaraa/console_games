package main

import (
	"fmt"
)

// initializers:

// initialize lengths with row length - 1
func InitLengths(lengths *[COLUMNS]int) {
	for i := 0; i < COLUMNS; i++ {
		(*lengths)[i] = ROWS - 1
	}

}

// initialize completed lines with falses(0)
func InitCompletedLines(completedLines *[ROWS]bool) {
	for i := 0; i < ROWS; i++ {
		(*completedLines)[i] = false
	}

}

// initialize tetris boolean map with falses(0)
func InitCheckList(checkList *[ROWS][COLUMNS]bool) {
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLUMNS; col++ {
			(*checkList)[row][col] = false
		}
	}
}

// initialize tetris map with dots
func InitTetrisMap(tetrisMap *[ROWS][COLUMNS]rune) {
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLUMNS; col++ {
			(*tetrisMap)[row][col] = '.'
		}
	}
}

// the other folks

// replace true in boolean tetris map with #s in the tetris map
func UpdateTetrisMap(tetrisMap *[ROWS][COLUMNS]rune,
	checkList *[ROWS][COLUMNS]bool) {

	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLUMNS; col++ {
			if (*checkList)[row][col] {
				(*tetrisMap)[row][col] = '#'
			}

		}
	}

}

// replace #s in tetris map with trues
func CheckTetrisMap(tetrisMap [ROWS][COLUMNS]rune,
	checkList *[ROWS][COLUMNS]bool,
	lengths *[COLUMNS]int) {

	// reversed checker to prevent stacked areas,
	// gone back to the original after having trouble with eliminateLines

	// this one is for updating tetris boolean map
	for col := 0; col < COLUMNS; col++ {
		for row := 0; row < ROWS; row++ {
			if tetrisMap[row][col] == '#' {
				(*checkList)[row][col] = true
			}
		}
	}
	// this one is for updating columns lengths
	for col := 0; col < COLUMNS; col++ {
		for row := 0; row < ROWS; row++ {
			if tetrisMap[row][col] == '#' {
				(*lengths)[col] = row - 1
				break
			}
		}
	}
}

// if line is completed mark its place in the completedLines array
func MarkDoneLines(tetrisBooleanMap *[ROWS][COLUMNS]bool,
	completedLines *[ROWS]bool,
	eliminatedLines *int) {

	for row := 0; row < ROWS; row++ {

		// if the first field in the line is filled check the next ones
		if (*tetrisBooleanMap)[row][0] &&
			(*tetrisBooleanMap)[row][1] &&
			(*tetrisBooleanMap)[row][2] &&
			(*tetrisBooleanMap)[row][3] &&
			(*tetrisBooleanMap)[row][4] &&
			(*tetrisBooleanMap)[row][5] &&
			(*tetrisBooleanMap)[row][6] &&
			(*tetrisBooleanMap)[row][7] &&
			(*tetrisBooleanMap)[row][8] &&
			(*tetrisBooleanMap)[row][9] {

			// I fucked up here not gonna lie ;)
			(*completedLines)[row] = true

			// increase eliminated lines by one
			(*eliminatedLines)++

		} // if
	} // for
}

// eliminate completed lines and shift upper lines down
func EliminateLines(tetrisMap *[ROWS][COLUMNS]rune,
	tetrisBooleanMap *[ROWS][COLUMNS]bool,
	completedLines *[ROWS]bool,
	columnsLengths *[COLUMNS]int,
	gameSpeed *float64) {

	// look for completed(filled)
	for row := 0; row < ROWS; row++ {
		if (*completedLines)[row] {
			// increase game speed
			(*gameSpeed) -= 0.1
			// set completion state to false
			(*completedLines)[row] = false
			// reset line
			for col := 0; col < COLUMNS; col++ {
				(*tetrisMap)[row][col] = '.'
				(*tetrisBooleanMap)[row][col] = false
				// since there's an eleminated line so lengths are increased
				(*columnsLengths)[col]++

			}

			// shift rows down
			// remeber we're going downnnn to up
			for col := 0; col < COLUMNS; col++ {
				// rowEL is no spanish or mexican it's just (row at Eleminated Line)
				for rowEL := row; rowEL > 0; rowEL-- {

					(*tetrisMap)[rowEL][col] = (*tetrisMap)[rowEL-1][col]
					(*tetrisBooleanMap)[rowEL][col] = (*tetrisBooleanMap)[rowEL-1][col]

				} // inner for 2
			} // inner for1
		} // if
	} // outer for

}

//
func PrintGameOver() {
	Clear()
	// print red game over
	fmt.Println(RED)
	fmt.Println("GAME OVER !!!!")
	fmt.Println(RESET)

}

func clearAboveLines(pa2cTetrisMap *[ROWS][COLUMNS]rune,
	uBlock Tetromino) {

	x := uBlock.X
	y := uBlock.Y

	// draw equivalent empty block on the tetris map
	for shapeRow := 0; shapeRow < y+uBlock.Height+2; shapeRow++ {

		for shapeCol := 0; shapeCol < COLUMNS; shapeCol++ {
			if y+shapeRow < ROWS && x+shapeCol < COLUMNS {

				(*pa2cTetrisMap)[y+shapeRow][x+shapeCol] = '.'
			}
		} //

	} //

}

func DropBlockOneRow(tetrisMap *[ROWS][COLUMNS]rune,
	block *Tetromino,
	destX int) {

	x := &block.X
	y := &block.Y

	clearAboveLines(tetrisMap, *block)
	// update coordinates
	*x = destX
	(*y)++ // drop one block

	// draw block on the tetris map
	for shapeRow := 0; shapeRow < 4; shapeRow++ {

		for shapeCol := 0; shapeCol < 4; shapeCol++ {
			if *y+shapeRow < ROWS && *x+shapeCol < COLUMNS &&
				block.Shape[shapeRow][shapeCol] == '#' {
				(*tetrisMap)[*y+shapeRow][*x+shapeCol] = (*block).Shape[shapeRow][shapeCol]
			} // if
		} // inner for
	} // outer for
}

func IsGameOver(anColsLengths [COLUMNS]int) bool {
	for col := 0; col < COLUMNS; col++ {
		if anColsLengths[col] == 0 {
			return true
		}
	}

	return false
}
