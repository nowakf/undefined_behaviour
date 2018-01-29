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

func NewTextButton(width int, content string, action func()) *TextButton {
	b := new(TextButton)
	b.Rect = cp.NewRect(1, width)
	b.text = cp.NewTitleText(width, content)
	b.butt = cp.NewButton(action, b.Rect)
	return b
}

func (t *TextButton) Draw(x, y int) []c.Cell {
	return t.butt.Draw(t.text.Draw(x, y))
}

func (t *TextButton) OnMouse(x int, y int, clicked bool) bool {
	return t.butt.OnMouse(x, y, clicked)
}
