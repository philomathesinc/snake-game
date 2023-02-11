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

	// Need 1 unit of rect on the window as Snake.
	var rectangles []*canvas.Rectangle
	for i := 1; i <= 20; i++ {
		rect := canvas.NewRectangle(color.White)
		rect.SetMinSize(fyne.Size{
			Width:  40,
			Height: 40,
		})
		rect.Resize(fyne.Size{
			Width:  40,
			Height: 40,
		})
		rectangles = append(rectangles, rect)
	}

	// snake := canvas.NewRectangle(color.White)
	// snake.SetMinSize(fyne.Size{
	// 	Width:  20,
	// 	Height: 20,
	// })

	// // snake.Move()
	// snake.Resize(fyne.NewSize(50, 50))

	content := container.NewWithoutLayout(rectangles...)

	w.SetContent(content)
	w.ShowAndRun()
}
