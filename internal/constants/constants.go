package constants

import (
	"image/color"

	"fyne.io/fyne/v2"
)

const (
	FinalSpaceWidth  = 840
	FinalSpaceHeight = 840
	SinglePix        = 40
	PixelCountLimit  = 21
)

var (
	CenterGamePixel = fyne.NewPos((FinalSpaceWidth-SinglePix)/2, (FinalSpaceHeight-SinglePix)/2)
	Green           = color.NRGBA{R: 0, G: 180, B: 0, A: 255}
)
