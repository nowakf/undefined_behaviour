package main

import (
	"github.com/nowakf/pixel/pixelgl"
	"time"
	"ub/data"
	e "ub/events"
	"ub/render"
	ui "ub/ui"
)

func run() {

	// load settings from a text file?
	// cool, but very unfriendly

	f := &data.FontLoader{} //is a collection of various data methods

	ren := render.New(f, 18)

	win := ren.Window

	uh, uw := ren.Stats() //gets the height/width

	save := data.NewData().LoadSave()

	config := e.NewWorldConfig() //gens a default world config

	//pconfig := e.PlayerConfig{} //and a default pconfig TODO

	w := e.NewWorld(config) //generates a world using the config
	if save != nil {
		w = e.LoadWorld(save)
	}

	n := names.New(500)
	for _, v := range n {
		println(v)
	}

	ev := e.NewEventSystem(w) //starts an event system

	u := ui.NewUI(uh, uw, win, ev)

	u.Start(uh, uw)
	//loading screen will go here:
	stack := u.Draw()
	ren.Update(stack)

	check := resized()

	fps := time.Tick(time.Second / 60)

	for !win.Closed() {

		<-fps

		if check(win) {
			println("resize!")
			uh, uw = ren.Stats()
			u.Resize(uh, uw)
		}

		if u.Event() {
			stack = u.Draw()
			ren.Update(stack)
			u.Update()
		}
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
