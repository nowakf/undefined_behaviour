package ui

import (
	c "cthu3/common"
	"github.com/faiface/pixel/pixelgl"
	"math"
)

//ui manages the stuff that's in view
type ui struct {
	h, w  int
	win   *pixelgl.Window
	x, y  int
	view  []c.Cell
	state State
}

//creates a new UI, returns a pointer
func NewUI(h, w int, win *pixelgl.Window) *ui {

	u := new(ui)
	u.h, u.w = h, w
	u.view = make([]c.Cell, h*w)
	u.state = NewSetup(h, w)
	u.win = win
	return u

}

func (u *ui) Draw() []c.Cell {

	diff := u.state.Draw(0, 0) //you should pass this the scrolling offset

	u.view = make([]c.Cell, 0)

	for _, cell := range diff {
		if cell.X < u.w && cell.Y < u.h {
			u.view = append(u.view, cell)
		}
	}

	return u.view

}

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

	mouseX, mouseY := u.win.MousePosition().XY()

	boundsH, boundsW := u.win.Bounds().H(), u.win.Bounds().W()

	mx, my := u.mousepos(mouseX, mouseY, boundsH, boundsW)

	x, y := u.floatToCellCoord(mx, my, u.w, u.h)

	mousePressed := u.win.Pressed(pixelgl.MouseButton1)
	mouseReleased := u.win.JustReleased(pixelgl.MouseButton1)

	switch {
	case u.x == x && u.y == y:
		if mousePressed || mouseReleased {
			function := u.state.OnMouse(x, y, mousePressed, mouseReleased)
			println(function())
			return true
		} else {
			return false
		}
	case mousePressed || mouseReleased:
		function := u.state.OnMouse(x, y, mousePressed, mouseReleased)
		u.state.OnMouse(u.x, u.y, false, false)
		u.x, u.y = x, y
		println(function())
		return true
	default:
		u.x, u.y = x, y
		function := u.state.OnMouse(x, y, mousePressed, mouseReleased)
		println(function())
		return true
	}

}
