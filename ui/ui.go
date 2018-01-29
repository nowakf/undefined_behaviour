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

	if u.mouse() {
		return true
	} else {
		return false
	}
}

func (u *ui) mouse() bool {

	px := u.win.MousePosition().X / u.win.Bounds().W()
	py := u.win.MousePosition().Y / u.win.Bounds().H()

	x := int(math.Floor(float64(u.w) * px))
	y := int(math.Floor(float64(u.h) * py))

	lMouseReleased := u.win.JustReleased(pixelgl.MouseButton1)

	if u.x == x && u.y == y && !lMouseReleased {
		return false
	} else {
		u.x, u.y = x, y
		function := u.state.OnMouse(x, y, lMouseReleased)
		println(function())
		return true
	}

}
