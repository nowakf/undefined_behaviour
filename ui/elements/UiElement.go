package elements

import (
	c "ub/common"
)

type UiElement interface {
	Rect
	Draw(x, y int) []c.Cell
	GetLast(x, y int) UiElement
	Identify() string
}

type Rect interface {
	GetRatio() (float64, float64)
	Resize(h int, w int)
	H() int
	W() int
}
