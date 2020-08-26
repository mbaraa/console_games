package main

import (
	"fmt"
	"os"
	"time"

	/*Const "./Constants"
	Common "./Shared"
	TF "./TetrisFuncs"
	"./Types"*/
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

	var a2cTetrisMainMap [ROWS][COLUMNS]rune
	// mark each taken place with true
	var a2bCheckList [ROWS][COLUMNS]bool
	// columns lengths
	var anColsLengths [COLUMNS]int
	// completed rows
	var abCompletedLines [ROWS]bool
	// lines eliminated
	var nLines int = 0

	InitLengths(&anColsLengths)
	InitCheckList(&a2bCheckList)
	InitTetrisMap(&a2cTetrisMainMap)
	InitCompletedLines(&abCompletedLines)

	// start dropping from the first row
	var nCurrY int = -1
	// dropping starts from the middle
	var nCurrX int = 3

	var nNewX = int(nCurrX)

	// print current tetris
	PrintMatrix(a2cTetrisMainMap)

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

	var auTetrominos [7]Tetromino = [7]Tetromino{square, I,
		L, LInverse,
		T, skew, skewInverse}

	// current tetromino index
	var nCurrTetromino int = 0

	// current tetromino object pointer
	var puBlock *Tetromino

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
		puBlock.Y = nCurrY - (puBlock.Height - 1) // height-1 because of the 0 based arrays

		// update statement 1:
		// overlapping checker1
		if nCurrY >= anColsLengths[nCurrX] {
			// DEBUG:
			//droppingRow = nColsLengths[col] // - 1
			// normal
			InitLengths(&anColsLengths)
			CheckTetrisMap(a2cTetrisMainMap, &a2bCheckList, &anColsLengths)
		}

		// update statement 2:
		// restart loop to avoid overlapping if a shape hits the ground
		if nCurrY >= 0 && a2bCheckList[nCurrY][nCurrX] {

			// use the next tetromino
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
		time.Sleep(time.Millisecond * 150)

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
			nCurrY <= anColsLengths[nCurrX] {

			DropBlockOneRow(&a2cTetrisMainMap, puBlock,
				nCurrX, nCurrY, nNewX)
		}

		nCurrX = nNewX

		////////////
		/*if( tetrisMainMap[0][col] == '#' && tetrisMainMap[1][col] == '#') {
		    printGameOverAndGTFOH();
		}*/

		// update statements 4:
		// clear screen, update maps and print current tetris map
		Clear()
		MarkDoneLines(&a2bCheckList, &abCompletedLines, &nLines)
		EliminateLines(&a2cTetrisMainMap, &a2bCheckList,
			&abCompletedLines, &anColsLengths)
		UpdateTetrisMap(&a2cTetrisMainMap, &a2bCheckList)
		PrintMatrix(a2cTetrisMainMap)
		// DEBUG:
		printBoolMtrx(a2bCheckList)

		// DEBUG:
		fmt.Printf("droppingRow: %d, colLength: %d\n", nCurrY, anColsLengths[nCurrX])
		fmt.Printf("tetromino's y: %d, tetromino's x: %d\n", puBlock.Y, puBlock.X)
		/*fmt.Println("lines completed:")
		for col := 0; col < Const.ROWS; col++ {

			fmt.Printf("%v ", bCompletedLines[col])

		}*/
		fmt.Println("\nCols lengths:")
		for col := 0; col < COLUMNS; col++ {

			fmt.Printf("%d ", anColsLengths[col])

		}
		fmt.Printf("\nLines: %d\n", nLines)
		//// debugging budies

	} // game loop

}

// DEBUG:
func printBoolMtrx(boolmtrx [ROWS][COLUMNS]bool) {
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			fmt.Printf("%v ", boolmtrx[i][j])
		}
		fmt.Println("")
	}
}
