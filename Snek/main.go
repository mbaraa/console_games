package main

import (
	"math/rand"
	"os"
	"time"

	kb "github.com/eiannone/keyboard"

	"snek/snakes"
	"snek/tui"
)

func main() {
	maxTermPoint, _ := tui.GetTermSize()
	plane := tui.NewXYPlane(maxTermPoint, "█", " ")

	key := make(chan kb.Key)
	go func() {
		for {
			_, _key, err := kb.GetSingleKey()
			if err != nil {
				panic(err)
			}
			if _key == kb.KeyEnd {
				os.Exit(0)
			}
			key <- _key
		}
	}()

	mxPt := plane.MaxPoint()
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixMicro())
		pt := tui.Point2{
			X: rand.Intn(mxPt.X),
			Y: rand.Intn(mxPt.Y),
		}
		plane.Mark(pt, tui.ColorBoldRed)
	}

	snek := snakes.NewSnake(tui.Point2{15, 3})
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

	for {
		old, dst := tui.Point2{}, tui.Point2{}

		switch <-key {
		case kb.KeyArrowUp:
			old, dst = snek.MoveY(1)
		case kb.KeyArrowDown:
			old, dst = snek.MoveY(-1)
		case kb.KeyArrowRight:
			old, dst = snek.MoveX(1)
		case kb.KeyArrowLeft:
			old, dst = snek.MoveX(-1)
		}

		if dst.X > plane.MaxPoint().X || dst.Y > plane.MaxPoint().Y ||
			dst.X < 0 || dst.Y < 0 {
			panic("blyat")
		}

		if plane.Marked(dst) {
			snek.AddNode()
		}
		plane.UnMark(old)
		plane.Mark(dst, tui.ColorBoldGreen)
		plane.Draw()
	}
}
