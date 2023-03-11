# SNAKE GAME

1. Game Window
    - [x] 840X840
    - [x] The game pixel will be 40 X 40; Which will make the game window have 21 X 21 game pixels.
2. Snake
    - [x] Constant length of 1
    - [x] Spawns in the center.
    - [x] Moves using the WASD keys.
    - [x] Snake dies on touching the game window.
    - [x] Increase length of snake on food consumption
    - [x] Snake dies by touching itself.
3. Food Pellets
    - [x] 1 unit dimension
    - [x] Spawns randomly within game window.
    - [x] Doesn't spawn on cells occupied by the snake
    - [x] Once "consumed" by the snake, new food pellet needs to be spawned
4. Score counter
    - Go up by one when snake head touches it.
        - [x] Score counting
        - [x] Score display widget

Snake:
- func New(length, width int) Snake {}
    - Sets length to 1.
    - Sets the position to X and Y.
    
- func Move() {}
    - This method just updates the position of all the snake nodes based on the direction set on the head node.
    - Does not update the canvas.

- func HeadPosition() fyne.Position{}
    - Can be used in multiple ways to detect food pellet and to detect collision with game window.

- func Grow() {}
    - This will add a snakeNode to the type Snake.

FoodPellet:
- func New(length, width int) FoodPellet{}
    - Creates a circle canvas object and sets the position to X,Y.