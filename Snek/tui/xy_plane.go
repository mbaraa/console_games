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
	cells         []ColoredCell
}

func NewXYPlane(maxPoint Point2, filled, empty string) *XYPlane {
	cells := make([]ColoredCell, maxPoint.X*maxPoint.Y)
	for i := range cells {
		cells[i] = NewCell(empty, ColorReset)
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
			sb.WriteString(p.cells[row*p.maxPoint.X+col].String())
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (p *XYPlane) Marked(_p Point2) (TermColor, bool) {
	if _p.X >= p.maxPoint.X || _p.Y >= p.maxPoint.Y || _p.X < 0 || _p.Y < 0 {
		return ColorReset, false
	}

	cell := p.cells[_p.Y*p.maxPoint.X+_p.X]

	return cell.Color(), cell.Content() != p.empty
}

func (p *XYPlane) Mark(_p Point2, color ...TermColor) error {
	if len(color) != 0 {
		return p.markPoint(_p, NewCell(p.filled, color[0]))
	}
	return p.markPoint(_p, NewCell(p.filled, ColorReset))
}

func (p *XYPlane) UnMark(_p Point2) error {
	return p.markPoint(_p, NewCell(p.empty, ColorReset))
}

func (p *XYPlane) markPoint(_p Point2, cell ColoredCell) error {
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
