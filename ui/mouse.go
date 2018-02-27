package ui

import (
	"github.com/faiface/pixel/pixelgl"
	"math"
	el "ub/ui/elements"
)

type mouse struct {
	win  *pixelgl.Window
	last el.UiElement
}

//returns if the mouse is over something
func (m *mouse) Event(h, w int, current state) bool {

	//get the raw mouse position
	mouseX, mouseY := m.win.MousePosition().XY()

	//get the bounds of the screen
	boundsH, boundsW := m.win.Bounds().H(), m.win.Bounds().W()

	//get the relative mouse position
	mx, my := m.mousepos(mouseX, mouseY, boundsH, boundsW)

	//get the cell coordinate
	x, y := m.floatToCellCoord(mx, my, w, h)

	//get the mousepress events
	mousePressed := m.win.Pressed(pixelgl.MouseButton1)
	mouseReleased := m.win.JustReleased(pixelgl.MouseButton1)

	object := current.GetLast(x, y)

	changed := object != m.last

	return m.parseClick(object, changed, mousePressed, mouseReleased)

}

func (m *mouse) mousepos(mouseX, mouseY, boundsH, boundsW float64) (float64, float64) {

	y := mouseY / boundsH
	x := mouseX / boundsW

	return x, y
}

//converts the mouse position to a cell co-ordinate
func (m *mouse) floatToCellCoord(fx, fy float64, width, height int) (int, int) {

	x := int(math.Floor(float64(width) * fx))
	y := height - int(math.Floor(float64(height)*fy))

	return x, y
}

func (m *mouse) checkIfCatcher(object el.UiElement) (el.KeyCatcher, bool) {
	var checked el.KeyCatcher
	checked, ok := object.(el.KeyCatcher)
	return checked, ok
}

func (m *mouse) checkClickable(object el.UiElement) (el.Clickable, bool) {
	var checked el.Clickable
	checked, ok := object.(el.Clickable)
	return checked, ok
}

func (m *mouse) parseClick(input el.UiElement, changed, mousePressed, mouseReleased bool) bool {

	object, isClickable := m.checkClickable(input)

	isNill := object == nil

	prev, prevIsClickable := m.checkClickable(m.last)

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
			return true
		} else {
			return false
		}
	case changed:
		if prevIsClickable {
			prev.Flush()
		}
		if isClickable && !isNill {
			object.OnMouse(false)
		}

		m.last = input

		return true
	default:
		return false
	}
}
