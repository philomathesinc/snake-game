package models

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"github.com/PhilomathesInc/snake-game/internal/constants"
)

type node struct {
	direction string
	canvasObj canvas.Rectangle
	next      *node
}

type Snake struct {
	head   *node
	tail   *node
	length int
}

func newSnakeNode() *node {
	snakeNode := node{
		direction: "up",
		canvasObj: canvas.Rectangle{
			FillColor:   constants.Green,
			StrokeColor: color.White,
			StrokeWidth: 1,
		},
	}
	snakeNode.next = nil
	snakeNode.canvasObj.Resize(fyne.NewSize(constants.SinglePix, constants.SinglePix))

	return &snakeNode
}

func (s *Snake) SnakeBodyHit() bool {
	for node := s.head.next; node != nil; node = node.next {
		if s.head.canvasObj.Position() == node.canvasObj.Position() {
			return true
		}
	}

	return false
}

func NewSnake() Snake {
	snake := Snake{}
	snake.head = newSnakeNode()
	snake.tail = snake.head

	return snake
}

func (s *Snake) Move(pos fyne.Position) {
	s.head.canvasObj.Move(pos)
}

func (s *Snake) BodyPositions() []fyne.CanvasObject {
	objs := []fyne.CanvasObject{}
	// Add all the snake nodes' CanvasObject to `objs`.
	for node := s.head; node != nil; node = node.next {
		objs = append(objs, &node.canvasObj)
	}
	return objs
}
