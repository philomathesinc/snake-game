package snake

import (
	"fmt"
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
	head      *node
	tail      *node
	length    int
	pixelSize float32
}

func (s *Snake) newSnakeNode() *node {
	snakeNode := node{
		direction: "up",
		canvasObj: canvas.Rectangle{
			FillColor:   constants.Green,
			StrokeColor: color.White,
			StrokeWidth: 1,
		},
	}
	snakeNode.next = nil
	snakeNode.canvasObj.Resize(fyne.NewSize(s.pixelSize, s.pixelSize))

	return &snakeNode
}

func New(pixelSize int, position fyne.Position) *Snake {
	snake := Snake{}
	snake.pixelSize = float32(pixelSize)
	snake.head = snake.newSnakeNode()
	snake.tail = snake.head
	snake.length = 1
	return &snake
}

func (s *Snake) HeadPosition() fyne.Position {
	return s.head.canvasObj.Position()
}

func (s *Snake) Direction() string {
	return s.head.direction
}

func (s *Snake) SetDirection(d string) {
	s.head.direction = d
}

func (s *Snake) SnakeBodyHit() bool {
	for node := s.head.next; node != nil; node = node.next {
		if s.HeadPosition() == node.canvasObj.Position() {
			return true
		}
	}

	return false
}

func (s *Snake) BodyPositions() []fyne.CanvasObject {
	objs := []fyne.CanvasObject{}
	// Add all the snake nodes' CanvasObject to `objs`.
	for node := s.head; node != nil; node = node.next {
		objs = append(objs, &node.canvasObj)
	}
	return objs
}

func (s *Snake) Move(pos fyne.Position) {
	oldPos := s.HeadPosition()
	s.head.canvasObj.Move(pos)
	// rest of the snake body move
	s.updateSnakeBody(oldPos)
}

func (s *Snake) updateSnakeBody(headOldPos fyne.Position) {
	oldPos := headOldPos
	tmp := s.head.next

	for tmp != nil {
		olderPosition := tmp.canvasObj.Position()
		tmp.canvasObj.Move(oldPos)
		oldPos = olderPosition
		// ToDo: g.window.Canvas().Refresh(&tmp.canvasObj)
		tmp = tmp.next
	}

	i := 0
	for node := s.head; node != nil; node = node.next {
		fmt.Printf("node %v: %v\n", i, node.canvasObj.Position())
		i++
	}
}

func (s *Snake) Grow() {
	newNode := s.newSnakeNode()
	s.tail.next = newNode
	s.tail = s.tail.next
	s.length++

	s.updateSnakeBody(s.HeadPosition())
}
