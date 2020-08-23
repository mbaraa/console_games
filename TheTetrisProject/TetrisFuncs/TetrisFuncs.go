package TetrisFuncs

import (
	"fmt"
	"os"

	Const "../Constants"
	Common "../Shared"
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

			if (*checkList)[row][col] == true {
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

		if (*tetrisBooleanMap)[row][0] == true &&
			(*tetrisBooleanMap)[row][1] == true &&
			(*tetrisBooleanMap)[row][2] == true &&
			(*tetrisBooleanMap)[row][3] == true &&
			(*tetrisBooleanMap)[row][4] == true &&
			(*tetrisBooleanMap)[row][5] == true &&
			(*tetrisBooleanMap)[row][6] == true &&
			(*tetrisBooleanMap)[row][7] == true &&
			(*tetrisBooleanMap)[row][8] == true &&
			(*tetrisBooleanMap)[row][9] == true {

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

		if (*completedLines)[row] == true {

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

/*transposedShape[2][3];
  char eqNoneTransposedShape[2][3];
*/

func DropBlockOneRow(tetrisMap *[Const.ROWS][Const.COLUMNS]rune,
	block *Common.Tetromino,
	currX, currY,
	destX int) {
	//int rotation) {

	// pi/2 rad rotation

	// swapping height and width is really reallly important trust me
	//swap(&block->height, &block->width);

	/*
	   for(; block->rotationsCounter <= block->rotation; block->rotationsCounter++) {


	   for(int row = 0; row < block->width; row++) {

	       for(int col = 0; col < block->height; col++) {

	           eqNoneTransposedShape[row][col] = block->eqNone[col][row];

	       }
	   }

	   for(int row = 0; row < block->width; row++) {

	       for(int col = 0; col < block->height; col++) {

	           transposedShape[row][col] = block->shape[col][row];

	       }
	   }

	   swap(&block->height, &block->width);

	   }
	   //int tmp = block->width;
	   //block->width = block->height;
	   //block->height = tmp;

	   int *x = &block->x;
	   int *y = &block->y;

	   // draw equivalent empty block on the tetris map
	   for(int shapeRow = 0; shapeRow < block -> height; shapeRow++) {

	       for(int shapeCol = 0; shapeCol < block -> width; shapeCol++) {

	           if(block->rotation > 0) {
	               tetrisMap[*y + shapeRow][*x + shapeCol] = eqNoneTransposedShape[shapeRow][shapeCol];
	           }

	           else {
	               tetrisMap[*y + shapeRow][*x + shapeCol] = block -> eqNone[shapeRow][shapeCol];
	           }
	       } //

	   } //

	   // update coordinates
	   *x = destX;
	   (*y)++; // drop one block

	   // draw block on the tetris map
	   for(int shapeRow = 0; shapeRow < block -> height; shapeRow++) {

	       for(int shapeCol = 0; shapeCol < block->width; shapeCol++) {

	           if(block->rotation > 0) {
	               tetrisMap[*y + shapeRow][*x + shapeCol] = transposedShape[shapeRow][shapeCol];
	           }
	           else {
	               tetrisMap[*y + shapeRow][*x + shapeCol] = block -> shape[shapeRow][shapeCol];
	           }
	       } //

	   } //
	*/

	// 0 rad rotation

	x := &block.X
	y := &block.Y

	// draw equivalent empty block on the tetris map
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
	// 0 rad rotation

} // dropBlockOneRow()
