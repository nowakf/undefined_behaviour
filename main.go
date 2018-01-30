package main

import (
	//c "cthu3/common"
	ev "cthu3/events"
	ui "cthu3/ui"
	"github.com/faiface/pixel/pixelgl"
)

func run() {

	win := newWindow()

	ren := newRender(win)

	uh, uw := ren.Stats()

	config := ev.NewWorldConfig()

	w := ev.NewWorld(config)

	ev := ev.NewEventSystem(w)

	u := ui.NewUI(uh, uw, win, ev)

	check := resized()
	for !win.Closed() {
		if check(win) {
			ren = newRender(win)
			uh, uw = ren.Stats()
			u = ui.NewUI(uh, uw, win, ev)
		}

		if u.Input() {
			cells := u.Draw()
			ren.update(cells)
		}
		//refactor to be in slower tick-rate loop:
		ev.Tick()
		//so it will be events.Poll here?

		win.Update()

	}
}
func resized() func(win *pixelgl.Window) bool {

	h := 0.0
	w := 0.0

	return func(win *pixelgl.Window) bool {

		hi := win.Bounds().H()
		wi := win.Bounds().W()
		if h != hi || w != wi {
			h = hi
			w = wi
			return true
		} else {
			return false
		}
	}
}

func main() {
	pixelgl.Run(run)
}
