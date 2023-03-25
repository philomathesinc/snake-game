package window

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

const (
	length      = 840
	singlePixel = 40
)

type Window struct {
	fyne.Window
}

func generateRandomNumber(limit int) int {
	var i int
	i = rand.Intn(limit)
	for i <= 1 {
		i = rand.Intn(limit)
	}
	return i * singlePixel
}

func (w *Window) RandomPosition() fyne.Position {

	var (
		i               fyne.Position
		pixelCountLimit = length/singlePixel - 1
	)

	xPos := generateRandomNumber(pixelCountLimit)
	yPos := generateRandomNumber(pixelCountLimit)
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
	for _, obj := range objs {
		obj.Refresh()
	}
}

func (w *Window) PixelSize() int {
	return singlePixel
}

func (w *Window) CenterPosition() fyne.Position {
	return fyne.NewPos((length-singlePixel)/2, (length-singlePixel)/2)
}

func (w *Window) Hit(p fyne.Position) bool {
	right := p.Y >= length
	fmt.Printf("p: %v\n", p)
	fmt.Println(right)
	left := p.X >= length
	fmt.Println(left)
	top := p.X <= 0
	fmt.Println(top)
	bottom := p.Y <= 0
	fmt.Println(bottom)

	// return !((gameInstance.snakeInstance.head.snakeObj.Position().Y == finalSpaceHeight) || (gameInstance.snakeInstance.head.snakeObj.Position().X == finalSpaceWidth) || (gameInstance.snakeInstance.head.snakeObj.Position().X < 0) || (gameInstance.snakeInstance.head.snakeObj.Position().Y < 0))
	return (right || left || top || bottom)
}
