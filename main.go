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
	finalSpaceWidth  = 400
	finalSpaceHeight = 400
	singlePix        = 40
)

var (
	green = color.NRGBA{R: 0, G: 180, B: 0, A: 255}
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.Size{
		Width:  finalSpaceWidth,
		Height: finalSpaceHeight,
	})

	w.CenterOnScreen()

	snake := canvas.NewRectangle(green)
	snake.Resize(fyne.NewSize(singlePix, singlePix))

	content := container.NewWithoutLayout(snake)

	w.SetContent(content)

	fmt.Println(snake.Position())

	go func() {
		for {
			time.Sleep(1 * time.Second)
			newPos := fyne.NewPos(snake.Position().X+singlePix, snake.Position().Y+singlePix)
			if newPos.X > finalSpaceWidth || newPos.Y > finalSpaceHeight {
				fmt.Println("Game over")
				os.Exit(0)
			}
			snake.Move(newPos)
			fmt.Println(snake.Position())
			snake.Refresh()
		}
	}()

	w.ShowAndRun()
}
