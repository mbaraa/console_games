package main

// tetromino struct
type Tetromino struct {
	/*	CenterX int // center of the shape

		CenterY int // center of the shape
	*/Height int // space taken on Y axis

	Width int // space taken on X axis

	Shape [4][4]rune // shape matrix

	EqNone [4][4]rune // equivalent none shape matrix

	X int // X axis

	Y int // Y axis

	Rotation int // rotation * pi/2 rad

	RotationsCounter int //
}

// private method to get equivalent none shape
func (tet Tetromino) getEqNone() [4][4]rune {
	return tet.EqNone
}

// private method to put the rotated tetromino back where it was
func (tet *Tetromino) shiftToOriginalPlace() {

	var (
		anXDist []int
		anYDist []int
	)

	for row := 0; row < 4; row++ {

		for col := 0; col < 4; col++ {

			if (*tet).Shape[row][col] == '#' {

				anXDist = append(anXDist, col)
				anYDist = append(anYDist, row)

			}

		}
	}

	var (
		nStartingRow int = GetMinArrayElement(anYDist)
		nStartingCol int = GetMinArrayElement(anXDist)
	)

	var a2cEmptyShape [4][4]rune = (*tet).EqNone

	var (
		newRow int
		newCol int
	)

	for row := nStartingRow; row < (*tet).Height+nStartingRow; row++ {

		for col := nStartingCol; col < (*tet).Width+nStartingCol; col++ {

			a2cEmptyShape[newRow][newCol] = (*tet).Shape[row][col]
			newCol++
		}

		newCol = 0
		newRow++

	}

	(*tet).Shape = a2cEmptyShape

}

// Rotate the shape Π/2 radians (euler's formula)
func (tet *Tetromino) Rotate90Degs() {

	// given a point p(x,y), a center of rotation c(a,b)
	//                  Z                           w
	/*
			 * using Euler's rotation formula
		     * Z' = -i(Z - w) + w
			 * we get the rotated point p' about c
	*/

	// ok let's start

	// declare complex numbers needed for rotation
	var Zprime, i, w complex128
	i = 0 + 1i
	// rotate about center of the shape c(1,2)
	w = 1 + 2i

	var newShape [4][4]rune

	newShape = tet.getEqNone()

	// store the rotated shape in the new matrix
	for row := 0.0; row < 4; row++ {

		for col := 0.0; col < 4; col++ {

			Z := complex(col, row)
			Zprime = i*(Z-w) + w

			x := int(real(Zprime))
			y := int(imag(Zprime)) - 1

			if x >= 0 && y >= 0 && tet.Shape[y][x] == '#' {
				newShape[int(row)][int(col)] = tet.Shape[y][x]

			}

		}

	}

	// last but not least swapping hieght & width
	SwapInt(&tet.Height, &tet.Width)

	// set the rotated shape
	tet.Shape = newShape

	tet.shiftToOriginalPlace()

}

// create a square tetromino
func (_ Tetromino) CreateSquareShape() Tetromino {

	return Tetromino{
		2,
		2,
		[4][4]rune{{'#', '#', '.', '.'},
			{'#', '#', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'}},

		[4][4]rune{{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'}},
		4,
		0,
		0,
		0}

}

// create an I shaped tetromino
func (_ Tetromino) CreateIShape() Tetromino {

	return Tetromino{
		4,
		1,
		[4][4]rune{{'#', '.', '.', '.'},
			{'#', '.', '.', '.'},
			{'#', '.', '.', '.'},
			{'#', '.', '.', '.'}},

		[4][4]rune{{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'}},

		4,
		0,
		0,
		0}

}

// create an skew shaped tetromino
func (_ Tetromino) CreateSkewShape() Tetromino {

	return Tetromino{
		3,
		2,
		[4][4]rune{{'.', '#', '.', '.'},
			{'#', '#', '.', '.'},
			{'#', '.', '.', '.'},
			{'.', '.', '.', '.'}},

		[4][4]rune{{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'}},
		4,
		0,
		0,
		0}

}

// create an skewInverse shaped tetromino
func (_ Tetromino) CreateSkewInverseShape() Tetromino {

	return Tetromino{
		3,
		2,
		[4][4]rune{{'#', '.', '.', '.'},
			{'#', '#', '.', '.'},
			{'.', '#', '.', '.'},
			{'.', '.', '.', '.'}},

		[4][4]rune{{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'}},
		4,
		0,
		0,
		0}

}

// create an L shaped tetromino
func (_ Tetromino) CreateLShape() Tetromino {

	return Tetromino{
		3,
		2,
		[4][4]rune{{'#', '.', '.', '.'},
			{'#', '.', '.', '.'},
			{'#', '#', '.', '.'},
			{'.', '.', '.', '.'}},

		[4][4]rune{{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'}},
		4,
		0,
		0,
		0}

}

// create an L inversed shaped tetromino
func (_ Tetromino) CreateLInverseShape() Tetromino {

	return Tetromino{
		3,
		2,
		[4][4]rune{{'.', '#', '.', '.'},
			{'.', '#', '.', '.'},
			{'#', '#', '.', '.'},
			{'.', '.', '.', '.'}},

		[4][4]rune{{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'}},
		4,
		0,
		0,
		0}

}

// create a T shaped tetromino
func (_ Tetromino) CreateTShape() Tetromino {

	return Tetromino{
		2,
		3,
		[4][4]rune{{'.', '#', '.', '.'},
			{'#', '#', '#', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'}},

		[4][4]rune{{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'},
			{'.', '.', '.', '.'}},
		4,
		0,
		0,
		0}

}
