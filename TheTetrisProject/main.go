package main

import (
	//"fmt"
	"time"
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
	var bCheckList[Const.ROWS][Const.COLUMNS] bool
	// columns lengths
	var nColsLengths[Const.COLUMNS] int
		// completed rows
	var bCompletedLines[Const.ROWS] bool
		// lines eliminated
	//var nLines int = 0

	TF.InitLengths(&nColsLengths)
	TF.InitCheckList(&bCheckList)
	TF.InitTetrisMap(&cTetrisMainMap)
	TF.InitCompletedLines(&bCompletedLines)

	// start dropping from the first row
	var droppingRow int = -1;
	// dropping starts from the middle
	var col = int(Const.COLUMNS / 2 - 1);

	// print current tetris
	Common.PrintMatrix(cTetrisMainMap);


	var newCol = int(col);


	/*Tetromino tetrominos[] = {square, straight,
	                       lShape, lShapeInverse,
	                       tShape, skew, skewInverse};

	  int counter = 0;

	  Tetromino *block;
	  //block = &tetrominos[counter];

	  block = &skew;
	  block->x = 4;
	  block->y = 0;
	*/


	var chr rune
	//var key keyboard.Key

	// game loop
	for {

		if( droppingRow == nColsLengths[col]) {
            TF.InitLengths(&nColsLengths)
            TF.CheckTetrisMap(&cTetrisMainMap, &bCheckList, &nColsLengths);
        }

		droppingRow++
		// replace dropped character with a dot
        if droppingRow < nColsLengths[newCol] && droppingRow > 0{

            cTetrisMainMap[droppingRow - 1][newCol] = '.'
        }

		if bCheckList[droppingRow][col] == true {
            /*if(counter >= 6) {
            counter = 0;
        	} else {
            	counter++;
        	}*/

            droppingRow = 0
            continue
        }

		go func() {
			chr, _, _ = keyboard.GetSingleKey()

		}()

		time.Sleep(time.Millisecond * 500)


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
        } else if newCol <= 0  {
            newCol = 0
        }

		// overlapping checker
        /*if( droppingRow == nColsLengths[col]) {
            initLengths(colsLengths);
            checkTetrisMap(tetrisMainMap, checkList, colsLengths);
        }*/

        /*if( tetrisMainMap[0][col] == '#' && tetrisMainMap[1][col] == '#') {
            printGameOverAndGTFOH();
        }*/





		if( droppingRow < nColsLengths[newCol] ) {
            cTetrisMainMap[droppingRow][newCol] = '#'
        }
		col = newCol
		//fmt.Printf("rune: %v key: %v \n", chr, key)


		Common.Clear()
		Common.PrintMatrix(cTetrisMainMap)

	} // game loop

}

var (

	straight Common.Tetromino = Common.Tetromino{
			4,
			1,
			[4][2]rune{ {'#', '\000'},
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



	skew Common.Tetromino = Common.Tetromino{
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

)
