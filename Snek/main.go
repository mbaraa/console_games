package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	kb "github.com/eiannone/keyboard"

	"snek/snakes"
	"snek/tui"
)

type ptsFunc func() (old, dst tui.Point2)

type dir uint

const (
	up dir = iota
	down
	left
	right
)

var (
	_          = func() bool { rand.Seed(time.Now().UnixMicro()); return true }()
	key        = kb.Key(rand.Intn(0xFFEE-0xFFEA) + 0xFFEA)
	stop       = false
	score      = 0
	last   dir = dir(rand.Intn(4))
	apples     = rand.Intn(300-25) + 25
)

func main() {
	maxTermPoint, _ := tui.GetTermSize()
	maxTermPoint.Y -= 5
	maxTermPoint.X -= 2
	plane := tui.NewXYPlane(maxTermPoint, "█", ".")

	mxPt := plane.MaxPoint()
	for i := 0; i < apples; i++ {
		rand.Seed(time.Now().UnixMicro())
		pt := tui.Point2{
			X: rand.Intn(mxPt.X),
			Y: rand.Intn(mxPt.Y),
		}
		plane.Mark(pt, tui.ColorBoldRed)
	}

	snek := snakes.NewSnake(tui.Point2{X: rand.Intn(mxPt.X), Y: rand.Intn(mxPt.Y)})
	for i := 0; i < 5; i++ {
		snek.AddNode()
	}
	nds := snek.Nodes()
	for p := nds.Front(); ; p = p.Next() {
		plane.Mark(p.Value.(tui.Point2), tui.ColorBoldGreen)
		if p == nds.Back() {
			break
		}
	}

	pts := make(chan ptsFunc)
	go func() {
		for {
			var err error
			switch key {
			case kb.KeyArrowUp:
				last = up
				go send(up, snek, pts)
			case kb.KeyArrowDown:
				last = down
				go send(down, snek, pts)
			case kb.KeyArrowRight:
				last = right
				go send(right, snek, pts)
			case kb.KeyArrowLeft:
				last = left
				go send(left, snek, pts)
			}
			_, key, err = kb.GetSingleKey()
			if err != nil {
				panic(err)
			}
			if key == kb.KeyEnd {
				os.Exit(0)
			}
		}
	}()
	receive(plane, snek, pts)
}

func send(movDir dir, snek *snakes.Snake, ch chan ptsFunc) {
	for {
		if last != movDir {
			return
		}
		var o, d tui.Point2
		switch movDir {
		case up:
			o, d = snek.MoveY(1)
		case down:
			o, d = snek.MoveY(-1)
		case right:
			o, d = snek.MoveX(1)
		case left:
			o, d = snek.MoveX(-1)
		}
		ch <- func() (old tui.Point2, dst tui.Point2) {
			return o, d
		}
	}
}

func receive(plane *tui.XYPlane, snek *snakes.Snake, ch chan ptsFunc) {
	for {
		pp, ok := <-ch
		if !ok { //|| stop {
			continue
		}
		o, d := (pp)()

		if color, marked := plane.Marked(d); marked && color != tui.ColorBoldGreen {
			snek.AddNode()
			score++
		}

		err := plane.Mark(d, tui.ColorBoldGreen)
		if err != nil || score == apples {
			fmt.Printf("\033[H\033[2J%s\nScore: %d\n",
				tui.ColorBoldRed.StringColored("Game Over!"),
				score,
			)
			return
		}

		plane.UnMark(o)
		plane.Draw()
		fmt.Printf("Score: %d\t Left Apples: %d\n", score, apples-score)
		time.Sleep(time.Second / 5)
	}
}
