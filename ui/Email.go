package ui

import (
	"ub/events"
	el "ub/ui/elements"
)

type emailViewer struct {
	*el.Table
	mailPipe chan (events.Event)
	mails    []message
}

type message struct {
	subject string
	sender  string
	content string
	options []string
}

func NewEmailViewer(h, w int, mailHook func(chan (events.Event))) *emailViewer {
	e := new(emailViewer)
	e.Table = el.NewTable(h, w)
	e.mailPipe = make(chan (events.Event))
	mailHook(e.mailPipe)
	return e
}

func (e *emailViewer) HasNew() bool {
	select {
	case hasNew := <-e.mailPipe:
		e.AddMail(hasNew)
		return true
	default:
		return false
	}
}

func (e *emailViewer) AddMail(m events.Event) {
}
