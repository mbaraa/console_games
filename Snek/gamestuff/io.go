package gamestuff

import (
	"fmt"
	"snek/snakes"
	"snek/tui"
	"time"

	kb "github.com/eiannone/keyboard"
)

type CellsFunc func() (old, dst tui.Cell)

type MoveDir uint

const (
	MoveUp MoveDir = 1 << iota
	MoveDown
	MoveRight
	MoveLeft
)

type GameIO struct {
	plane         *tui.XYPlane
	snake         *snakes.Snake
	cells         chan CellsFunc
	lastMove      MoveDir
	score, apples int
	speed         int
	key           kb.Key
}

func NewGameIO(plane *tui.XYPlane, level snakes.Level) *GameIO {
	mdPt := tui.Point2{
		X: plane.MaxPoint().X / 2,
		Y: plane.MaxPoint().Y / 2,
	}
	snake := snakes.NewSnake(mdPt)
	for i := 0; i < 8; i++ {
		snake.AddNode()
	}

	for n := snake.Nodes().Front(); ; n = n.Next() {
		plane.Mark(n.Value.(tui.Cell))
		if n.Next() == nil {
			break
		}
	}

	apples := snakes.GetApples(plane.MaxPoint(), level, tui.ColorBoldRed)
	for _, apple := range apples {
		plane.Mark(&apple)
	}

	return &GameIO{
		plane:    plane,
		snake:    snake,
		cells:    make(chan CellsFunc),
		score:    0,
		apples:   len(apples),
		speed:    getSpeed(level),
		lastMove: MoveRight,
		key:      kb.KeyArrowRight,
	}
}

func getSpeed(level snakes.Level) int {
	switch level {
	default:
		fallthrough
	case snakes.LevelNormal:
		return 8
	case snakes.LevelCrazy:
		return 16
	case snakes.LevelInsane:
		return 32
	case snakes.LevelSnek:
		return 64
	case snakes.LevelBerzerk:
		return 128
	}
}

func (g *GameIO) PrintStuff() {
	for {
		getCells, ok := <-g.cells
		if !ok {
			continue
		}

		o, d := getCells()

		if cell, marked := g.plane.Marked(d.Position()); marked && cell.Name() == "apple" {
			g.snake.AddNode()
			g.score++
			g.apples--
		}

		err := g.plane.Mark(d)
		if err != nil || g.apples == 0 {
			fmt.Printf("\033[H\033[2J%s\nScore: %d\n",
				tui.ColorBoldRed.StringColored("Game Over!"),
				g.score,
			)
			return
		}

		g.plane.UnMark(o)
		g.plane.Draw()
		fmt.Printf("Score: %d\tLeft Apples: %d\n", g.score, g.apples)
		time.Sleep(time.Second / time.Duration(g.speed))
	}
}

func (g *GameIO) GetInput() {
	var err error
	for {
		switch g.key {
		case kb.KeyArrowUp:
			if g.lastMove != MoveDown {
				g.lastMove = MoveUp
				go g.sendPts(MoveUp)
			}
		case kb.KeyArrowDown:
			if g.lastMove != MoveUp {
				g.lastMove = MoveDown
				go g.sendPts(MoveDown)
			}
		case kb.KeyArrowRight:
			if g.lastMove != MoveLeft {
				g.lastMove = MoveRight
				go g.sendPts(MoveRight)
			}
		case kb.KeyArrowLeft:
			if g.lastMove != MoveRight {
				g.lastMove = MoveLeft
				go g.sendPts(MoveLeft)
			}
		}
		_, g.key, err = kb.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if g.key == kb.KeyEsc {
			return
		}
	}
}

func (g *GameIO) sendPts(currMove MoveDir) {
	for {
		if g.lastMove != currMove {
			return
		}
		var old, dst tui.Cell
		switch currMove {
		case MoveUp:
			old, dst = g.snake.MoveY(1)
		case MoveDown:
			old, dst = g.snake.MoveY(-1)
		case MoveRight:
			old, dst = g.snake.MoveX(1)
		case MoveLeft:
			old, dst = g.snake.MoveX(-1)
		}
		g.cells <- func() (o, d tui.Cell) {
			return old, dst
		}
	}
}
