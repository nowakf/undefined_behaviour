package elements

import (
	c "ub/common"
)

type button struct {
	action func() string
	*container
	mode   buttonMode
	active bool
}

type buttonMode int

const (
	none buttonMode = iota
	hover
	clicked
	wasSelected
)

func newButton(action func() string, hitbox *container) *button {
	b := new(button)
	b.action = action
	b.container = hitbox
	b.active = true
	return b
}

func (b *button) Light() {
	switch b.mode {
	case hover:
		b.foreground = c.White
		b.background = c.LightGrey
	case clicked:
		b.foreground = c.DarkGrey
		b.background = c.White
	case wasSelected:
		b.foreground = c.DarkGrey
		b.background = c.LightGrey
	case none:
		b.foreground = c.LightGrey
		b.background = c.Blank
	default:
		println("check the button logic")
	}
}

func (b *button) Deactivate(wasClicked bool) {
	b.active = false
	if wasClicked {
		b.mode = wasSelected
	}
}

func (b *button) Flush() {
	if b.active {
		b.mode = none
	}
	b.Light()
}

func (b *button) OnMouse(click bool) {
	if b.active {
		if click {
			b.mode = clicked
		} else {
			b.mode = hover
		}
	}
}

func (b *button) Do() {
	if b.active {
		b.action()
	}
}
