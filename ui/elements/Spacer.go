package elements

import (
	c "cthu3/common"
	cp "cthu3/ui/elements/components"
)

type Spacer struct {
	*cp.Rect
}

func NewSpacer(height_width ...int) *Spacer {
	if len(height_width) > 2 {
		panic("2d only, please.")
	}
	s := new(Spacer)
	s.Rect = cp.NewRect(height_width[0], height_width[1])
	return s

}

func (s *Spacer) Draw(x int, y int) []c.Cell {
	cell := c.Cell{X: x, Y: y, Content: ' '}
	return []c.Cell{cell}
}

func (s *Spacer) OnMouse(x int, y int, clicked bool) bool {
	return false
}
