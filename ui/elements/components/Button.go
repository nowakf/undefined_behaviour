package components

import c "cthu3/common"

type Button struct {
	action func()
	*Rect
	mouseOver bool
}

func NewButton(action func(), hitbox *Rect) *Button {
	b := new(Button)
	b.action = action
	b.Rect = hitbox
	return b
}

func (b Button) OnMouse(x, y int, clicked bool) bool {
	if x < b.W() && y < b.H() {
		if clicked {
			b.action()
			return true
		} else {
			b.mouseOver = true
			return true
		}
	} else {
		return false
	}
}

func (b Button) Draw(input []c.Cell) []c.Cell {
	if b.mouseOver {
		for _, cell := range input {
			cell.Content = 'O'
		}
	}
	return input
}
