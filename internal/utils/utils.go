package utils

import (
	"math/rand"

	"github.com/PhilomathesInc/snake-game/internal/constants"
)

func RandomNumber(limit int) int {
	var i int
	i = rand.Intn(limit)
	for i <= 1 {
		i = rand.Intn(limit)
	}
	return i * constants.SinglePix
}
