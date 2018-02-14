package elements

import (
	"fmt"
	c "ub/common"
)

type TextButton struct {
	*button
	*rect
	text *text
}

func NewTextButton(width int, content string, action func() string) *TextButton {
	b := new(TextButton)
	b.rect = newrect(1, width)
	b.text = newtitletext(width, content)
	b.button = newButton(action, b.rect)
	return b
}

func (t *TextButton) Draw(x, y int) []c.Cell {
	cells := t.text.Draw(x, y)
	switch t.mode {
	case hover:
		cells = append(cells, t.Light(x, y)...)
	case clicked:
		cells = t.Light(x, y)
	default:
		//do nothing
	}
	return cells
}
func (t *TextButton) Identify() string {
	return fmt.Sprintf("a textbutton, origin unkown, dimensions %v, %v", t.H(), t.W())
}

func (t *TextButton) GetLast(x, y int) UiElement {
	return t
}
