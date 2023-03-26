# SNAKE GAME

## How to Run
```
go mod tidy
FYNE_THEME=dark go run main.go
```

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