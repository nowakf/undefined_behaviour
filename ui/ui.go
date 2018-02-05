// ui manages user input, the abstract display of objects, and
// sends messages to the event system, when input changes game state.
package ui

import (
	"github.com/faiface/pixel/pixelgl"
	"math"
	c "ub/common"
	"ub/events"
	el "ub/ui/elements"
)

// a ui displays the stuff that's in view, manages changes between states,
// and sends game-state-altering input to the event system
type ui struct {
	h, w int

	win    *pixelgl.Window
	events *events.EventSystem

	x, y int

	view         []c.Cell
	currentState state
	states       []el.UiElement
}

//creates a new UI, returns a pointer
func NewUI(h, w int, win *pixelgl.Window, e *events.EventSystem) *ui {

	u := new(ui)
	u.events = e
	u.h, u.w = h, w
	u.view = make([]c.Cell, h*w)
	u.states = u.initStates(h, w)
	u.win = win
	return u
}

type state int

const (
	aSetup state = iota
	aEmail
)

var states = []state{
	aSetup,
	aEmail,
}

func (u *ui) initStates(h, w int) []el.UiElement {
	m := make([]el.UiElement, len(states))
	m[aSetup] = NewSetup(h, w)
	m[aEmail] = NewEmail(h, w)
	return m
}

func (u *ui) Draw() []c.Cell {

	diff := u.states[u.currentState].Draw(0, 0) //you should pass this the scrolling offset

	u.view = make([]c.Cell, 0)

	for _, cell := range diff {
		if cell.X < u.w && cell.Y < u.h {
			u.view = append(u.view, cell)
		}
	}

	return u.view

}

//this will check what input there is, then return true if it exists
func (u *ui) Input() bool {

	return u.mouse()

}

//gets the position from the window
func (u *ui) mousepos(mouseX, mouseY, boundsH, boundsW float64) (float64, float64) {

	y := mouseY / boundsH
	x := mouseX / boundsW

	return x, y
}

//converts the mouse position to a cell co-ordinate
func (u *ui) floatToCellCoord(fx, fy float64, width, height int) (int, int) {

	x := int(math.Floor(float64(width) * fx))
	y := height - int(math.Floor(float64(height)*fy))

	return x, y
}

//returns if the mouse is over something
func (u *ui) mouse() bool {

	//get the raw mouse position
	mouseX, mouseY := u.win.MousePosition().XY()

	//get the bounds of the screen
	boundsH, boundsW := u.win.Bounds().H(), u.win.Bounds().W()

	//get the relative mouse position
	mx, my := u.mousepos(mouseX, mouseY, boundsH, boundsW)

	//get the cell coordinate
	x, y := u.floatToCellCoord(mx, my, u.w, u.h)

	//get the mousepress events
	mousePressed := u.win.Pressed(pixelgl.MouseButton1)
	mouseReleased := u.win.JustReleased(pixelgl.MouseButton1)

	//apply:
	switch {
	//if mouse hasn't moved,
	case u.x == x && u.y == y:
		//and if it has been clicked, update
		if mousePressed || mouseReleased {
			function := u.states[u.currentState].OnMouse(x, y, mousePressed, mouseReleased)
			//return the function
			println(string(u.currentState))
			println(function())
			return true
		} else {
			//if it has neither been moved nor pressed, do nothing
			return false
		}
	//if the mouse is pressed and moved,
	case mousePressed || mouseReleased:
		function := u.states[u.currentState].OnMouse(x, y, mousePressed, mouseReleased)
		//return the function
		u.states[u.currentState].OnMouse(u.x, u.y, false, false)
		//update the previous position
		u.x, u.y = x, y
		println(function())
		return true
	//if the mouse is moved
	default:
		u.x, u.y = x, y
		//update the previous position
		function := u.states[u.currentState].OnMouse(x, y, mousePressed, mouseReleased)
		println(function())
		//tell the new button the mouse is here
		return true
	}

}
