package pellet

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Pellet struct {
	*canvas.Circle
}

func New(s fyne.Size, p fyne.Position) *Pellet {
	c := canvas.NewCircle(color.White)
	c.Resize(s)
	c.Move(p)

	return &Pellet{
		c,
	}
}
