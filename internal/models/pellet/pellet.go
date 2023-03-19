package pellet

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Pellet struct {
	*canvas.Circle
}

func New(s int, p fyne.Position) *Pellet {
	c := canvas.NewCircle(color.White)
	c.Resize(fyne.NewSize(
		float32(s),
		float32(s),
	))
	c.Move(p)

	return &Pellet{
		c,
	}
}

func (p *Pellet) Display() fyne.CanvasObject {
	return p.Circle
}

func (p *Pellet) Hit(pos fyne.Position) bool {
	return pos == p.Position()
}
