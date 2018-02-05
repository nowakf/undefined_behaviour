package main

import (
	//c "ub/common"
	"github.com/faiface/pixel/pixelgl"
	ev "ub/events"
	ui "ub/ui"
)

func run() {

	win := newWindow() //makes a pixelgl window

	data := newData() //is a collection of various data methods

	ren := newRender(win, data) //initializes the ui renderer

	uh, uw := ren.Stats() //gets the height/width

	config := ev.NewWorldConfig() //gens a default world config

	w := ev.NewWorld(config) //generates a world using the config

	ev := ev.NewEventSystem(w) //starts an event system

	u := ui.NewUI(uh, uw, win, ev) //makes a new ui

	check := resized()
	for !win.Closed() {

		if check(win) { //if the window is resized, redraw
			ren = newRender(win, data)
			uh, uw = ren.Stats()
			u = ui.NewUI(uh, uw, win, ev)
		}

		if u.Input() { //only run the main loop if there has been valid input
			cells := u.Draw()
			ren.update(cells)
		}
		//refactor to be in slower tick-rate loop:
		ev.Tick()
		//so it will be events.Poll here?

		win.Update()

	}
}

//resized() returns a function that can be called to check if the window has been resized
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

//for graphics related reasons, pixelgl must be in the main block
func main() {
	pixelgl.Run(run)
}
