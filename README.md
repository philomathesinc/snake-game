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
- func New() Snake {}
    - Sets length to 1.
    
- func (s *Snake) Move(fyne.Position) {}
    - This method just updates the position of all the snake nodes based on the direction set on the head node.
    - Does not update the canvas.

- func (s *Snake) HeadPosition() fyne.Position{}
    - Can be used in multiple ways to detect food pellet and to detect collision with game window.

- func (s *Snake) Grow() {}
    - This will add a snakeNode to the type Snake.

- func (s *Snake) Accelerate() {}
    - Increase the game speed, based on our existing design.
    - Ideally, speed acceleration should be on the snake object only.

Pellet:
- func New(pos fyne.Position) Pellet{}
    - Creates a circle canvas object and sets the position to pos.
- func Position() (TBD)
- ToDo : Once "consumed" by the snake, new food pellet needs to be spawned

Window:
- func (w *Window) randomPosition() fyne.Position {} 
    - return random position limited by length and width of window
- func (w *Window) Refresh() {}
    - g.window.Canvas().Refresh(&tmp.canvasObj)
- func (w *Window) UpdateContent() {}
    - 
- func Boundary() (TBD)

Score counter:
- func New() ScoreCounter{}
    - Initializes score to zero
    - Returns a text box with "Score: X"
- func Increment() {}
    - Increments the score by one.

Game:
- func bootstrap() []fyne.CanvasObject{}
    - Snake.New()
    - Pellet.New()
    - ScoreCounter.New()

- func Start() {}
    - bootstrap()
    - Start below go routines
        - Check food pellet consumption
            - Snake.Grow()
            - Window.UpdateContent(game, snake)
            - Snake.Accelerate()
            - ScoreCounter.Increment()
        - Check boundary hit
            - GameOver()
        - Check self hit on snake
            - GameOver()
        - Snake.Move()
        - GameWindow.SetContent()
