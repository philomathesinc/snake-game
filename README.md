# SNAKE GAME

1. Game Window
    - [x] 840X840
    - [x] The game pixel will be 40 X 40; Which will make the game window have 21 X 21 game pixels.
2. Snake
    - [x] Constant length of 1
    - [x] Spawns in the center.
    - [x] Moves using the WASD keys.
    - [x] Snake dies on touching the game window.
3. Food Pellets
    - [x] 1 unit dimension
    - [x] Spawns randomly within game window.
    - [x] Doesn't spawn on cells occupied by the snake
    - [x] Once "consumed" by the snake, new food pellet needs to be spawned
4. Score counter
    - Go up by one when snake head touches it.
        - [x] Score counting
        - [ ] Score display widget
---

This is a simple implementation of the classic game "Snake" using the
Fyne GUI toolkit. The game window is set to a size of `840x840`, and the game
space is divided into `21x21` grid squares of `40x40` pixels each.

The game starts with a single block of `green` color representing the snake's
head, with its position set to the `center` of the game window. The snake's body
is represented by a series of `fyne.Position` values that correspond to the grid
squares it occupies. The direction in which the snake moves is determined by
the WASD keys, and the snake's movement speed is controlled by a `time.Sleep`
function call in the game loop.

A `white` circle represents the food pellet, which is placed at a random position
on the game board. When the snake's head collides with the food pellet, the
pellet disappears, and a new one is placed at a different random position, also the
score is incremented. If the snake's head collides with any of the walls, the
game is over.

The code uses Go's `fyne` package to create the graphical user interface and
`time` package to control the game's loop speed.

