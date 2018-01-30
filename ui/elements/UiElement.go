package elements

import (
	c "cthu3/common"
)

type UiElement interface {
	W() int
	H() int
	SetW(width int)
	SetH(height int)
	Draw(x, y int) []c.Cell
	OnMouse(x, y int, pressed bool, released bool) func() string
}
