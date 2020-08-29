package main

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	// start listening to keyboard
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	// end listening to keyboard when the program is done
	defer keyboard.Close()

	// da tetris map
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
	var nCurrY int = 0
	// dropping starts from the middle
	var nCurrX int = 4

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

			// append non-zero values to channel
			if cKeyChar > 0 {
				chcKeyChar <- cKeyChar
			}

		}
	}()

	// rotation state, need to clean lines above the rotated tetromino
	var bIsRotUsed bool = false
	// game state boolean, well that's really obvious :)
	var bGameOn bool = true

	///////////////////////////////////////////////////////////////////////////
	///////////////////////// game loop //////////////////////////////////////
	/////////////////////////////////////////////////////////////////////////

	for bGameOn {
		puBlock = &auTetrominos[nCurrTetromino]
		puBlock.X = nCurrX
		puBlock.Y = nCurrY - (puBlock.Height - 1) // height-1 because of the 0 based arrays

		// update statement 1:
		// overlapping checker1
		if nCurrY >= anColsLengths[nCurrX] {

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

			puBlock.X = 4 // the middle of the tetris map
			nCurrX = 4

			nCurrY = 0
			continue
		}

		/*go func() {

			chr, _, _ = keyboard.GetSingleKey()

		}()
		*/
		// pause the game so it won't go crazy fast
		time.Sleep(time.Millisecond * 100)

		// move left / right, rotate, and quit(pause in future) controls
		// must be done parallel to the rest of the loop to avoid deadlocks
		go func() {
			// get keystroke charachter from its channel
			chr := <-chcKeyChar

			if chr == 'A' || chr == 'a' {
				nCurrX--
			} else if chr == 'D' || chr == 'd' {
				nCurrX++
			} else if chr == 'W' || chr == 'w' {
				bIsRotUsed = true
				puBlock.Rotate90Degs()

			} else if chr == 'q' {
				bGameOn = false
			}
		}()

		// boundaries
		if nCurrX >= 9-puBlock.Width+1 {
			nCurrX = 9 - puBlock.Width + 1
		} else if nCurrX <= 0 {
			nCurrX = 0
		}

		// update statement 3:
		// move block on the tetris map
		nCurrY++
		if puBlock.Y >= 0 &&
			nCurrY >= 0 &&
			nCurrY <= anColsLengths[nCurrX] {

			DropBlockOneRow(&a2cTetrisMainMap, puBlock, nCurrX)
		}

		// last checking statement:
		// game over!
		if IsGameOver(anColsLengths) {
			bGameOn = false
		}

		// update statements 4:
		// clear screen, update maps and print current tetris map
		Clear()
		MarkDoneLines(&a2bCheckList, &abCompletedLines, &nLines)
		EliminateLines(&a2cTetrisMainMap, &a2bCheckList,
			&abCompletedLines, &anColsLengths)
		UpdateTetrisMap(&a2cTetrisMainMap, &a2bCheckList)
		PrintMatrix(a2cTetrisMainMap)
		fmt.Printf("\n Lines: %d\n", nLines)

		// DEBUG:
		//printBoolMtrx(a2bCheckList)
		/*
			// DEBUG:
			fmt.Printf("droppingRow: %d, colLength: %d\n", nCurrY, anColsLengths[nCurrX])
			fmt.Printf("tetromino's y: %d, tetromino's x: %d\n", puBlock.Y, puBlock.X)
			fmt.Printf("tetromino's H: %d, tetromino's W: %d\n", puBlock.Height, puBlock.Width)
			/*fmt.Println("lines completed:")
			for col := 0; col < ROWS; col++ {

				fmt.Printf("%v ", abCompletedLines[col])

			}
			fmt.Println("\nCols lengths:")
			for col := 0; col < COLUMNS; col++ {

				fmt.Printf("%d ", anColsLengths[col])

			}
			//// debugging budies
		*/
		// DEBUG:
		// clear above lines from hashes
		if bIsRotUsed {
			for row := 0; row < nCurrY+puBlock.Height+1; row++ {

				for col := 0; col < COLUMNS; col++ {

					a2cTetrisMainMap[row][col] = '.'
				}

			}
			bIsRotUsed = false
		}

	} // game loop

	PrintGameOver()

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
