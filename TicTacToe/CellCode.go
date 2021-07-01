package main

import "strconv"

const (
	EMPTY_CELL = iota
	X
	O
)

// CellCode indicates what does a TicTacToe cell have
// and take a value from the constants above
type CellCode uint

func GetCellCode(cell CellCode, cellIndex int) string {
	switch cell {
	case X:
		return "X"
	case O:
		return "O"

	default:
		return strconv.Itoa(cellIndex + 1)
	}
}
