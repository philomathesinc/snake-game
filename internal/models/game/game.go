package game

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/PhilomathesInc/snake-game/internal/models/pellet"
	"github.com/PhilomathesInc/snake-game/internal/models/scorecounter"
	"github.com/PhilomathesInc/snake-game/internal/models/snake"
	"github.com/PhilomathesInc/snake-game/internal/models/window"
)

type Game struct {
	window *window.Window
	snake  *snake.Snake
	score  *scorecounter.ScoreCounter
	pellet *pellet.Pellet
}

func New() *Game {
	w := window.New(app.New())
	snakePos := w.CenterPosition()
	pelletPos := w.RandomPosition()
	s := snake.New(w.PixelSize(), snakePos)
	p := pellet.New(w.PixelSize(), pelletPos)

	fmt.Printf("Snake pos: %v, Pellet pos: %v", snakePos, pelletPos)
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

func Start() {
	// Window initialization
	rand.Seed(time.Now().UnixNano())
	g := New()
	g.window.Canvas().SetOnTypedKey(g.steerSnake)
	g.window.UpdateContent(g.canvasObjects()...)
	go g.gameLoop()
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
	}
}

func (g *Game) gameLoop() {
	for {
		time.Sleep(time.Millisecond * 200)

		// Pellet Consumption - Score goes up by one when snake head touches it.
		if g.pellet.Hit(g.snake.HeadPosition()) {
			g.pellet = pellet.New(g.window.PixelSize(), g.window.RandomPosition())
			g.snake.Grow()
			g.score.Increment()

			objs := g.snake.BodyPositions()
			objs = append(objs, g.pellet.Display(), g.score.Display())
			g.window.UpdateContent(objs...)
		}
		// Snake dies on touching the game window
		if g.window.Hit(g.snake.HeadPosition()) {
			fmt.Println("DEBUG: inside g.window.Hit")
			over()
		}
		// Snake dies on touching it's own body.
		if g.snake.SnakeBodyHit() {
			over()
		}

		g.moveSnake()
		g.window.UpdateContent(g.canvasObjects()...)
	}
}

func (g *Game) moveSnake() {
	// move headNode
	var newPos fyne.Position
	switch g.snake.Direction() {
	case "up":
		newPos = fyne.NewPos(
			g.snake.HeadPosition().X,
			g.snake.HeadPosition().Y-float32(g.window.PixelSize()),
		)
	case "down":
		newPos = fyne.NewPos(
			g.snake.HeadPosition().X,
			g.snake.HeadPosition().Y+float32(g.window.PixelSize()),
		)
	case "left":
		newPos = fyne.NewPos(
			g.snake.HeadPosition().X-float32(g.window.PixelSize()),
			g.snake.HeadPosition().Y,
		)
	case "right":
		newPos = fyne.NewPos(
			g.snake.HeadPosition().X+float32(g.window.PixelSize()),
			g.snake.HeadPosition().Y,
		)
	}

	g.snake.Move(newPos)

	// ToDo: Refresh should be in window package
	// for node := g.SnakeInstance.head; node != nil; node = node.next {
	// 	g.window.Canvas().Refresh(&node.canvasObj)
	// }

}

func over() {
	fmt.Println("Game over!!")
	os.Exit(0)
}
