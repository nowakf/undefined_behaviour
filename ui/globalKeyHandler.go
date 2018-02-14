package ui

import (
	"github.com/faiface/pixel/pixelgl"
	el "ub/ui/elements"
)

type globalKeyHandler struct {
	*pixelgl.Window
}

func (g *globalKeyHandler) Escape() el.KeyCatcher {
	//top level, so no parent
	return g
}
