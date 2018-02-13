package ui

import (
	//	"ub/events"
	el "ub/ui/elements"
)

type emailViewer struct {
	*el.Table
	mailBox email
	mail    []email
}

type message struct {
	subject string
	sender  string
	content string
	options []string
}

func NewEmailViewer(h, w int, mail email) *emailViewer {
	e := new(emailViewer)
	e.Table = el.NewTable(h, w)
	e.mailBox = mail
	return e
}

func (e *emailViewer) HasNew() bool {
	if e.mailBox != nil {
		e.AddMail(e.mailBox)
		e.mailBox = nil
		return true
	} else {
		return false
	}
}

func (e *emailViewer) AddMail(email) {
}
