package elements

import (
	c "ub/common"
)

type textbox struct {
	*rect
	*text
}

func NewTextbox(width int, content string) *textbox {
	t := new(textbox)
	t.text, _ = newbodytext(width-2, content) // to give it some padding
	t.rect = newrect(t.text.H()+1, width)
	return t
}

func (t *textbox) Identify() string {
	return t.text.Content()
}

func (t *textbox) Draw(x, y int) []c.Cell {
	return t.text.Draw(x+1, y) //offset for a border
}

func (t *textbox) GetLast(x, y int) UiElement {
	return t
}
