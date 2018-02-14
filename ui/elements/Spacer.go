package elements

import (
	"fmt"
	c "ub/common"
)

type Spacer struct {
	*rect
}

func NewSpacer(height, width int) *Spacer {
	s := new(Spacer)
	s.rect = newrect(height, width)
	return s

}

func (s *Spacer) GetLast(x, y int) UiElement {
	return s
}
func (s *Spacer) Draw(x int, y int) []c.Cell {
	cell := c.Cell{
		X:          x,
		Y:          y,
		Letter:     ' ',
		Foreground: c.Grey,
		Background: c.White,
	}
	return []c.Cell{cell}
}

func (s *Spacer) Identify() string {
	return fmt.Sprintf("spacer, origin unknown, dimensions (%v, %v)", s.H(), s.W())
}

func (s *Spacer) OnMouse(x int, y int, pressed bool, released bool) func() string {
	return func() string { return fmt.Sprintf("spacer at %v,%v", x, y) }
}
