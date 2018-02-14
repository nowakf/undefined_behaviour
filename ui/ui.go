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

	global el.KeyCatcher

	win    *pixelgl.Window
	player *events.Actor

	last el.UiElement

	view         []c.Cell
	currentState state
	states       []state

	focused el.KeyCatcher
}

//creates a new UI, returns a pointer
func NewUI(h, w int, win *pixelgl.Window, e *events.Actor) *ui {

	u := new(ui)
	u.player = e
	u.h, u.w = h, w
	u.view = make([]c.Cell, h*w)
	u.focused = &globalKeyHandler{win}
	u.states = u.initStates(h, w)
	u.currentState = u.states[0] //take the first state, setup
	u.win = win
	return u
}

func (u *ui) initStates(h, w int) []state {
	states := []state{
		NewSetup(h, w, u.player),
		NewEmailViewer(h, w, u.player.MailBox),
	}
	return states
}

func (u *ui) HasNew() bool {

	n := false
	i := 0

	for n == false && i < len(u.states) {
		n = u.states[i].HasNew()
		i++
	}

	return n
}

func (u *ui) Draw() []c.Cell {

	diff := u.currentState.Draw(0, 0) //you should pass this the scrolling offset

	u.view = make([]c.Cell, 0)

	for _, cell := range diff {
		if cell.X < u.w && cell.Y < u.h {
			u.view = append(u.view, cell)
		}
	}

	return u.view

}

//this will check what input there is, then return true if it exists
func (u *ui) Event() bool {

	inputEvent := u.mouse()
	switch {
	case inputEvent:
		//event will probs be some kind of interface
		return true
	case u.HasNew():
		println("new something!")
		return true
	default:
		return false
	}

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

	object := u.currentState.GetLast(x, y)

	changed := object != u.last

	//catcher, ok := u.checkIfCatcher(object)

	return u.parseClick(object, changed, mousePressed, mouseReleased)

}
func (u *ui) checkIfCatcher(object el.UiElement) (el.KeyCatcher, bool) {
	var checked el.KeyCatcher
	checked, ok := object.(el.KeyCatcher)
	return checked, ok
}

func (u *ui) checkClickable(object el.UiElement) (el.Clickable, bool) {
	var checked el.Clickable
	checked, ok := object.(el.Clickable)
	return checked, ok
}

func (u *ui) parseClick(input el.UiElement, changed, mousePressed, mouseReleased bool) bool {

	object, isClickable := u.checkClickable(input)
	isNill := object == nil
	prev, prevIsClickable := u.checkClickable(u.last)

	switch {
	case changed && mousePressed:
		return false
	case mousePressed:
		if isClickable && !isNill {
			object.OnMouse(true)
		}
		return true
	case mouseReleased:
		if prevIsClickable {
			prev.Flush()
			prev.Do()
			return true //to hover over something else
		} else {
			return false
		}
	case changed:
		if isClickable && !isNill {
			object.OnMouse(false)
		}

		if prevIsClickable {
			prev.Flush()
		}
		if !isNill {
			u.last = object.(el.UiElement)
		}
		return true //to hover over something new
	default:
		return false
	}
}
