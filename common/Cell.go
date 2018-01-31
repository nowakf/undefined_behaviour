package common

import "github.com/faiface/pixel"

type Cell struct {
	X, Y       int
	Letter     rune
	Foreground pixel.RGBA
	Background pixel.RGBA
}

var (
	Grey  pixel.RGBA = pixel.RGB(0.8, 0.8, 0.8)
	White            = pixel.RGB(1.0, 1.0, 1.0)
	Red              = pixel.RGB(1.0, 0.0, 0.0)
	Black            = pixel.RGB(0.0, 0.0, 0.0)
)

var Colors = []pixel.RGBA{
	Grey,
	White,
	Red,
	Black,
}
