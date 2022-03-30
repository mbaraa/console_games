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
	cells         []string
}

func NewXYPlane(maxPoint Point2, filled, empty string) *XYPlane {
	cells := make([]string, maxPoint.X*maxPoint.Y)
	for i := range cells {
		cells[i] = empty
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
			sb.WriteString(p.cells[row*p.maxPoint.X+col])
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (p *XYPlane) Marked(_p Point2) bool {
	if _p.X >= p.maxPoint.X || _p.Y >= p.maxPoint.Y || _p.X < 0 || _p.Y < 0 {
		return false
	}

	return p.cells[_p.Y*p.maxPoint.X+_p.X] != p.empty
}

func (p *XYPlane) Mark(_p Point2, color ...TermColor) error {
	if len(color) != 0 {
		return p.markPoint(_p, string(color[0])+p.filled+string(ColorReset))
	}
	return p.markPoint(_p, p.filled)
}

func (p *XYPlane) UnMark(_p Point2) error {
	return p.markPoint(_p, p.empty)
}

func (p *XYPlane) markPoint(_p Point2, cell string) error {
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
