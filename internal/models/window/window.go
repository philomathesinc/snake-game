package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/PhilomathesInc/snake-game/internal/utils"
)

const (
	length      = 840
	singlePixel = 40
)

type Window struct {
	fyne.Window
}

func (w *Window) RandomPosition() fyne.Position {
	var (
		i               fyne.Position
		pixelCountLimit = length / singlePixel
	)
	xPos := utils.RandomNumber(pixelCountLimit)
	yPos := utils.RandomNumber(pixelCountLimit)
	i = fyne.NewPos(float32(xPos), float32(yPos))

	return i
}

func New(a fyne.App) *Window {
	w := a.NewWindow("Snake Game")
	w.Resize(fyne.Size{
		Width:  length,
		Height: length,
	})
	w.CenterOnScreen()

	return &Window{w}
}

func (w *Window) UpdateContent(objs ...fyne.CanvasObject) {
	w.SetContent(container.NewWithoutLayout(objs...))
}

func (w *Window) PixelSize() int {
	return singlePixel
}

func (w *Window) CenterPosition() fyne.Position {
	return fyne.NewPos((length-singlePixel)/2, (length-singlePixel)/2)
}
