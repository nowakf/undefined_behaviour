package elements

import (
	"fmt"
	c "ub/common"
)

type Spacer struct {
	*container
}

func NewSpacer(parent *Node, h, w int) *Spacer {
	s := new(Spacer)
	s.container = NewContainer(s, parent, h, w)
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
		Foreground: c.Blank,
		Background: c.Blank,
	}
	return []c.Cell{cell}
}

func (s *Spacer) Identify() string {
	return fmt.Sprintf("spacer, origin unknown, dimensions (%v, %v)", s.H(), s.W())
}

func (s *Spacer) OnMouse(x int, y int, pressed bool, released bool) func() string {
	return func() string { return fmt.Sprintf("spacer at %v,%v", x, y) }
}
