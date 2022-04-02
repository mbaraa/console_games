package gamestuff

import (
	"snek/snakes"
	"snek/tui"
)

func Start(level snakes.Level) {
	maxTermPoint, _ := tui.GetTermSize()
	maxTermPoint.Y -= 5
	plane := tui.NewXYPlane(maxTermPoint, "█", " ")

	io := NewGameIO(plane, level)
	go io.GetInput()
	io.PrintStuff()
}
