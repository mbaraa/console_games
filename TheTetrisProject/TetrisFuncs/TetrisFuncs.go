package TetrisFuncs

import Const "../Constants"

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


// replace #s in tetris map with trues
func CheckTetrisMap( tetrisMap *[Const.ROWS][Const.COLUMNS]rune, checkList *[Const.ROWS][Const.COLUMNS]bool, lengths *[Const.COLUMNS]int ) {

    /* reversed checker to prevent stacked areas,
     *   gone back to the original after having trouble with eliminateLines
     */
   for col := 0; col < Const.COLUMNS; col++ {

        for row := 0; row < Const.ROWS; row++ {

            if (*tetrisMap)[row][col] == '#' {

                (*checkList)[row][col] = true;
                (*lengths)[col]--;
                //break;

            } // if
        } // inner for
    } // outer for

} // void checkTetrisMap
