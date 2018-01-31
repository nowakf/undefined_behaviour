package elements

import (
	"time"
	c "ub/common"
)

type button struct {
	action func() string
	*rect
	MouseOver bool
	Clicked   bool
}

func newButton(action func() string, hitbox *rect) *button {
	b := new(button)
	b.action = action
	b.rect = hitbox
	return b
}

func (b *button) outline(x, y int) []c.Cell {
	cells := make([]c.Cell, 0)
	for i := x; i < x+b.W(); i++ {
		overline := c.Cell{
			X:          i,
			Y:          y,
			Letter:     '‾',
			Foreground: c.Grey,
			Background: c.White}
		underline := c.Cell{
			X:          i,
			Y:          y,
			Letter:     '_',
			Foreground: c.Grey,
			Background: c.White}
		cells = append(cells, underline, overline)
	}
	for j := y; j < y+b.H(); j++ {
		leftline := c.Cell{
			X:          x - 1,
			Y:          j,
			Letter:     '|',
			Foreground: c.Grey,
			Background: c.White}
		rightline := c.Cell{
			X:          x + b.W(),
			Y:          j,
			Letter:     '|',
			Foreground: c.Grey,
			Background: c.White}
		cells = append(cells, leftline, rightline)
	}
	b.MouseOver = false
	return cells
}

func (b *button) Draw(x, y int) []c.Cell {
	switch {
	case b.MouseOver:
		return b.outline(x, y)
	case b.Clicked:
		return b.fill(x, y)
	default:
		return []c.Cell{c.Cell{
			X:          0,
			Y:          0,
			Letter:     ' ',
			Foreground: c.White,
			Background: c.Grey,
		}}
	}
}

func (b *button) OnMouse(x, y int, pressed bool, released bool) func() string {
	switch {
	case pressed:
		b.Clicked = true
		return func() string { return "mouse was on this button, but it released elsewhere" }
	case released:
		b.Clicked = false
		return b.action
	default:
		b.Clicked = false
		b.MouseOver = true
		return func() string { return "mouse over a button!" }
	}
}

func (b *button) fill(xoffset, yoffset int) []c.Cell {
	cells := make([]c.Cell, b.W()*b.H())
	length := len(cells)
	for i := 0; i < length; i++ {
		x := i % b.W()
		y := i / b.W()
		cell := c.Cell{
			X:          x + xoffset,
			Y:          y + yoffset,
			Letter:     '█',
			Foreground: c.Grey,
			Background: c.White,
		}
		cells[i] = cell
	}
	return cells
}

func (b *button) click() {
	b.Clicked = true
	time.Sleep(time.Millisecond * time.Duration(100))
	b.Clicked = false

}
