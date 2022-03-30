package shapes

import "snek/tui"

type Shape interface {
	GetPoints() []tui.Point2
}

func makeLineHorz(start, end tui.Point2) (pts []tui.Point2) {
	if start.X > end.X {
		start, end = end, start
	}
	for x := start.X; x <= end.X; x++ {
		pts = append(pts, tui.Point2{X: x, Y: start.Y})
	}
	return
}

func makeLineVert(start, end tui.Point2) (pts []tui.Point2) {
	if start.Y > end.Y {
		start, end = end, start
	}
	for y := start.Y; y <= end.Y; y++ {
		pts = append(pts, tui.Point2{X: start.X, Y: y})
	}
	return
}

func GetFilledPoints(start, end tui.Point2) []tui.Point2 {
	topLeftPoint := tui.Point2{X: start.X, Y: end.Y}
	bottomRightPoint := tui.Point2{X: end.X, Y: start.Y}

	bottomLine := makeLineHorz(start, bottomRightPoint)
	rightLine := makeLineVert(bottomRightPoint, end)
	topLine := makeLineHorz(end, topLeftPoint)
	leftLine := makeLineVert(topLeftPoint, start)

	return append(bottomLine, append(rightLine, append(topLine, leftLine...)...)...)
}

type Rect struct {
	pts []tui.Point2
}

func NewRect(start, end tui.Point2) *Rect {
	return &Rect{
		pts: GetFilledPoints(start, end),
	}
}

func (r *Rect) GetPoints() []tui.Point2 {
	return r.pts
}
