package elements

import (
	c "ub/common"
)

type UiElement interface {
	W() int
	H() int
	SetW(width int)
	SetH(height int)
	Draw(x, y int) []c.Cell
	GetLast(x, y int) UiElement
	Identify() string
}
