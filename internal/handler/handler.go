package handler

import (
	"context"
	"math/rand"
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/PhilomathesInc/snake-game/internal/models/game"
)

func Start(ctx context.Context) {
	rand.Seed(time.Now().UnixNano())

	app := app.New()
	game.Start(ctx, app)
}
