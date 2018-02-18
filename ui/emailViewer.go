package ui

import (
	"ub/events"
	el "ub/ui/elements"
)

type emailViewer struct {
	root   *el.Node
	player *events.Actor
	world  *events.World
	*el.Table
	mail []message
}

type email interface {
	GetTitle() string
	GetInstigator() *events.Actor
	GetBody() string
	GetOptions() []*events.Action
	GetDepth() int
}

type message struct {
	subject string
	sender  string
	content string
	options []*el.TextButton
}

func NewEmailViewer(h, w int, player *events.Actor) *emailViewer {
	e := new(emailViewer)
	e.root = el.Root()
	e.player = player
	e.Table = el.NewTable(e.root, h, w)
	e.initFrames()
	return e
}

func (e *emailViewer) initFrames() {
}

func (e *emailViewer) HasNew() bool {
	mail, number := e.player.History().NewMail()()
	if number > 0 {
		e.AddMail(mail)
		return true
	} else {
		return false
	}
}

func (e *emailViewer) AddMail(n email) {
	m := message{
		subject: e.ornamentTitle(n),
		sender:  e.toEmail(n),
		content: n.GetBody(),
		options: e.parseOpts(n),
	}
	e.mail = append(e.mail, m)

}

func (e *emailViewer) ornamentTitle(a email) string {
	subject := ""
	for i := 0; i <= a.GetDepth(); i++ {
		subject += "Re:"
	}
	return subject + a.GetTitle()
}
func (e *emailViewer) toEmail(a email) string {
	return a.GetInstigator().Name + "@" + a.GetInstigator().Org.EmailRoot
}
func (e *emailViewer) addTitle(email) {
}
func (e *emailViewer) Send(sent string) {
}

func (e *emailViewer) parseOpts(a email) []*el.TextButton {
	buttons := make([]*el.TextButton, 0)
	h := 1
	w := 15
	for _, c := range a.GetOptions() {
		b := el.NewTextButton(e.root, h, w, c.Title, func() string {
			c.Apply(e.player, a.GetInstigator(), e.world)
			for _, button := range buttons {
				button.Deactivate(false)
			}

			return "applied" + c.Title
		})
		buttons = append(buttons, b)
	}
	return buttons

}
