package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const (
	finalSpaceWidth  = 840
	finalSpaceHeight = 840
	singlePix        = 40
	pixelCountLimit  = 21
)

type game struct {
	window          fyne.Window
	snakeInstance   snake
	score           uint
	pellet          fyne.CanvasObject
	scoreDisplayBox *canvas.Text
}

type snakeNode struct {
	direction string
	position  fyne.Position
	snakeObj  canvas.Rectangle
	next      *snakeNode
}

type snake struct {
	head   *snakeNode
	tail   *snakeNode
	length int
	// body []fyne.Position
}

var (
	green        = color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	white        = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	gameInstance = game{}
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.Size{
		Width:  finalSpaceWidth,
		Height: finalSpaceHeight,
	})

	w.CenterOnScreen()

	gameInstance.snakeInstance = newSnake()
	centerGamePixel := fyne.NewPos((finalSpaceWidth-singlePix)/2, (finalSpaceHeight-singlePix)/2)
	gameInstance.snakeInstance.head.snakeObj.Move(centerGamePixel)

	gameInstance.pellet = foodPellet()
	gameInstance.scoreDisplayBox = canvas.NewText(fmt.Sprintf("Score: %d", 0), color.White)
	content := container.NewWithoutLayout(&gameInstance.snakeInstance.head.snakeObj, gameInstance.pellet, gameInstance.scoreDisplayBox)
	w.SetContent(content)
	w.Canvas().SetOnTypedKey(printKeys)

	gameInstance.window = w

	go gameLoop()
	w.ShowAndRun()
}

func foodPellet() fyne.CanvasObject {
	pellet := *canvas.NewCircle(white)
	pellet.Resize(fyne.NewSize(singlePix, singlePix))

	pellet.Move(randomPositionInGameWindow())

	return &pellet
}

func printKeys(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyW:
	case fyne.KeyUp:
		gameInstance.snakeInstance.head.direction = "up"
	case fyne.KeyS:
	case fyne.KeyDown:
		gameInstance.snakeInstance.head.direction = "down"
	case fyne.KeyA:
	case fyne.KeyLeft:
		gameInstance.snakeInstance.head.direction = "left"
	case fyne.KeyD:
	case fyne.KeyRight:
		gameInstance.snakeInstance.head.direction = "right"
	}
}

func gameLoop() {
	for {
		time.Sleep(time.Millisecond * 400)

		switch gameInstance.snakeInstance.head.direction {
		case "up":
			oldPos := gameInstance.snakeInstance.head.snakeObj.Position()
			// headNode move
			newPos := fyne.NewPos(gameInstance.snakeInstance.head.snakeObj.Position().X, gameInstance.snakeInstance.head.snakeObj.Position().Y-singlePix)
			gameInstance.snakeInstance.head.snakeObj.Move(newPos)
			updateSnakeBody(oldPos)
		case "down":
			oldPos := gameInstance.snakeInstance.head.snakeObj.Position()
			// headNode move
			newPos := fyne.NewPos(gameInstance.snakeInstance.head.snakeObj.Position().X, gameInstance.snakeInstance.head.snakeObj.Position().Y+singlePix)
			gameInstance.snakeInstance.head.snakeObj.Move(newPos)
			updateSnakeBody(oldPos)
		case "left":
			oldPos := gameInstance.snakeInstance.head.snakeObj.Position()
			// headNode move
			newPos := fyne.NewPos(gameInstance.snakeInstance.head.snakeObj.Position().X-singlePix, gameInstance.snakeInstance.head.snakeObj.Position().Y)
			gameInstance.snakeInstance.head.snakeObj.Move(newPos)
			updateSnakeBody(oldPos)
		case "right":
			oldPos := gameInstance.snakeInstance.head.snakeObj.Position()
			// headNode move
			newPos := fyne.NewPos(gameInstance.snakeInstance.head.snakeObj.Position().X+singlePix, gameInstance.snakeInstance.head.snakeObj.Position().Y)
			gameInstance.snakeInstance.head.snakeObj.Move(newPos)
			updateSnakeBody(oldPos)
		}

		// Snake dies on touching the game window.
		if !checkIfWindowHit() {
			gameOver()
		}

		// Score goes up by one when snake head touches it.
		if checkIfPelletHit() {
			gameInstance.pellet = foodPellet()
			gameInstance.score++
			gameInstance.scoreDisplayBox = canvas.NewText(fmt.Sprintf("Score: %d", gameInstance.score), color.White)
			gameInstance.window.SetContent(container.NewWithoutLayout(&gameInstance.snakeInstance.head.snakeObj, gameInstance.pellet, gameInstance.scoreDisplayBox))
			increaseSnakeLength()

			fmt.Printf("gameInstance.score: %v\n", gameInstance.score)
		}

		for node := gameInstance.snakeInstance.head; node != nil; node = node.next {
			gameInstance.window.Canvas().Refresh(&node.snakeObj)
		}
	}
}

func randomPositionInGameWindow() fyne.Position {
	var i fyne.Position
	xPos := randomNumber(pixelCountLimit)
	yPos := randomNumber(pixelCountLimit)
	i = fyne.NewPos(float32(xPos), float32(yPos))
	for node := gameInstance.snakeInstance.head; node != nil; node = node.next {
		if i == node.snakeObj.Position() {
			xPos = randomNumber(pixelCountLimit)
			yPos = randomNumber(pixelCountLimit)
			i = fyne.NewPos(float32(xPos), float32(yPos))
		}
	}
	fmt.Println("food pellet position:", i)
	return i
}

func randomNumber(limit int) int {
	var i int
	i = rand.Intn(limit)
	for i <= 1 {
		i = rand.Intn(limit)
	}
	return i * singlePix
}

func checkIfWindowHit() bool {
	return !((gameInstance.snakeInstance.head.snakeObj.Position().Y == finalSpaceHeight) || (gameInstance.snakeInstance.head.snakeObj.Position().X == finalSpaceWidth) || (gameInstance.snakeInstance.head.snakeObj.Position().X < 0) || (gameInstance.snakeInstance.head.snakeObj.Position().Y < 0))
}

func checkIfPelletHit() bool {
	return gameInstance.snakeInstance.head.snakeObj.Position() == gameInstance.pellet.Position()
}

func gameOver() {
	fmt.Println("Game over!!")
	os.Exit(0)
}

func newSnake() snake {
	snake := snake{}
	snake.head = newSnakeNode()
	snake.tail = snake.head

	return snake
}

func newSnakeNode() *snakeNode {
	snakeNode := snakeNode{
		direction: "up",
		snakeObj:  *canvas.NewRectangle(green),
	}
	snakeNode.next = nil
	snakeNode.snakeObj.Resize(fyne.NewSize(singlePix, singlePix))

	return &snakeNode
}

func updateSnakeBody(oldPos fyne.Position) {
	tmp := gameInstance.snakeInstance.head.next

	for tmp != nil {
		tmp.next.snakeObj.Move(oldPos)
		oldPos = tmp.snakeObj.Position()
		tmp = tmp.next
	}
}

func increaseSnakeLength() {
	node := newSnakeNode()
	snake := gameInstance.snakeInstance
	headPos := snake.head.snakeObj.Position()
	gameInstance.snakeInstance.head.snakeObj.Move(headPos)

	snake.tail.next = node
	snake.tail = snake.tail.next
}
