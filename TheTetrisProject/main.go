package main

import (
	"fmt"
	"time"
	//"os"
	Const "./Constants"
	Common "./Shared"
	TF "./TetrisFuncs"
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
	var col int = Const.COLUMNS/2 - 1

	// print current tetris
	Common.PrintMatrix(cTetrisMainMap)

	var newCol = int(col)

	/*var tetrominos [7]Tetromino = { square, straight,
	                       lShape, lShapeInverse,
	                       tShape, skew, skewInverse}

	 /* int counter = 0;
*/
	var block *Common.Tetromino
	//block = &tetrominos[counter];

	block = &skew
	block.X = 4
	block.Y = 0

	var chr rune
	//var key keyboard.Key

	// game loop
	for {

		block.Y = droppingRow - (block.Height - 1)

		// overlapping checker
		if droppingRow == nColsLengths[col] {
			fmt.Println(Const.RED)
			fmt.Println("what up bitch")
			fmt.Println(Const.RESET)
			TF.InitLengths(&nColsLengths)
			TF.CheckTetrisMap(&cTetrisMainMap, &bCheckList, &nColsLengths)
		}

		/*droppingRow++
		// replace dropped character with a dot
		if droppingRow < nColsLengths[newCol] && droppingRow > 0 {

			cTetrisMainMap[droppingRow-1][newCol] = '.'
		}*/

		if 	droppingRow >= 0 &&
			droppingRow <= nColsLengths[col] &&
			bCheckList[droppingRow][col] == true {
			/*if(counter >= 6) {
			      counter = 0;
			  	} else {
			      	counter++;
			  	}*/
			fmt.Println(Const.RED)
			fmt.Println("I Werk Bitch")
			fmt.Println(Const.RESET)
			//os.Exit(0)
			droppingRow = 0
			continue
		}

		go func() {
			chr, _, _ = keyboard.GetSingleKey()

		}()

		time.Sleep(time.Millisecond * 300)

		// move left / right
		if chr == 'A' || chr == 'a' {
			newCol--
		} else if chr == 'D' || chr == 'd' {
			newCol++
		}
		/*else if chr == 'W' || chr == 'w' {
		    if block->rotation >= 3 {
		        block->rotation = 0
		        block->rotationsCounter = 0;
		    }
		    else {
		        block->rotation++;
		    }
		}*/

		// setting boundaries
		if newCol >= 9 {
			newCol = 9
		} else if newCol <= 0 {
			newCol = 0
		}

		droppingRow++
		if 	block.Y >= 0 &&
			droppingRow >= 0 &&
			droppingRow <= nColsLengths[col] {

			TF.DropBlockOneRow(&cTetrisMainMap, block,
						   col, droppingRow, newCol)
		}

		col = newCol

		// overlapping checker
		/*if( droppingRow == nColsLengths[col]) {
		    initLengths(colsLengths);
		    checkTetrisMap(tetrisMainMap, checkList, colsLengths);
		}*/

		/*if( tetrisMainMap[0][col] == '#' && tetrisMainMap[1][col] == '#') {
		    printGameOverAndGTFOH();
		}*/

		/*if droppingRow < nColsLengths[newCol] {
			cTetrisMainMap[droppingRow][newCol] = '#'
		}*/

		//fmt.Printf("rune: %v key: %v \n", chr, key)

		Common.Clear()
		TF.MarkDoneLines(&bCheckList, &bCompletedLines, &nLines)
		TF.EliminateLines(&cTetrisMainMap, &bCheckList,
							&bCompletedLines, &nColsLengths)
		TF.UpdateTetrisMap(&cTetrisMainMap, &bCheckList)
		Common.PrintMatrix(cTetrisMainMap)
		fmt.Println(droppingRow, nColsLengths[col])

		//// exists for debugging
        fmt.Println("lines completed:")
        for col := 0; col < Const.ROWS; col++ {

            fmt.Printf("%v ", bCompletedLines[col])

        }
        fmt.Println("\nlines lengths:")
        for col := 0; col < Const.COLUMNS; col++ {

            fmt.Printf("%d ", nColsLengths[col])

        }
        fmt.Printf("\nLine: %d", nLines)
        //// debugging budies

	} // game loop

}

var (
	straight Common.Tetromino = Common.Tetromino {
		4,
		1,
		[4][2]rune{{'#', '\000'},
			{'#', '\000'},
			{'#', '\000'},
			{'#', '\000'}},

		[4][2]rune{{'.', '\000'},
			{'.', '\000'},
			{'.', '\000'},
			{'.', '\000'}},
		4,
		0,
		0,
		0}

	skew Common.Tetromino = Common.Tetromino {
		3,
		2,
		[4][2]rune{{'.', '#'},
			{'#', '#'},
			{'#', '.'},
			{'\000', '\000'}},
		[4][2]rune{{'.', '.'},
			{'.', '.'},
			{'.', '.'},
			{'\000', '\000'}},
		4,
		0,
		0,
		0}
/*
	lShape Common.Tetromino = Common.Tetromino {
    	3,
	    2,
    	[4][2]rune{ {'#', '.'},
      				{'#', '.'},
      				{'#', '#'},
      				{'\0', '\0'}},

    	[4][2]rune{ {'.', '.'},
      				{'.', '.'},
      				{'.', '.'},

		4,
    	0,
    	0,
    	0}


	lShapeInverse Common.Tetromino = Common.Tetromino {
	   	3,
	    2,
	   	[4][2]rune{ {'.', '#'},
	   				{'.', '#'},
	  				{'#', '#'},
	  				{'\0', '\0'}},

	   	[4][2]rune{ {'.', '.'},
	   				{'.', '.'},
	   				{'.', '.'},

		4,
	    0,
	    0,
	    0}

	tShape Common.Tetromino = Common.Tetromino {
	   	3,
	    2,
	   	[4][2]rune{ {'#', '.'},
	   				{'#', '#'},
	  				{'#', '.'},
	  				{'\0', '\0'}},

	   	[4][2]rune{ {'.', '.'},
	   				{'.', '.'},
	   				{'.', '.'},

		4,
	    0,
	    0,
	    0}

	skew Common.Tetromino = Common.Tetromino {
	   	3,
	    2,
	   	[4][2]rune{ {'.', '#'},
	   				{'#', '#'},
	  				{'#', '.'},
	  				{'\0', '\0'}},

	   	[4][2]rune{ {'.', '.'},
	   				{'.', '.'},
	   				{'.', '.'},

		4,
	    0,
	    0,
	    0}

	skew Common.Tetromino = Common.Tetromino {
		3,
		2,
		[4][2]rune{ {'#', '.'},
					{'#', '#'},
					{'.', '#'},
					{'\0', '\0'}},

		[4][2]rune{ {'.', '.'},
					{'.', '.'},
					{'.', '.'},

		4,
		0,
		0,
		0}
*/

)
