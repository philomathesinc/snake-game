# SNAKE GAME

1. Game Window
    - [ ] 840X840
    - [ ] The game pixel will be 40 X 40; Which will make the game window have 21 X 21 game pixels.
2. Snake
    - [ ] Constant length of 1
    - [ ] Spawns in the center.
    - [ ] Moves using the WASD keys.
    - [ ] Snake dies on touching the game window.
3. Food Pellets
    - [x] 1 unit dimension
    - [x] Spawns randomly within game window.
    - [ ] Doesn't spawn on cells occupied by the snake - Not possible till snake consumption of pellet is done.
    - [ ] Once "consumed" by the snake, new food pellet needs to be spawned
4. Score counter
    - Go up by one when snake head touches it.
---
Snake
    - Length is by default 2 units.
    - Snake increases in 1 unit length when it touches the food pellet.
    - Snake dies by touching itself.