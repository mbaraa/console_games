package TetrisFuncs

import (
	"fmt"
	"os"

	Const "../Constants"
	Common "../Shared"
	"../Types"
)

// initializers:

// initialize lengths with row length - 1
func InitLengths(lengths *[Const.COLUMNS]int) {

	for i := 0; i < Const.COLUMNS; i++ {
		(*lengths)[i] = Const.ROWS - 1
	} // for

} // initLengths()

// initialize completed lines with falses(0)
func InitCompletedLines(completedLines *[Const.ROWS]bool) {

	for i := 0; i < Const.ROWS; i++ {
		(*completedLines)[i] = false
	} // for

} // initCompletedLines()

// initialize tetris boolean map with falses(0)
func InitCheckList(checkList *[Const.ROWS][Const.COLUMNS]bool) {

	for row := 0; row < Const.ROWS; row++ {

		for col := 0; col < Const.COLUMNS; col++ {
			(*checkList)[row][col] = false
		} // inner for
	} // outer for
} // initCheckList()

// initialize tetris map with dots
func InitTetrisMap(tetrisMap *[Const.ROWS][Const.COLUMNS]rune) {

	for row := 0; row < Const.ROWS; row++ {
		for col := 0; col < Const.COLUMNS; col++ {

			(*tetrisMap)[row][col] = '.'

		} // inner for
	} // outer for
} // initTetrisMap()

// replace true in boolean tetris map with #s in the tetris map
func UpdateTetrisMap(tetrisMap *[Const.ROWS][Const.COLUMNS]rune,
	checkList *[Const.ROWS][Const.COLUMNS]bool) {

	for row := 0; row < Const.ROWS; row++ {

		for col := 0; col < Const.COLUMNS; col++ {

			if (*checkList)[row][col] {
				(*tetrisMap)[row][col] = '#'
			} // if

		} // for
	} // for

} // updateTetrisMap()

// replace #s in tetris map with trues
func CheckTetrisMap(tetrisMap [Const.ROWS][Const.COLUMNS]rune,
	checkList *[Const.ROWS][Const.COLUMNS]bool,
	lengths *[Const.COLUMNS]int) {

	/* reversed checker to prevent stacked areas,
	 *   gone back to the original after having trouble with eliminateLines
	 */
	for col := 0; col < Const.COLUMNS; col++ {

		for row := 0; row < Const.ROWS; row++ {

			if tetrisMap[row][col] == '#' {

				(*checkList)[row][col] = true
				(*lengths)[col]--
				//break

			} // if
		} // inner for
	} // outer for

} // checkTetrisMap()

// if line is copleted mark its place in the completedLines array
func MarkDoneLines(tetrisBooleanMap *[Const.ROWS][Const.COLUMNS]bool,
	completedLines *[Const.ROWS]bool,
	eliminatedLines *int) {

	for row := 0; row < Const.ROWS; row++ {

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

			// I fucked up here not gonna lie
			(*completedLines)[row] = true

			// increase eliminated lines by one
			(*eliminatedLines)++

			/*for(int col = 0; col < COLUMNS; col++){

			    if( tetrisBooleanMap[row][col] != 1 ) {
			        break;
			    }
			    if( tetrisBooleanMap[row][9] == 1 ) {

			    }

			}*/

		} // crazy if
	} // outer for

} // void markDoneLines

