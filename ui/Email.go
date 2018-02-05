package ui

import el "ub/ui/elements"

type email struct {
	*el.Table
}

func NewEmail(h, w int) *email {
	e := new(email)
	e.Table = el.NewTable(h, w)
	return e
}

func (e *email) AddMail() {
}
