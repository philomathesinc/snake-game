package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.Size{
		Width:  800,
		Height: 800,
	})

	w.CenterOnScreen()

	gameWindow := getGameWindow(20, 20)

	w.SetPadded(false)
	w.SetContent(gameWindow)
	w.ShowAndRun()
}

func getGameWindow(rows, cols int) fyne.CanvasObject {
	var rectangles []fyne.CanvasObject
	for i := 0; i < rows*cols; i++ {
		rect := canvas.NewRectangle(color.White)
		rectangles = append(rectangles, rect)
	}

	return container.NewAdaptiveGrid(20, rectangles...)

}
