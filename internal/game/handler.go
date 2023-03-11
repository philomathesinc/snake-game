package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	"github.com/PhilomathesInc/snake-game/internal/constants"
	"github.com/PhilomathesInc/snake-game/internal/models"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	w := models.NewGameWindow(app.New())

	s := models.NewSnake()
	scoreDisplayBox := canvas.NewText(fmt.Sprintf("Score: %d", 0), color.White)
	g := models.NewGame(
		w,
		s,
		0,
		scoreDisplayBox)

	p := models.FoodPellet(&g)
	g.Pellet = p

	// Put the Snake's head in the center of the game space
	s.Move(constants.CenterGamePixel)
	w.UpdateContent(&g, &s)
	w.Canvas().SetOnTypedKey(g.SteerSnake)

	go g.GameLoop()
	w.ShowAndRun()
}
