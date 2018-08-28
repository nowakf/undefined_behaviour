package maze

import (
	"golang.org/x/image/colornames"
	"image/color"
)

const (
	Fog = iota
)

type cloud struct {
	name  string
	color color.RGBA
}

var Clouds = [...]cloud{
	Fog: cloud{
		name:  "Fog",
		color: colornames.Whitesmoke,
	},
}
