package models

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/PhilomathesInc/snake-game/internal/constants"
	"github.com/PhilomathesInc/snake-game/internal/utils"
)

type Game struct {
	window          GameWindow
	SnakeInstance   Snake
	score           uint
	Pellet          fyne.CanvasObject
	ScoreDisplayBox *canvas.Text
}

func NewGame(
	window GameWindow,
	snakeInstance Snake, score uint, scoreDisplayBox *canvas.Text,
) Game {
	return Game{
		window:          window,
		SnakeInstance:   snakeInstance,
		score:           score,
		Pellet:          nil,
		ScoreDisplayBox: scoreDisplayBox,
	}
}

func (g *Game) SteerSnake(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyW:
		g.SnakeInstance.head.direction = "up"
	case fyne.KeyUp:
		g.SnakeInstance.head.direction = "up"
	case fyne.KeyS:
		g.SnakeInstance.head.direction = "down"
	case fyne.KeyDown:
		g.SnakeInstance.head.direction = "down"
	case fyne.KeyA:
		g.SnakeInstance.head.direction = "left"
	case fyne.KeyLeft:
		g.SnakeInstance.head.direction = "left"
	case fyne.KeyD:
		g.SnakeInstance.head.direction = "right"
	case fyne.KeyRight:
		g.SnakeInstance.head.direction = "right"
	}
}

func FoodPellet(g *Game) fyne.CanvasObject {
	pellet := *canvas.NewCircle(color.White)
	pellet.Resize(fyne.NewSize(constants.SinglePix, constants.SinglePix))

	pellet.Move(g.randomPositionInGameWindow())

	return &pellet
}

func (g *Game) randomPositionInGameWindow() fyne.Position {
	var i fyne.Position
	xPos := utils.RandomNumber(constants.PixelCountLimit)
	yPos := utils.RandomNumber(constants.PixelCountLimit)
	i = fyne.NewPos(float32(xPos), float32(yPos))
	for node := g.SnakeInstance.head; node != nil; node = node.next {
		if i == node.canvasObj.Position() {
			xPos = utils.RandomNumber(constants.PixelCountLimit)
			yPos = utils.RandomNumber(constants.PixelCountLimit)
			i = fyne.NewPos(float32(xPos), float32(yPos))
		}
	}
	fmt.Println("food pellet position:", i)
	return i
}

func (g *Game) GameLoop() {
	for {
		time.Sleep(time.Millisecond * 200)

		switch g.SnakeInstance.head.direction {
		case "up":
			oldPos := g.SnakeInstance.head.canvasObj.Position()
			// headNode move
			newPos := fyne.NewPos(g.SnakeInstance.head.canvasObj.Position().X, g.SnakeInstance.head.canvasObj.Position().Y-constants.SinglePix)
			g.SnakeInstance.head.canvasObj.Move(newPos)
			// rest of the snake body move
			g.updateSnakeBody(oldPos)
		case "down":
			oldPos := g.SnakeInstance.head.canvasObj.Position()
			// headNode move
			newPos := fyne.NewPos(g.SnakeInstance.head.canvasObj.Position().X, g.SnakeInstance.head.canvasObj.Position().Y+constants.SinglePix)
			g.SnakeInstance.head.canvasObj.Move(newPos)
			// rest of the snake body move
			g.updateSnakeBody(oldPos)
		case "left":
			oldPos := g.SnakeInstance.head.canvasObj.Position()
			// headNode move
			newPos := fyne.NewPos(g.SnakeInstance.head.canvasObj.Position().X-constants.SinglePix, g.SnakeInstance.head.canvasObj.Position().Y)
			g.SnakeInstance.head.canvasObj.Move(newPos)
			// rest of the snake body move
			g.updateSnakeBody(oldPos)
		case "right":
			oldPos := g.SnakeInstance.head.canvasObj.Position()
			// headNode move
			newPos := fyne.NewPos(g.SnakeInstance.head.canvasObj.Position().X+constants.SinglePix, g.SnakeInstance.head.canvasObj.Position().Y)
			g.SnakeInstance.head.canvasObj.Move(newPos)
			// rest of the snake body move
			g.updateSnakeBody(oldPos)
		}

		// Snake dies on touching it's own body.
		if g.SnakeInstance.SnakeBodyHit() {
			g.gameOver()
		}

		// Snake dies on touching the game window.
		if g.windowHit() {
			g.gameOver()
		}

		// Score goes up by one when snake head touches it.
		if g.pelletHit() {
			g.Pellet = FoodPellet(g)
			g.score++
			g.ScoreDisplayBox = canvas.NewText(fmt.Sprintf("Score: %d", g.score), color.White)
			g.window.SetContent(container.NewWithoutLayout(&g.SnakeInstance.head.canvasObj, g.Pellet, g.ScoreDisplayBox))
			g.increaseSnakeLength()
		}

		for node := g.SnakeInstance.head; node != nil; node = node.next {
			g.window.Canvas().Refresh(&node.canvasObj)
		}
	}
}

func (g *Game) pelletHit() bool {
	return g.SnakeInstance.head.canvasObj.Position() == g.Pellet.Position()
}

func (g *Game) windowHit() bool {
	return ((g.SnakeInstance.head.canvasObj.Position().Y == constants.FinalSpaceHeight) || (g.SnakeInstance.head.canvasObj.Position().X == constants.FinalSpaceWidth) || (g.SnakeInstance.head.canvasObj.Position().X < 0) || (g.SnakeInstance.head.canvasObj.Position().Y < 0))
}

func (g *Game) gameOver() {
	fmt.Println("Game over!!")
	os.Exit(0)
}
