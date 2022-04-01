package snakes

import (
	"container/list"
	"snek/tui"
)

type Snake struct {
	nodes         *list.List
	closeMoveChan func()
	keepMoving    bool
}

func NewSnake(initialPosition tui.Point2) *Snake {
	nodes := list.New()
	nodes.PushBack(initialPosition)

	return &Snake{
		nodes: nodes,
	}
}

func (s *Snake) Len() int {
	return s.nodes.Len()
}

func (s *Snake) AddNode() *Snake {
	p := s.nodes.Back().Value.(tui.Point2)
	p.X--
	s.nodes.PushBack(p)

	return s
}

func (s *Snake) Nodes() *list.List {
	return s.nodes
}

func (s *Snake) CheckMove() bool {
	return s.keepMoving
}

func (s *Snake) MoveY(steps int) (old, dst tui.Point2) {
	dstPt := s.nodes.Front().Value.(tui.Point2)
	dstPt.Y += steps
	return s.move(dstPt)
}

func (s *Snake) MoveX(steps int) (old, dst tui.Point2) {
	dstPt := s.nodes.Front().Value.(tui.Point2)
	dstPt.X += steps
	return s.move(dstPt)
}

func (s *Snake) move(p tui.Point2) (old, dst tui.Point2) {
	s.nodes.PushFront(p)
	last := s.nodes.Back()
	s.nodes.Remove(last)

	return last.Value.(tui.Point2), p
}
