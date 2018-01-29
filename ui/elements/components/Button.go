package components

type Button struct {
	action func() string
	*Rect
	MouseOver bool
}

func NewButton(action func() string, hitbox *Rect) *Button {
	b := new(Button)
	b.action = action
	b.Rect = hitbox
	return b
}

func (b Button) OnMouse(x, y int, clicked bool) func() string {
	if x <= b.W() && y <= b.H() {
		if clicked {
			return b.action
		} else {
			b.MouseOver = true
			return func() string { return "mouse over a button!" }
		}
	} else {
		return func() string { return "mouse over the buttons cell, but not on the button itself" }
	}
}
