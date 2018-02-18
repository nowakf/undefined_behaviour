package elements

import c "ub/common"

type InputBox struct {
	*container
	field string
}

func NewInputBox(parent *Node, h, w int, field string) *InputBox {
	i := new(InputBox)
	i.container = NewContainer(i, parent, h, w)
	i.field = field
	return i
}

func (i *InputBox) Draw(x int, y int) []c.Cell {
	panic("not implemented")
}

func (i *InputBox) Do() {
}

func (i *InputBox) GetLast(x, y int) UiElement {
	return i
}

func (i *InputBox) OnMouse(x int, y int, pressed bool, released bool) func() string {
	panic("not implemented")
}

func (i *InputBox) Identify() string {
	panic("not implemented")
}
