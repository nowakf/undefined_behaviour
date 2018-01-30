package main

import (
	//c "cthu3/common"
	"cthu3/data"
	ev "cthu3/events"
	ui "cthu3/ui"
	"github.com/faiface/pixel/pixelgl"
)

func run() {

	win := newWindow()

	ren := newRender(win)

	uh, uw := ren.Stats()

	ui := ui.NewUI(uh, uw, win)

	data := data.NewData()

	config := ev.NewWorldConfig()

	w := ev.NewWorld(config, data)

	ev := ev.NewEventSystem(data, w)

	ev.Test()

	for !win.Closed() {

		if ui.Input() {
			cells := ui.Draw()
			ren.update(cells)
		}
		//refactor to be in slower tick-rate loop:
		//and this:
		ev.Tick()

		win.Update()

	}
}

func main() {
	pixelgl.Run(run)
}
