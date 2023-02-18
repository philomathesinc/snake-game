/*
This Go code is a simple implementation of the classic game "Snake" using the
Fyne GUI toolkit. The game window is set to a size of 840x840, and the game
space is divided into 21x21 grid squares of 40x40 pixels each.

The game starts with a single block of green color representing the snake's
head, with its position set to the center of the game window. The snake's body
is represented by a series of fyne.Position values that correspond to the grid
squares it occupies. The direction in which the snake moves is determined by
the WASD keys, and the snake's movement speed is controlled by a time.Sleep
function call in the game loop.

A white circle represents the food pellet, which is placed at a random position
on the game board. When the snake's head collides with the food pellet, the
pellet disappears, and a new one is placed at a different random position. The
score is incremented, and the snake grows by one block. If the snake's head
collides with any of the walls, the game is over.

The code uses Go's "fyne" package to create the graphical user interface and
"time" package to control the game's loop speed.
*/

package main

// Import the necessary packages, including fyne - the GUI
import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const (
    DELAY = 300
	finalSpaceWidth  = 840
	finalSpaceHeight = 840
	singlePix        = 40
)

// Types
type game struct {
	window        fyne.Window
	snakeInstance snake
	score         uint
	pellet        fyne.CanvasObject
}

type snake struct {
	body      []fyne.Position
	direction string
	snakeObj  canvas.Rectangle
}

// Global variables
var (
	green        = color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	white        = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	gameInstance = game{}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.Size{
		Width:  finalSpaceWidth,
		Height: finalSpaceHeight,
	})

	w.CenterOnScreen()

	gameInstance.snakeInstance = snake{
		direction: "up",
		snakeObj:  *canvas.NewRectangle(green),
	}
	gameInstance.snakeInstance.snakeObj.Resize(fyne.NewSize(singlePix, singlePix))
	centerGamePixel := fyne.NewPos((finalSpaceWidth-singlePix)/2, (finalSpaceHeight-singlePix)/2)
	gameInstance.snakeInstance.body = append(gameInstance.snakeInstance.body, centerGamePixel)
	gameInstance.snakeInstance.snakeObj.Move(centerGamePixel)

	gameInstance.pellet = foodPellet()
	content := container.NewWithoutLayout(&gameInstance.snakeInstance.snakeObj, gameInstance.pellet)
	w.SetContent(content)
	w.Canvas().SetOnTypedKey(changeDirection)

	gameInstance.window = w

	go gameLoop()
	w.ShowAndRun()
}

func foodPellet() fyne.CanvasObject {
	pellet := *canvas.NewCircle(white)
	pellet.Resize(fyne.NewSize(singlePix, singlePix))

	pellet.Move(randomPositionInGameWindow())

	return &pellet
}

// Set the snake direction according to the key pressed; from fyne.KeyEvent
func changeDirection(ev *fyne.KeyEvent) {
	if ev.Name == fyne.KeyW {
		gameInstance.snakeInstance.direction = "up"
	} else if ev.Name == fyne.KeyS {
		gameInstance.snakeInstance.direction = "down"
	} else if ev.Name == fyne.KeyA {
		gameInstance.snakeInstance.direction = "left"
	} else if ev.Name == fyne.KeyD {
		gameInstance.snakeInstance.direction = "right"
	}
}

/*The gameLoop function is the main game loop. It runs in an infinite loop and
handles the movement of the snake, the score, and the game over conditions.
In each iteration of the loop, the function waits for a few milliseconds using
the time.Sleep function, and then moves the snake in the direction specified
by the gameInstance.snakeInstance.direction variable.*/

func gameLoop() {
	for {
		time.Sleep(time.Millisecond * DELAY)

        // Based on the direction move the snake to the next position.
		switch gameInstance.snakeInstance.direction {
		case "up":
			newPos := fyne.NewPos(
				gameInstance.snakeInstance.snakeObj.Position().X,
				gameInstance.snakeInstance.snakeObj.Position().Y - singlePix,
			)
			gameInstance.snakeInstance.snakeObj.Move(newPos)
			gameInstance.snakeInstance.body = append(gameInstance.snakeInstance.body[1:1], newPos)
		case "down":
			newPos := fyne.NewPos(
                gameInstance.snakeInstance.snakeObj.Position().X,
				gameInstance.snakeInstance.snakeObj.Position().Y + singlePix,
            )
			gameInstance.snakeInstance.snakeObj.Move(newPos)
			gameInstance.snakeInstance.body = append(gameInstance.snakeInstance.body[1:1], newPos)
		case "left":
			newPos := fyne.NewPos(gameInstance.snakeInstance.snakeObj.Position().X-singlePix, gameInstance.snakeInstance.snakeObj.Position().Y)
			gameInstance.snakeInstance.snakeObj.Move(newPos)
			gameInstance.snakeInstance.body = append(gameInstance.snakeInstance.body[1:1], newPos)
		case "right":
			newPos := fyne.NewPos(gameInstance.snakeInstance.snakeObj.Position().X+singlePix, gameInstance.snakeInstance.snakeObj.Position().Y)
			gameInstance.snakeInstance.snakeObj.Move(newPos)
			gameInstance.snakeInstance.body = append(gameInstance.snakeInstance.body[1:1], newPos)
		}

		// Snake dies on touching the game window.
		if checkIfWindowHit() { gameOver() }

		// Score goes up by one when snake head touches it.
		if checkIfPelletHit() {
			gameInstance.pellet = foodPellet()
			gameInstance.window.SetContent(container.NewWithoutLayout(&gameInstance.snakeInstance.snakeObj, gameInstance.pellet))

			gameInstance.score++

			fmt.Printf("gameInstance.score: %v\n", gameInstance.score)
		}

		gameInstance.window.Canvas().Refresh(&gameInstance.snakeInstance.snakeObj)
	}
}

func randomPositionInGameWindow() fyne.Position {
	var i fyne.Position
	xPos := randomNumber(22)
	yPos := randomNumber(22)
	i = fyne.NewPos(float32(xPos), float32(yPos))
	for i == gameInstance.snakeInstance.body[0] {
		xPos = randomNumber(22)
		yPos = randomNumber(22)
		i = fyne.NewPos(float32(xPos), float32(yPos))
	}
	return i
}

func randomNumber(limit int) int {
	var i int
	i = rand.Intn(limit)
	for i <= 1 {
		i = rand.Intn(limit)
	}
	return i * singlePix
}

func checkIfWindowHit() bool {
	return (
        (gameInstance.snakeInstance.snakeObj.Position().Y == finalSpaceHeight) ||
        (gameInstance.snakeInstance.snakeObj.Position().X == finalSpaceWidth) ||
        (gameInstance.snakeInstance.snakeObj.Position().X < 0) ||
        (gameInstance.snakeInstance.snakeObj.Position().Y < 0))
}

func checkIfPelletHit() bool {
	return gameInstance.snakeInstance.snakeObj.Position() == gameInstance.pellet.Position()
}

func gameOver() {
	fmt.Println("Game over!!")
	os.Exit(0)
}
