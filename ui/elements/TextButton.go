package elements

import (
	c "ub/common"
)

type TextButton struct {
	butt *button
	*rect
	text *text
}

func NewTextButton(width int, content string, action func() string) *TextButton {
	b := new(TextButton)
	b.rect = newrect(1, width)
	b.text = newtitletext(width, content)
	b.butt = newButton(action, b.rect)
	return b
}

func (t *TextButton) Draw(x, y int) []c.Cell {
	cells := t.text.Draw(x, y)
	switch {
	case t.butt.MouseOver:
		cells = append(cells, t.butt.Draw(x, y)...)
	case t.butt.Clicked:
		cells = t.butt.Draw(x, y)
	}
	return cells
}

func (t *TextButton) OnMouse(x int, y int, pressed bool, released bool) func() string {
	return t.butt.OnMouse(x, y, pressed, released)
}
