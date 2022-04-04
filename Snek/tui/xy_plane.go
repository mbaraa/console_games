package tui

import (
	"fmt"
	"strings"
)

type Point2 struct {
	X, Y int
}

type XYPlane struct {
	filled, empty string
	maxPoint      Point2
	cells         []Cell
}

func NewXYPlane(maxPoint Point2, filled, empty string) *XYPlane {
	cells := make([]Cell, maxPoint.X*maxPoint.Y)
	for y := 0; y < maxPoint.Y; y++ {
		for x := 0; x < maxPoint.X; x++ {
			cells[y*maxPoint.X+x] = NewColoredCell("empty", empty, ColorReset, Point2{X: x, Y: y})
		}
	}

	return &XYPlane{
		filled:   filled,
		empty:    empty,
		maxPoint: maxPoint,
		cells:    cells,
	}
}

func (p *XYPlane) Draw() {
	fmt.Println("\033[H\033[2J")
	fmt.Println(p.String())
}

func (p *XYPlane) String() string {
	sb := new(strings.Builder)

	for row := p.maxPoint.Y - 1; row >= 0; row-- {
		for col := 0; col < p.maxPoint.X; col++ {
			currPoint := p.cells[row*p.maxPoint.X+col]
			if currPoint != nil {
				sb.WriteString(p.cells[row*p.maxPoint.X+col].String())
			} else {
				sb.WriteString(p.empty)
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (p *XYPlane) Marked(_p Point2) (Cell, bool) {
	if _p.X >= p.maxPoint.X || _p.Y >= p.maxPoint.Y || _p.X < 0 || _p.Y < 0 {
		return nil, false
	}

	cell := p.cells[_p.Y*p.maxPoint.X+_p.X]

	if cell == nil {
		return nil, true
	}

	return cell, cell.Content() != p.empty
}

func (p *XYPlane) Mark(cell Cell) error {
	return p.markPoint(cell)
}

func (p *XYPlane) UnMark(cell Cell) error {
	return p.markPoint(NewColoredCell("empty", p.empty, cell.Color(), cell.Position()))
}

func (p *XYPlane) markPoint(cell Cell) error {
	_p := cell.Position()
	if _p.X >= p.maxPoint.X || _p.Y >= p.maxPoint.Y {
		return ErrPlaneOverflow
	}

	if _p.X < 0 || _p.Y < 0 {
		return ErrPlaneUnderflow
	}

	p.cells[_p.Y*p.maxPoint.X+_p.X] = cell
	return nil

}

func (p *XYPlane) MaxPoint() Point2 {
	return p.maxPoint
}
