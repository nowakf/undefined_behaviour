package elements

import (
	c "ub/common"
)

type button struct {
	action func() string
	*rect
	mode buttonMode
}

type buttonMode int

const (
	none buttonMode = iota
	hover
	clicked
)

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
	return cells
}

func (b *button) Light(x, y int) []c.Cell {
	switch b.mode {
	case hover:
		return b.outline(x, y)
	case clicked:
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

func (b *button) Flush() {
	b.mode = none
}

func (b *button) OnMouse(click bool) {
	if click {
		b.mode = clicked
	} else {
		b.mode = hover
	}
}

func (b *button) Do() {
	b.action()
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