// eliminate completed lines and shift upper lines down
func EliminateLines(tetrisMap *[Const.ROWS][Const.COLUMNS]rune,
	tetrisBooleanMap *[Const.ROWS][Const.COLUMNS]bool,
	completedLines *[Const.ROWS]bool,
	columnsLengths *[Const.COLUMNS]int) {

	// look for completed(filled) lines from down to up, it's more logical that way :)
	for row := Const.ROWS - 1; row >= 0; row-- {

		if (*completedLines)[row] {

			// set completion state to false
			(*completedLines)[row] = false

			// reset line
			for col := 0; col < Const.COLUMNS; col++ {
				(*tetrisMap)[row][col] = '.'
				(*tetrisBooleanMap)[row][col] = false
				// since there's an eleminated line so lengths are increased
				(*columnsLengths)[col]++

			} // inner for1

			// shift rows down
			// remeber we're going downnnn to up
			for col := 0; col < Const.COLUMNS; col++ {
				// rowEL is no spanish or mexican it's just (row at Eleminated Line)
				for rowEL := row; rowEL > 0; rowEL-- {

					(*tetrisMap)[rowEL][col] = (*tetrisMap)[rowEL-1][col]
					(*tetrisBooleanMap)[rowEL][col] = (*tetrisBooleanMap)[rowEL-1][col]

				} // inner inner for
			} // inner for 2

		} // if

	} // outer for

} // eliminateLines()

// quit the game
func PrintGameOverAndGTFOH() {
	Common.Clear()
	// print red game over
	fmt.Println(Const.RED)
	fmt.Println("GAME OVER !!!!")
	fmt.Println(Const.RESET)
	os.Exit(0)

} // printGameOverAndGTFOH()

func clearAboveLines(pa2cTetrisMap *[Const.ROWS][Const.COLUMNS]rune,
	uBlock Types.Tetromino) {

	//
	for y := 0; y < uBlock.Y+1; y++ {
		for x := uBlock.X; x < uBlock.Width+uBlock.X; x++ {
			//if uBlock.Y < Const.ROWS && uBlock.X+x < Const.COLUMNS { //&&
			if uBlock.Y >= 0 && uBlock.X >= 0 {

				(*pa2cTetrisMap)[y][x] = '.'
			}
		}
	}

}

func DropBlockOneRow(tetrisMap *[Const.ROWS][Const.COLUMNS]rune,
	block *Types.Tetromino,
	currX, currY, // usless shits, will be removed :)
	destX int) {

	x := &block.X
	y := &block.Y

	// draw equivalent empty block on the tetris map
	for shapeRow := 0; shapeRow < 4; shapeRow++ {

		for shapeCol := 0; shapeCol < 4; shapeCol++ {
			if *y+shapeRow < Const.ROWS && *x+shapeCol < Const.COLUMNS {
				//block.Shape[shapeRow][shapeCol] == '#' {

				(*tetrisMap)[*y+shapeRow][*x+shapeCol] = (*block).EqNone[shapeRow][shapeCol]
			}
		} //

	} //

	//clearAboveLines(tetrisMap, *block)
	// update coordinates
	*x = destX
	(*y)++ // drop one block

	// draw block on the tetris map
	for shapeRow := 0; shapeRow < 4; shapeRow++ {

		for shapeCol := 0; shapeCol < 4; shapeCol++ {
			if *y+shapeRow < Const.ROWS && *x+shapeCol < Const.COLUMNS { //&&
				//block.Shape[shapeRow][shapeCol] == '#' {
				(*tetrisMap)[*y+shapeRow][*x+shapeCol] = (*block).Shape[shapeRow][shapeCol]
			}
		} //

	} //

	/*	// draw equivalent empty block on the tetris map
		for shapeRow := 0; shapeRow < (*block).Height; shapeRow++ {

			for shapeCol := 0; shapeCol < (*block).Width; shapeCol++ {

				(*tetrisMap)[*y+shapeRow][*x+shapeCol] = (*block).EqNone[shapeRow][shapeCol]

			} //

		} //

		// update coordinates
		*x = destX
		(*y)++ // drop one block

		// draw block on the tetris map
		for shapeRow := 0; shapeRow < (*block).Height; shapeRow++ {

			for shapeCol := 0; shapeCol < (*block).Width; shapeCol++ {

				(*tetrisMap)[*y+shapeRow][*x+shapeCol] = (*block).Shape[shapeRow][shapeCol]

			} //

		} //
	*/
} // dropBlockOneRow()
