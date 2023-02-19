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
)

type game struct {
	window          fyne.Window
	snakeInstance   snake
	score           uint
	pellet          fyne.CanvasObject
	scoreDisplayBox *canvas.Text
}

type snake struct {
	body      []fyne.Position
	direction string
	snakeObj  canvas.Rectangle
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

	gameInstance.snakeInstance = snake{
		direction: "up",
		snakeObj:  *canvas.NewRectangle(green),
	}
	gameInstance.snakeInstance.snakeObj.Resize(fyne.NewSize(singlePix, singlePix))
	centerGamePixel := fyne.NewPos((finalSpaceWidth-singlePix)/2, (finalSpaceHeight-singlePix)/2)
	gameInstance.snakeInstance.body = append(gameInstance.snakeInstance.body, centerGamePixel)
	gameInstance.snakeInstance.snakeObj.Move(centerGamePixel)

	gameInstance.pellet = foodPellet()
	gameInstance.scoreDisplayBox = canvas.NewText(fmt.Sprintf("Score: %d", 0), color.White)
	content := container.NewWithoutLayout(&gameInstance.snakeInstance.snakeObj, gameInstance.pellet, gameInstance.scoreDisplayBox)
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
	if ev.Name == fyne.KeyW {
		gameInstance.snakeInstance.direction = "up"
	} else if ev.Name == fyne.KeyS {
		gameInstance.snakeInstance.direction = "down"
	} else if ev.Name == fyne.KeyA {
		gameInstance.snakeInstance.direction = "left"
	} else if ev.Name == fyne.KeyD {
		gameInstance.snakeInstance.direction = "right"
	}
}

func gameLoop() {
	for {
		time.Sleep(time.Millisecond * 300)

		switch gameInstance.snakeInstance.direction {
		case "up":
			newPos := fyne.NewPos(gameInstance.snakeInstance.snakeObj.Position().X, gameInstance.snakeInstance.snakeObj.Position().Y-singlePix)
			gameInstance.snakeInstance.snakeObj.Move(newPos)
			gameInstance.snakeInstance.body = append(gameInstance.snakeInstance.body[1:1], newPos)
		case "down":
			newPos := fyne.NewPos(gameInstance.snakeInstance.snakeObj.Position().X, gameInstance.snakeInstance.snakeObj.Position().Y+singlePix)
			gameInstance.snakeInstance.snakeObj.Move(newPos)
			gameInstance.snakeInstance.body = append(gameInstance.snakeInstance.body[1:1], newPos)
		case "left":
			newPos := fyne.NewPos(gameInstance.snakeInstance.snakeObj.Position().X-singlePix, gameInstance.snakeInstance.snakeObj.Position().Y)
			gameInstance.snakeInstance.snakeObj.Move(newPos)
			gameInstance.snakeInstance.body = append(gameInstance.snakeInstance.body[1:1], newPos)
		case "right":
			newPos := fyne.NewPos(gameInstance.snakeInstance.snakeObj.Position().X+singlePix, gameInstance.snakeInstance.snakeObj.Position().Y)
			gameInstance.snakeInstance.snakeObj.Move(newPos)
			gameInstance.snakeInstance.body = append(gameInstance.snakeInstance.body[1:1], newPos)
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
			gameInstance.window.SetContent(container.NewWithoutLayout(&gameInstance.snakeInstance.snakeObj, gameInstance.pellet, gameInstance.scoreDisplayBox))

			fmt.Printf("gameInstance.score: %v\n", gameInstance.score)
		}

		gameInstance.window.Canvas().Refresh(&gameInstance.snakeInstance.snakeObj)
	}
}

func randomPositionInGameWindow() fyne.Position {
	var i fyne.Position
	xPos := randomNumber(22)
	yPos := randomNumber(22)
	i = fyne.NewPos(float32(xPos), float32(yPos))
	for i == gameInstance.snakeInstance.body[0] {
		xPos = randomNumber(22)
		yPos = randomNumber(22)
		i = fyne.NewPos(float32(xPos), float32(yPos))
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
	return !((gameInstance.snakeInstance.snakeObj.Position().Y == finalSpaceHeight) || (gameInstance.snakeInstance.snakeObj.Position().X == finalSpaceWidth) || (gameInstance.snakeInstance.snakeObj.Position().X < 0) || (gameInstance.snakeInstance.snakeObj.Position().Y < 0))
}

func checkIfPelletHit() bool {
	return gameInstance.snakeInstance.snakeObj.Position() == gameInstance.pellet.Position()
}

func gameOver() {
	fmt.Println("Game over!!")
	os.Exit(0)
}

// Bug : food pellet generation stops at
// {640 840}
