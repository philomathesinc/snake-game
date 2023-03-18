package window

import (
	"fyne.io/fyne/v2"
	"github.com/PhilomathesInc/snake-game/internal/utils"
)

const (
	length      = 840
	singlePixel = 40
)

var (
	centerGamePixel = fyne.NewPos((length-singlePixel)/2, (length-singlePixel)/2)
	pixelCountLimit = length / singlePixel
)

type Window struct {
	fyne.Window
}

func (w *Window) RandomPosition() fyne.Position {
	var i fyne.Position
	xPos := utils.RandomNumber(pixelCountLimit)
	yPos := utils.RandomNumber(pixelCountLimit)
	i = fyne.NewPos(float32(xPos), float32(yPos))

	return i
}

// func NewGameWindow(a fyne.App) GameWindow {
// 	w := a.NewWindow("Snake Game")
// 	w.Resize(fyne.Size{
// 		Width:  constants.FinalSpaceWidth,
// 		Height: constants.FinalSpaceHeight,
// 	})
// 	w.CenterOnScreen()

// 	return GameWindow{w}
// }

// func (w *GameWindow) UpdateContent(g *Game, s *Snake) {
// 	objs := s.BodyPositions()
// 	objs = append(objs, g.Pellet)
// 	objs = append(objs, g.ScoreDisplayBox)
// 	content := container.NewWithoutLayout(objs...)

// 	w.SetContent(content)
// }
