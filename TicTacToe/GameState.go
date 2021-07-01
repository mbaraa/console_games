package main

const (
	GAME_ON = iota
	X_WON
	O_WON
	TIE
	GAME_OVER
)

// GameState indicates the status of the game
// and takes a value of the above constants
type GameState uint
