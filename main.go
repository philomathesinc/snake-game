package main

import (
	"fmt"
	"image/color"
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
	window fyne.Window
	snake  canvas.Rectangle
}

var (
	green        = color.NRGBA{R: 0, G: 180, B: 0, A: 255}
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

	gameInstance.snake = *canvas.NewRectangle(green)
	gameInstance.snake.Resize(fyne.NewSize(singlePix, singlePix))

	gameInstance.snake.Move(fyne.NewPos((finalSpaceWidth-singlePix)/2, (finalSpaceHeight-singlePix)/2))

	content := container.NewWithoutLayout(&gameInstance.snake)
	w.SetContent(content)
	w.Canvas().SetOnTypedKey(printKeys)

	gameInstance.window = w

	go runAlways()
	w.ShowAndRun()
}

func printKeys(ev *fyne.KeyEvent) {
	if ev.Name == fyne.KeyW {
		fmt.Println("Move up")
		newPos := fyne.NewPos(gameInstance.snake.Position().X, gameInstance.snake.Position().Y-singlePix)
		gameInstance.snake.Move(newPos)
	} else if ev.Name == fyne.KeyS {
		newPos := fyne.NewPos(gameInstance.snake.Position().X, gameInstance.snake.Position().Y+singlePix)
		gameInstance.snake.Move(newPos)
	} else if ev.Name == fyne.KeyA {
		newPos := fyne.NewPos(gameInstance.snake.Position().X-singlePix, gameInstance.snake.Position().Y)
		gameInstance.snake.Move(newPos)
	} else if ev.Name == fyne.KeyD {
		newPos := fyne.NewPos(gameInstance.snake.Position().X+singlePix, gameInstance.snake.Position().Y)
		gameInstance.snake.Move(newPos)
	}

	gameInstance.window.Canvas().Refresh(&gameInstance.snake)
}

func runAlways() {
	for {
		time.Sleep(time.Second)
	}
}
