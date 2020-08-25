package main

import (
	"fmt"
	"os"
	"time"

	Const "./Constants"
	Common "./Shared"
	TF "./TetrisFuncs"
	"./Types"
	"github.com/eiannone/keyboard"
)

func main() {
	// start listening to keyboard
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	// end listening to keyboard when the program is done
	defer func() {
		_ = keyboard.Close()
	}()

	var cTetrisMainMap [Const.ROWS][Const.COLUMNS]rune
	// mark each taken place with true
	var bCheckList [Const.ROWS][Const.COLUMNS]bool
	// columns lengths
	var nColsLengths [Const.COLUMNS]int
	// completed rows
	var bCompletedLines [Const.ROWS]bool
	// lines eliminated
	var nLines int = 0

	TF.InitLengths(&nColsLengths)
	TF.InitCheckList(&bCheckList)
	TF.InitTetrisMap(&cTetrisMainMap)
	TF.InitCompletedLines(&bCompletedLines)

	// start dropping from the first row
	var droppingRow int = -1
	// dropping starts from the middle
	var col int = 3

	// print current tetris
	Common.PrintMatrix(cTetrisMainMap)

	var newCol = int(col)

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

	var counter int = 0

	var block *Types.Tetromino

	var chr rune

	//var key keyboard.Key

	go func() {
		for {
			chr, _, _ = keyboard.GetSingleKey()
		}
	}()

	///////////////////////////////////////////////////////////////////////////
	///////////////////////// game loop //////////////////////////////////////
	/////////////////////////////////////////////////////////////////////////

	for {
		block = &tetrominos[counter]
		block.X = col
		block.Y = droppingRow - (block.Height - 1)

		// update statement 1:
		// overlapping checker1
		if droppingRow >= nColsLengths[col] {
			//droppingRow = nColsLengths[col] // - 1
			TF.InitLengths(&nColsLengths)
			TF.CheckTetrisMap(cTetrisMainMap, &bCheckList, &nColsLengths)
		}

		// update statement 2:
		// restart loop if a shape hits the ground
		if droppingRow >= 0 && bCheckList[droppingRow][col] {

			if counter >= 6 {
				counter = 0
			} else {
				counter++
			}

			droppingRow = 0
			continue
		}

		/*go func() {

			chr, _, _ = keyboard.GetSingleKey()

		}()
		*/
		time.Sleep(time.Millisecond * 100)

		// move left / right, rotate, and quit(pause in future) controls
		if chr == 'A' || chr == 'a' {
			newCol--
		} else if chr == 'D' || chr == 'd' {
			newCol++
		} else if chr == 'W' || chr == 'w' {
			block.Rotate90Degs()

		} else if chr == 'q' {
			os.Exit(0)
		}

		// boundaries
		if newCol >= 9 {
			newCol = 9
		} else if newCol <= 0 {
			newCol = 0
		}

		// update statement 3:
		// move block on the tetris map
		droppingRow++
		if block.Y >= 0 &&
			droppingRow >= 0 &&
			droppingRow <= nColsLengths[col] {

			TF.DropBlockOneRow(&cTetrisMainMap, block,
				col, droppingRow, newCol)
		}

		col = newCol

		////////////
		/*if( tetrisMainMap[0][col] == '#' && tetrisMainMap[1][col] == '#') {
		    printGameOverAndGTFOH();
		}*/

		// update statements 4:
		// clear screen, update maps and print current tetris map
		Common.Clear()
		TF.MarkDoneLines(&bCheckList, &bCompletedLines, &nLines)
		TF.EliminateLines(&cTetrisMainMap, &bCheckList,
			&bCompletedLines, &nColsLengths)
		TF.UpdateTetrisMap(&cTetrisMainMap, &bCheckList)
		Common.PrintMatrix(cTetrisMainMap)

		//// exists for debugging

		fmt.Printf("droppingRow: %d, colLength: %d\n", droppingRow, nColsLengths[col])

		/*fmt.Println("lines completed:")
		for col := 0; col < Const.ROWS; col++ {

			fmt.Printf("%v ", bCompletedLines[col])

		}*/
		fmt.Println("\nCols lengths:")
		for col := 0; col < Const.COLUMNS; col++ {

			fmt.Printf("%d ", nColsLengths[col])

		}
		fmt.Printf("\nLines: %d", nLines)
		//// debugging budies

	} // game loop

}
