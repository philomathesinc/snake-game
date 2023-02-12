package main

import (
	"image/color"

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

	w.ShowAndRun()
}

func printKeys(ev *fyne.KeyEvent) {
	if ev.Name == fyne.KeyW {
		gameInstance.snake.Position().AddXY(0, singlePix)
	} else if ev.Name == fyne.KeyS {
		gameInstance.snake.Position().AddXY(0, -singlePix)
	} else if ev.Name == fyne.KeyA {
		gameInstance.snake.Position().AddXY(-singlePix, 0)
	} else if ev.Name == fyne.KeyD {
		gameInstance.snake.Position().AddXY(singlePix, 0)
	}

	gameInstance.window.Canvas().Refresh(&gameInstance.snake)
}
