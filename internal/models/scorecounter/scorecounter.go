package scorecounter

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type ScoreCounter struct {
	display *canvas.Text
	count   int
}

func New() *ScoreCounter {
	return &ScoreCounter{
		display: canvas.NewText(fmt.Sprintf("Score: %d", 0), color.White),
		count:   0,
	}
}

func (sc *ScoreCounter) Increment() {
	sc.count++
}

func (sc *ScoreCounter) Display() fyne.CanvasObject {
	sc.display = canvas.NewText(fmt.Sprintf("Score: %d", sc.count), color.White)

	return sc.display
}
