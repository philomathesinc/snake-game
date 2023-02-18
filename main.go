package main

import (
	"fmt"
	"image/color"
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
	window        fyne.Window
	snakeInstance snake
	score         uint
	pellet        fyne.CanvasObject
}

type snake struct {
	direction string
	snakeObj  canvas.Rectangle
}

var (
	green        = color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	white        = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	gameInstance = game{}
)

func main() {
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
	gameInstance.snakeInstance.snakeObj.Move(fyne.NewPos((finalSpaceWidth-singlePix)/2, (finalSpaceHeight-singlePix)/2))

	gameInstance.pellet = foodPellet()
	content := container.NewWithoutLayout(&gameInstance.snakeInstance.snakeObj, gameInstance.pellet)
	w.SetContent(content)
	w.Canvas().SetOnTypedKey(printKeys)

	gameInstance.window = w

	go gameLoop()
	w.ShowAndRun()
}

func foodPellet() fyne.CanvasObject {
	pellet := *canvas.NewCircle(white)
	pellet.Resize(fyne.NewSize(singlePix, singlePix))
	pellet.Move(fyne.NewPos(80, 80))

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
		time.Sleep(time.Millisecond * 600)

		switch gameInstance.snakeInstance.direction {
		case "up":
			newPos := fyne.NewPos(gameInstance.snakeInstance.snakeObj.Position().X, gameInstance.snakeInstance.snakeObj.Position().Y-singlePix)
			gameInstance.snakeInstance.snakeObj.Move(newPos)
		case "down":
			newPos := fyne.NewPos(gameInstance.snakeInstance.snakeObj.Position().X, gameInstance.snakeInstance.snakeObj.Position().Y+singlePix)
			gameInstance.snakeInstance.snakeObj.Move(newPos)
		case "left":
			newPos := fyne.NewPos(gameInstance.snakeInstance.snakeObj.Position().X-singlePix, gameInstance.snakeInstance.snakeObj.Position().Y)
			gameInstance.snakeInstance.snakeObj.Move(newPos)
		case "right":
			newPos := fyne.NewPos(gameInstance.snakeInstance.snakeObj.Position().X+singlePix, gameInstance.snakeInstance.snakeObj.Position().Y)
			gameInstance.snakeInstance.snakeObj.Move(newPos)
		}

		gameInstance.window.Canvas().Refresh(&gameInstance.snakeInstance.snakeObj)

		// Snake dies on touching the game window.
		if !checkIfWindowHit() {
			gameOver()
		}

		// Score goes up by one when snake head touches it.
		if checkIfPelletHit() {
			gameInstance.score++

			fmt.Printf("gameInstance.score: %v\n", gameInstance.score)
		}
	}
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
