package snakes

import (
	"container/list"
	"snek/tui"
)

type SnakeCell struct {
	position tui.Point2
	content  string
	color    tui.TermColor
}

func NewSnakeCell(content string, color tui.TermColor, position tui.Point2) *SnakeCell {
	return &SnakeCell{
		position: position,
		color:    color,
		content:  content,
	}
}

func (s *SnakeCell) Position() tui.Point2 {
	return s.position
}

func (s *SnakeCell) Name() string {
	return "snake"
}

func (s *SnakeCell) Color() tui.TermColor {
	return s.color
}

func (s *SnakeCell) Content() string {
	return s.content
}

func (s *SnakeCell) String() string {
	return s.color.StringColored(s.content)
}

type Snake struct {
	nodes         *list.List
	closeMoveChan func()
	keepMoving    bool
}

func NewSnake(initialPosition tui.Point2) *Snake {
	nodes := list.New()
	nodes.PushBack(&SnakeCell{
		position: initialPosition,
		content:  "█",
		color:    tui.ColorBoldGreen,
	})

	return &Snake{
		nodes: nodes,
	}
}

func (s *Snake) Len() int {
	return s.nodes.Len()
}

func (s *Snake) AddNode() *Snake {
	c := s.nodes.Back().Value.(tui.Cell)
	s.nodes.PushBack(&SnakeCell{
		position: tui.Point2{
			X: c.Position().X - 1,
			Y: c.Position().Y,
		},
		content: "█",
		color:   c.Color(),
	})

	return s
}

func (s *Snake) Nodes() *list.List {
	return s.nodes
}

func (s *Snake) CheckMove() bool {
	return s.keepMoving
}

func (s *Snake) MoveY(steps int) (old, dst tui.Cell) {
	dstCell := s.nodes.Front().Value.(tui.Cell)
	dstPt := dstCell.Position()
	dstPt.Y += steps
	return s.move(NewSnakeCell("█", dstCell.Color(), dstPt))
}

func (s *Snake) MoveX(steps int) (old, dst tui.Cell) {
	dstCell := s.nodes.Front().Value.(tui.Cell)
	dstPt := dstCell.Position()
	dstPt.X += steps
	return s.move(NewSnakeCell("█", dstCell.Color(), dstPt))
}

func (s *Snake) move(c tui.Cell) (old, dst tui.Cell) {
	s.nodes.PushFront(c)
	last := s.nodes.Back()
	s.nodes.Remove(last)

	return last.Value.(tui.Cell), c
}
