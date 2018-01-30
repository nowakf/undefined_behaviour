package elements

import (
	c "cthu3/common"
	"fmt"
)

type Spacer struct {
	*rect
}

func NewSpacer(height, width int) *Spacer {
	s := new(Spacer)
	s.rect = newrect(height, width)
	return s

}

func (s *Spacer) Draw(x int, y int) []c.Cell {
	cell := c.Cell{X: x, Y: y, Content: ' '}
	return []c.Cell{cell}
}

func (s *Spacer) OnMouse(x int, y int, pressed bool, released bool) func() string {
	return func() string { return fmt.Sprintf("spacer at %v,%v", x, y) }
}
