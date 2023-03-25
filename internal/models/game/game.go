package game

import (
	"context"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/PhilomathesInc/snake-game/internal/models/pellet"
	"github.com/PhilomathesInc/snake-game/internal/models/scorecounter"
	"github.com/PhilomathesInc/snake-game/internal/models/snake"
	"github.com/PhilomathesInc/snake-game/internal/models/window"
)

type Game struct {
	window   *window.Window
	snake    *snake.Snake
	score    *scorecounter.ScoreCounter
	pellet   *pellet.Pellet
	isPaused bool
	isOver   bool
}

func New(app fyne.App) *Game {
	w := window.New(app)
	s := snake.New(w.PixelSize(), w.CenterPosition())
	p := pellet.New(w.PixelSize(), w.RandomPosition())
	sc := scorecounter.New()

	return &Game{
		window: w,
		snake:  s,
		score:  sc,
		pellet: p,
	}
}

func (g *Game) canvasObjects() []fyne.CanvasObject {
	objs := g.snake.BodyPositions()
	objs = append(objs, g.pellet.Display(), g.score.Display())
	return objs
}

func Start(ctx context.Context, app fyne.App) {
	// Window initialization
	g := New(app)
	// Set the event handler for key presses
	g.window.Canvas().SetOnTypedKey(g.steerSnake)
	// Run the game loop
	go g.gameLoop()
	// Show the window
	g.window.ShowAndRun()
}

func (g *Game) steerSnake(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyW, fyne.KeyUp:
		g.snake.SetDirection("up")
	case fyne.KeyS, fyne.KeyDown:
		g.snake.SetDirection("down")
	case fyne.KeyA, fyne.KeyLeft:
		g.snake.SetDirection("left")
	case fyne.KeyD, fyne.KeyRight:
		g.snake.SetDirection("right")
	case fyne.KeySpace, fyne.KeyP:
		g.pause()
		// While paused the printKeys function is still called and the direction of the snake is updated even while paused.
	}
}

func (g *Game) pause() {
	if g.isPaused {
		g.isPaused = false
		return
	}

	g.isPaused = true
}

func (g *Game) gameLoop() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		// Game is paused or over.
		if g.isPaused || g.isOver {
			ticker.Stop()
		}

		g.snake.Move()
		g.window.UpdateContent(g.canvasObjects()...)

		// Pellet Consumption
		if g.pellet.Hit(g.snake.HeadPosition()) {
			g.pellet = pellet.New(g.window.PixelSize(), g.window.RandomPosition())
			g.snake.Grow()
			g.score.Increment()
			g.window.UpdateContent(g.canvasObjects()...)
		}
		// Snake hitting the window boundary.
		if g.window.Hit(g.snake.HeadPosition()) {
			g.over()
		}
		// Snake hitting itself.
		if g.snake.BodyHit() {
			g.over()
		}
	}
}

func (g *Game) over() {
	g.isOver = true
	text1 := canvas.NewText("Game Over", color.White)
	text2 := g.score.Display()
	gameOverContainer := container.NewVBox(text1, text2)
	content := container.New(layout.NewCenterLayout(), gameOverContainer)
	g.window.SetContent(content)
}
