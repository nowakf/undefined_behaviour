package elements

import (
	"fmt"
	c "ub/common"
)

type TextButton struct {
	*container
	*button
	*text
}

func NewTextButton(parent *Node, h, w int, content string, action func() string) *TextButton {

	b := new(TextButton)
	b.container = NewContainer(b, parent, h, w)
	b.text = newtitletext(content, b.container)
	b.button = newButton(action, b.container)
	return b
}

func (t *TextButton) Draw(x, y int) []c.Cell {
	t.Light()
	cells := t.text.Draw(x, y)
	return cells
}
func (t *TextButton) Identify() string {
	return fmt.Sprintf("a textbutton, origin unkown, dimensions %v, %v", t.container.H(), t.container.W())
}

func (t *TextButton) GetLast(x, y int) UiElement {
	return t
}
