package elements

import (
	c "ub/common"
)

type textbox struct {
	*container
	i *text
}

//NewTextbox returns a textbox, with the width as fraction, with the int
//refering to the number under the line
func NewTextbox(parent *Node, h, w int, content string) *textbox {
	t := new(textbox)
	t.container = NewContainer(t, parent, h, w)
	t.i = newbodytext(content, t.container)
	t.foreground = c.LightGrey
	t.background = c.Blank
	return t
}

func (t *textbox) Identify() string {
	return t.i.Content()
}

func (t *textbox) Draw(x, y int) []c.Cell {
	return t.i.Draw(x+1, y) //offset for a border
}

func (t *textbox) GetLast(x, y int) UiElement {
	return t
}
