package ui

import el "cthu3/ui/elements"

type Email struct {
	*el.Table
}

func NewEmail(h, w int) *Email {
	e := new(Email)
	e.Table = el.NewTable(h, w)
	return e
}

func (e *Email) AddMail() {
}
