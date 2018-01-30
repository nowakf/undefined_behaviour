package elements

import (
	c "cthu3/common"
	cp "cthu3/ui/elements/components"
)

type TextButton struct {
	butt *cp.Button
	*cp.Rect
	text *cp.Text
}

func NewTextButton(width int, content string, action func() string) *TextButton {
	b := new(TextButton)
	b.Rect = cp.NewRect(1, width)
	b.text = cp.NewTitleText(width, content)
	b.butt = cp.NewButton(action, b.Rect)
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
