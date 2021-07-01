package main

type GameData struct {
	xScore        uint
	oScore        uint
	cells         []CellCode
	currentPlayer CellCode
}

var instance *GameData = nil

func NewGameData() *GameData {
	if instance == nil {
		instance = &GameData{
			xScore:        0,
			oScore:        0,
			cells:         make([]CellCode, 9),
			currentPlayer: X,
		}
	}
	return instance
}

func (gd *GameData) GetXScore() uint {
	return gd.xScore
}

func (gd *GameData) GetOScore() uint {
	return gd.oScore
}

func (gd *GameData) GetCells() []CellCode {
	return gd.cells
}

func (gd *GameData) GetCurrentPlayer() CellCode {
	return gd.currentPlayer
}

func (gd *GameData) ResetGame() {
	gd.resetScores()
	gd.ResetCells()
}

func (gd *GameData) resetScores() {
	gd.xScore = 0
	gd.oScore = 0
}

func (gd *GameData) ResetCells() {
	gd.cells = make([]CellCode, 9)
}

func (gd *GameData) MarkCell(cellNumber uint) {
	gd.markCellForPlayer(gd.currentPlayer, cellNumber)
}

func (gd *GameData) markCellForPlayer(cellCode CellCode, cellNumber uint) {
	if cellNumber <= 9 {
		gd.cells[cellNumber] = cellCode
	}
}

func (gd *GameData) ChangePlayer() {
	switch gd.currentPlayer {
	case X:
		gd.currentPlayer = O
		break
	case O:
		gd.currentPlayer = X
		break
	}
}

func (gd *GameData) IncrementXScore() {
	gd.incrementScore(X)
}

func (gd *GameData) IncrementOScore() {
	gd.incrementScore(O)
}

func (gd *GameData) incrementScore(code CellCode) {
	switch code {
	case X:
		gd.xScore++
		break
	case O:
		gd.oScore++
		break
	}
}

func (gd *GameData) AreAllCellsOccupied() bool {
	occupiedCells := 0
	for _, cell := range gd.cells {
		if cell != EMPTY_CELL {
			occupiedCells++
		}
	}

	return occupiedCells == 9
}
