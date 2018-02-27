package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	c "ub/common"
)

func newWindow() *pixelgl.Window {
	winConfig := pixelgl.WindowConfig{
		Title:     "X",
		Bounds:    pixel.R(0, 0, 1024, 768),
		Resizable: true,
		//		Monitor: pixelgl.PrimaryMonitor(),
	}

	win, err := pixelgl.NewWindow(winConfig)
	c.Check(err)
	win.SetSmooth(false)

	return win
}
