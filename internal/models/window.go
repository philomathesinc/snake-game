package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/PhilomathesInc/snake-game/internal/constants"
)

type GameWindow struct {
	fyne.Window
}

func NewGameWindow(a fyne.App) GameWindow {
	w := a.NewWindow("Snake Game")
	w.Resize(fyne.Size{
		Width:  constants.FinalSpaceWidth,
		Height: constants.FinalSpaceHeight,
	})
	w.CenterOnScreen()

	return GameWindow{w}
}

func (w *GameWindow) UpdateContent(g *Game, s *Snake) {
	objs := s.BodyPositions()
	objs = append(objs, g.Pellet)
	objs = append(objs, g.ScoreDisplayBox)
	content := container.NewWithoutLayout(objs...)

	w.SetContent(content)
}
