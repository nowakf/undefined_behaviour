package ui

import (
	"github.com/faiface/pixel/pixelgl"
	el "ub/ui/elements"
)

type keyboard struct {
	context el.KeyCatcher
	win     *pixelgl.Window
}

func (k *keyboard) Event() bool {
	context, ok := k.context.DoIfValid(k.win.Typed(), k.win)
	if context != nil {
		k.context = context
		return ok
	} else {
		return ok
	}
}
