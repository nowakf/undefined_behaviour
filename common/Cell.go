package common

import (
	"github.com/faiface/pixel"
)

type Cell struct {
	X, Y       int
	Letter     rune
	Foreground pixel.RGBA
	Background pixel.RGBA
}

var (
	LightGrey pixel.RGBA = pixel.RGB(192.0/255.0, 197.0/255.0, 206.0/255.0)
	DarkGrey             = pixel.RGB(58.0/255.0, 58.0/255.0, 58.0/255.0)
	White                = pixel.RGB(1.0, 1.0, 1.0)
	Red                  = pixel.RGB(1.0, 0.0, 0.0)
	Black                = pixel.RGB(0.0, 0.0, 0.0)
	Blank                = pixel.RGBA{R: 0.0, G: 0.0, B: 0.0, A: 0.0}
)

var Colors = []pixel.RGBA{
	DarkGrey,
	LightGrey,
	White,
	Red,
	Black,
	Blank,
}
