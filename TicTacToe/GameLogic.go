package main

type GameLogic struct {
	data *GameData
}

func NewGameLogic(data *GameData) *GameLogic {
	return &GameLogic{data}
}

func (gl *GameLogic) CheckWinner() GameState {
	var state GameState
	if gl.checkIfXWon() {
		state = X_WON
	} else if gl.checkIfOWon() {
		state = O_WON
	} else {
		state = GAME_ON
	}
	gl.updateScore(state)
	return state
}

func (gl *GameLogic) checkIfOWon() bool {
	return gl.data.GetCells()[0] == gl.data.GetCells()[1] && gl.data.GetCells()[1] == gl.data.GetCells()[2] &&
		gl.data.GetCells()[0] == O || // first row
		gl.data.GetCells()[3] == gl.data.GetCells()[4] && gl.data.GetCells()[4] == gl.data.GetCells()[5] &&
			gl.data.GetCells()[3] == O || // second row
		gl.data.GetCells()[6] == gl.data.GetCells()[7] && gl.data.GetCells()[7] == gl.data.GetCells()[8] &&
			gl.data.GetCells()[6] == O || // third row
		// columns
		gl.data.GetCells()[0] == gl.data.GetCells()[3] && gl.data.GetCells()[3] == gl.data.GetCells()[6] &&
			gl.data.GetCells()[0] == O || // first column
		gl.data.GetCells()[1] == gl.data.GetCells()[4] && gl.data.GetCells()[4] == gl.data.GetCells()[7] &&
			gl.data.GetCells()[1] == O || // second column
		gl.data.GetCells()[2] == gl.data.GetCells()[5] && gl.data.GetCells()[5] == gl.data.GetCells()[8] &&
			gl.data.GetCells()[2] == O || // third column
		// diagonals
		gl.data.GetCells()[0] == gl.data.GetCells()[4] && gl.data.GetCells()[4] == gl.data.GetCells()[8] &&
			gl.data.GetCells()[0] == O || // first diagonal
		gl.data.GetCells()[2] == gl.data.GetCells()[4] && gl.data.GetCells()[4] == gl.data.GetCells()[6] &&
			gl.data.GetCells()[2] == O // second diagonal
}

func (gl *GameLogic) checkIfXWon() bool {
	return gl.data.GetCells()[0] == gl.data.GetCells()[1] && gl.data.GetCells()[1] == gl.data.GetCells()[2] &&
		gl.data.GetCells()[0] == X || // first row
		gl.data.GetCells()[3] == gl.data.GetCells()[4] && gl.data.GetCells()[4] == gl.data.GetCells()[5] &&
			gl.data.GetCells()[3] == X || // second row
		gl.data.GetCells()[6] == gl.data.GetCells()[7] && gl.data.GetCells()[7] == gl.data.GetCells()[8] &&
			gl.data.GetCells()[6] == X || // third row
		// columns
		gl.data.GetCells()[0] == gl.data.GetCells()[3] && gl.data.GetCells()[3] == gl.data.GetCells()[6] &&
			gl.data.GetCells()[0] == X || // first column
		gl.data.GetCells()[1] == gl.data.GetCells()[4] && gl.data.GetCells()[4] == gl.data.GetCells()[7] &&
			gl.data.GetCells()[1] == X || // second column
		gl.data.GetCells()[2] == gl.data.GetCells()[5] && gl.data.GetCells()[5] == gl.data.GetCells()[8] &&
			gl.data.GetCells()[2] == X || // third column
		// diagonals
		gl.data.GetCells()[0] == gl.data.GetCells()[4] && gl.data.GetCells()[4] == gl.data.GetCells()[8] &&
			gl.data.GetCells()[0] == X || // first diagonal
		gl.data.GetCells()[2] == gl.data.GetCells()[4] && gl.data.GetCells()[4] == gl.data.GetCells()[6] &&
			gl.data.GetCells()[2] == X // second diagonal
}

func (gl *GameLogic) updateScore(playerCode GameState) {
	switch playerCode {
	case X_WON:
		gl.data.IncrementXScore()
		break
	case O_WON:
		gl.data.IncrementOScore()
		break
	}
}
