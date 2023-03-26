package game

import (
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
	app      fyne.App
	window   *window.Window
	snake    *snake.Snake
	score    *scorecounter.ScoreCounter
	pellet   *pellet.Pellet
	isPaused bool
	isOver   bool
	speed    int
}

func New(app fyne.App) *Game {
	w := window.New(app, "Snake Game")
	s := snake.New(w.PixelSize(), w.CenterPosition())
	p := pellet.New(w.PixelSize(), w.RandomPosition())
	sc := scorecounter.New()

	return &Game{
		app:    app,
		window: w,
		snake:  s,
		score:  sc,
		pellet: p,
		speed:  0,
	}
}

func (g *Game) canvasObjects() []fyne.CanvasObject {
	objs := g.snake.BodyPositions()
	objs = append(objs, g.pellet.Display(), g.score.Display())
	return objs
}

func (g *Game) Start() {
	// Set the event handler for key presses
	g.window.Canvas().SetOnTypedKey(g.steerSnake)

	// Run the game loop
	go g.gameLoop()
	// 	Display the game window
	g.window.Show()
}

func (g *Game) steerSnake(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyW, fyne.KeyUp:
		if !g.isPaused {
			g.snake.SetDirection("up")
		}
	case fyne.KeyS, fyne.KeyDown:
		if !g.isPaused {
			g.snake.SetDirection("down")
		}
	case fyne.KeyA, fyne.KeyLeft:
		if !g.isPaused {
			g.snake.SetDirection("left")
		}
	case fyne.KeyD, fyne.KeyRight:
		if !g.isPaused {
			g.snake.SetDirection("right")
		}
	case fyne.KeySpace, fyne.KeyP:
		g.togglePause()
	}
}

func (g *Game) togglePause() {
	if g.isPaused {
		g.isPaused = false

		return
	}

	g.isPaused = true
}

func (g *Game) gameLoop() {
	ticker := time.NewTicker(200 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			// Game is paused.
			if g.isPaused {
				continue
			}

			// Game is over.
			if g.isOver {
				ticker.Stop()
				continue
			}

			// Move the snake
			g.snake.Move()

			// Update the window contents - snake, pellet, score
			g.window.UpdateContent(g.canvasObjects()...)

			// Pellet Consumption
			if g.pellet.Hit(g.snake.HeadPosition()) {
				g.increaseSpeed()
				tickerSpeed := (time.Duration(200.0-g.speed) * time.Millisecond)
				ticker = time.NewTicker(tickerSpeed)
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
}

func (g *Game) increaseSpeed() {
	switch speed := g.speed; {
	case speed < 50:
		g.speed += 10
	case speed < 150:
		g.speed += 5
	case speed <= 200:
		g.speed += 2
	}
}

func (g *Game) over() {
	// Set Game Over state
	g.isOver = true

	// Get updated score and increase font size
	scoreDisplay := g.score.Display()
	scoreDisplay.(*canvas.Text).TextSize = 50

	// Create and display the Game Over window
	gameOverContainer := container.NewVBox(scoreDisplay)
	content := container.New(layout.NewCenterLayout(), gameOverContainer)
	gameOverWindow := window.New(g.app, "Game Over")
	gameOverWindow.Resize(fyne.Size{
		Width:  300,
		Height: 100,
	})
	gameOverWindow.SetContent(content)
	gameOverWindow.Show()
}
