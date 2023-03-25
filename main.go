package main

import (
	"context"

	"github.com/PhilomathesInc/snake-game/internal/handler"
)

func main() {
	handler.Start(context.Background())
}
