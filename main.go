package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	isPaused        bool
}

type snakeNode struct {
	direction string
	snakeObj  canvas.Rectangle
	next      *snakeNode
}

type snake struct {
	head   *snakeNode
	tail   *snakeNode
	length int
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
	content := container.NewWithoutLayout(
		&gameInstance.snakeInstance.head.snakeObj,
		gameInstance.pellet,
		gameInstance.scoreDisplayBox,
	)
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
	case fyne.KeyW, fyne.KeyUp:
		gameInstance.snakeInstance.head.direction = "up"
	case fyne.KeyS, fyne.KeyDown:
		gameInstance.snakeInstance.head.direction = "down"
	case fyne.KeyA, fyne.KeyLeft:
		gameInstance.snakeInstance.head.direction = "left"
	case fyne.KeyD, fyne.KeyRight:
		gameInstance.snakeInstance.head.direction = "right"
	case fyne.KeySpace, fyne.KeyP:
		gamePaused()
		// While paused the printKeys function is still called and the direction of the snake is updated even while paused.
	}
}

func gamePaused() {
	if gameInstance.isPaused {
		gameInstance.isPaused = false
		return
	}

	gameInstance.isPaused = true
}

func gameLoop() {
	for {
		time.Sleep(time.Millisecond * 200)

		if gameInstance.isPaused {
			continue
		}

		switch gameInstance.snakeInstance.head.direction {
		case "up":
			oldPos := gameInstance.snakeInstance.head.snakeObj.Position()
			// headNode move
			newPos := fyne.NewPos(
				gameInstance.snakeInstance.head.snakeObj.Position().X,
				gameInstance.snakeInstance.head.snakeObj.Position().Y-singlePix,
			)
			gameInstance.snakeInstance.head.snakeObj.Move(newPos)
			// rest of the snake body move
			updateSnakeBody(oldPos)
		case "down":
			oldPos := gameInstance.snakeInstance.head.snakeObj.Position()
			// headNode move
			newPos := fyne.NewPos(
				gameInstance.snakeInstance.head.snakeObj.Position().X,
				gameInstance.snakeInstance.head.snakeObj.Position().Y+singlePix,
			)
			gameInstance.snakeInstance.head.snakeObj.Move(newPos)
			// rest of the snake body move
			updateSnakeBody(oldPos)
		case "left":
			oldPos := gameInstance.snakeInstance.head.snakeObj.Position()
			// headNode move
			newPos := fyne.NewPos(
				gameInstance.snakeInstance.head.snakeObj.Position().X-singlePix,
				gameInstance.snakeInstance.head.snakeObj.Position().Y,
			)
			gameInstance.snakeInstance.head.snakeObj.Move(newPos)
			// rest of the snake body move
			updateSnakeBody(oldPos)
		case "right":
			oldPos := gameInstance.snakeInstance.head.snakeObj.Position()
			// headNode move
			newPos := fyne.NewPos(
				gameInstance.snakeInstance.head.snakeObj.Position().X+singlePix,
				gameInstance.snakeInstance.head.snakeObj.Position().Y,
			)
			gameInstance.snakeInstance.head.snakeObj.Move(newPos)
			// rest of the snake body move
			updateSnakeBody(oldPos)
		}

		// Snake dies on touching it's own body.
		if snakeBodyHit() {
			gameOver()
		}

		// Snake dies on touching the game window.
		if !windowHit() {
			gameOver()
		}

		// Score goes up by one when snake head touches it.
		if pelletHit() {
			gameInstance.pellet = foodPellet()
			gameInstance.score++
			gameInstance.scoreDisplayBox = canvas.NewText(
				fmt.Sprintf("Score: %d", gameInstance.score),
				color.White,
			)
			gameInstance.window.SetContent(
				container.NewWithoutLayout(
					&gameInstance.snakeInstance.head.snakeObj,
					gameInstance.pellet,
					gameInstance.scoreDisplayBox,
				),
			)
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

func snakeBodyHit() bool {
	for node := gameInstance.snakeInstance.head.next; node != nil; node = node.next {
		if gameInstance.snakeInstance.head.snakeObj.Position() == node.snakeObj.Position() {
			return true
		}
	}
	return false
}

func windowHit() bool {
	return !((gameInstance.snakeInstance.head.snakeObj.Position().Y == finalSpaceHeight) || (gameInstance.snakeInstance.head.snakeObj.Position().X == finalSpaceWidth) || (gameInstance.snakeInstance.head.snakeObj.Position().X < 0) || (gameInstance.snakeInstance.head.snakeObj.Position().Y < 0))
}

func pelletHit() bool {
	return gameInstance.snakeInstance.head.snakeObj.Position() == gameInstance.pellet.Position()
}

func gameOver() {
	fmt.Println("Game over!!")
	text1 := canvas.NewText("Game Over!!", color.White)
	text2 := canvas.NewText(fmt.Sprintf("SCORE : %d", gameInstance.score), color.White)
	text1.Alignment = fyne.TextAlignCenter
	text2.Alignment = fyne.TextAlignCenter
	score := container.NewVBox(text1, text2)
	content := container.New(layout.NewCenterLayout(), score)
	gameInstance.window.SetContent(content)
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
		snakeObj: canvas.Rectangle{
			FillColor:   green,
			StrokeColor: color.White,
			StrokeWidth: 1,
		},
	}
	snakeNode.next = nil
	snakeNode.snakeObj.Resize(fyne.NewSize(singlePix, singlePix))

	return &snakeNode
}

func updateSnakeBody(headOldPos fyne.Position) {
	oldPos := headOldPos
	tmp := gameInstance.snakeInstance.head.next

	for tmp != nil {
		olderPosition := tmp.snakeObj.Position()
		tmp.snakeObj.Move(oldPos)
		oldPos = olderPosition
		gameInstance.window.Canvas().Refresh(&tmp.snakeObj)
		tmp = tmp.next
	}

	i := 0
	for node := gameInstance.snakeInstance.head; node != nil; node = node.next {
		fmt.Printf("node %v: %v, %v\n", i, node.snakeObj.Position().X, node.snakeObj.Position().Y)
		i++
	}
}

func increaseSnakeLength() {
	snake := gameInstance.snakeInstance

	newNode := newSnakeNode()
	snake.tail.next = newNode
	snake.tail = snake.tail.next
	snake.length++

	updateSnakeBody(snake.head.snakeObj.Position())
	gameInstance.snakeInstance = snake

	objs := []fyne.CanvasObject{}
	for node := gameInstance.snakeInstance.head; node != nil; node = node.next {
		objs = append(objs, &node.snakeObj)
	}
	objs = append(objs, gameInstance.pellet)
	objs = append(objs, gameInstance.scoreDisplayBox)
	gameInstance.window.SetContent(container.NewWithoutLayout(objs...))
}
