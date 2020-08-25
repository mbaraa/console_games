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

	var a2cTetrisMainMap [Const.ROWS][Const.COLUMNS]rune
	// mark each taken place with true
	var a2bCheckList [Const.ROWS][Const.COLUMNS]bool
	// columns lengths
	var anColsLengths [Const.COLUMNS]int
	// completed rows
	var abCompletedLines [Const.ROWS]bool
	// lines eliminated
	var nLines int = 0

	TF.InitLengths(&anColsLengths)
	TF.InitCheckList(&a2bCheckList)
	TF.InitTetrisMap(&a2cTetrisMainMap)
	TF.InitCompletedLines(&abCompletedLines)

	// start dropping from the first row
	var nCurrY int = -1
	// dropping starts from the middle
	var nCurrX int = 3

	var nNewX = int(nCurrX)

	// print current tetris
	Common.PrintMatrix(a2cTetrisMainMap)

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

	var auTetrominos [7]Types.Tetromino = [7]Types.Tetromino{square, I,
		L, LInverse,
		T, skew, skewInverse}

	// current tetromino index
	var nCurrTetromino int = 0

	// current tetromino object pointer
	var puBlock *Types.Tetromino

	// keystroke channels
	chcKeyChar := make(chan rune)
	//chuKeyCode := make(chan keyboard.Key) will be added in the near future

	// sending keystrokes to thier channels
	// this operation mus be done parallel to the main loop to avoid deadlocks
	// and to get keys w/o latincy
	go func() {
		for {
			cKeyChar, _, _ := keyboard.GetSingleKey()

			// append non-zero values to chaannel
			if cKeyChar > 0 {
				chcKeyChar <- cKeyChar
			}

		}
	}()

	///////////////////////////////////////////////////////////////////////////
	///////////////////////// game loop //////////////////////////////////////
	/////////////////////////////////////////////////////////////////////////

	for {
		puBlock = &auTetrominos[nCurrTetromino]
		puBlock.X = nCurrX
		puBlock.Y = nCurrY - (puBlock.Height - 1)

		// update statement 1:
		// overlapping checker1
		if nCurrY >= anColsLengths[nCurrTetromino] {
			//droppingRow = nColsLengths[col] // - 1
			TF.InitLengths(&anColsLengths)
			TF.CheckTetrisMap(a2cTetrisMainMap, &a2bCheckList, &anColsLengths)
		}

		// update statement 2:
		// restart loop if a shape hits the ground
		if nCurrY >= 0 && a2bCheckList[nCurrY][nCurrX] {

			if nCurrTetromino >= 6 {
				nCurrTetromino = 0
			} else {
				nCurrTetromino++
			}

			nCurrY = 0
			continue
		}

		/*go func() {

			chr, _, _ = keyboard.GetSingleKey()

		}()
		*/
		time.Sleep(time.Millisecond * 100)

		// move left / right, rotate, and quit(pause in future) controls
		// must be done parallel to the rest of the loop to avoid deadlocks
		go func() {
			chr := <-chcKeyChar

			if chr == 'A' || chr == 'a' {
				nNewX--
			} else if chr == 'D' || chr == 'd' {
				nNewX++
			} else if chr == 'W' || chr == 'w' {
				puBlock.Rotate90Degs()

			} else if chr == 'q' {
				os.Exit(0)
			}
		}()
		// boundaries
		if nNewX >= 9 {
			nNewX = 9
		} else if nNewX <= 0 {
			nNewX = 0
		}

		// update statement 3:
		// move block on the tetris map
		nCurrY++
		if puBlock.Y >= 0 &&
			nCurrY >= 0 &&
			nCurrY <= anColsLengths[nCurrTetromino] {

			TF.DropBlockOneRow(&a2cTetrisMainMap, puBlock,
				nCurrX, nCurrY, nNewX)
		}

		nCurrX = nNewX

		////////////
		/*if( tetrisMainMap[0][col] == '#' && tetrisMainMap[1][col] == '#') {
		    printGameOverAndGTFOH();
		}*/

		// update statements 4:
		// clear screen, update maps and print current tetris map
		Common.Clear()
		TF.MarkDoneLines(&a2bCheckList, &abCompletedLines, &nLines)
		TF.EliminateLines(&a2cTetrisMainMap, &a2bCheckList,
			&abCompletedLines, &anColsLengths)
		TF.UpdateTetrisMap(&a2cTetrisMainMap, &a2bCheckList)
		Common.PrintMatrix(a2cTetrisMainMap)

		//// exists for debugging

		fmt.Printf("droppingRow: %d, colLength: %d\n", nCurrX, anColsLengths[nCurrX])

		/*fmt.Println("lines completed:")
		for col := 0; col < Const.ROWS; col++ {

			fmt.Printf("%v ", bCompletedLines[col])

		}*/
		fmt.Println("\nCols lengths:")
		for col := 0; col < Const.COLUMNS; col++ {

			fmt.Printf("%d ", anColsLengths[col])

		}
		fmt.Printf("\nLines: %d", nLines)
		//// debugging budies

	} // game loop

}
