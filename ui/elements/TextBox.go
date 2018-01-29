package elements

import (
	c "cthu3/common"
	cp "cthu3/ui/elements/components"
)

type TextBox struct {
	*cp.Rect
	*cp.Text
}

func NewTextBox(width int, content string) *TextBox {
	t := new(TextBox)
	t.Text, _ = cp.NewBodyText(width-2, content) // to give it some padding
	t.Rect = cp.NewRect(t.Text.H()+1, width)
	return t
}

func (t *TextBox) Content() string {
	return t.Text.Content()
}

func (t *TextBox) Draw(x, y int) []c.Cell {
	return t.Text.Draw(x+1, y) //offset for a border
}

func (t *TextBox) OnMouse(x int, y int, clicked bool) bool {
	println(t.Content)
	return true
}
