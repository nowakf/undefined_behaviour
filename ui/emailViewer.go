package ui

import (
	"ub/events"
	el "ub/ui/elements"
)

type emailViewer struct {
	*viewer
	*linker
	*notifier
	name
	titles   *el.Table
	body     *el.Table
	contacts []contact
	mail     map[contact]message
}

type mailBox interface {
	Mail(chan email) func() int
}

func NewEmailViewer(v *viewer, l *linker, m mailBox) *emailViewer {
	e := new(emailViewer)
	e.viewer = v
	e.linker = l
	e.notifier = NewNotifier(m.Mail)
	return e
}

func (e *emailViewer) Start() {
}
func (e *emailViewer) Update() {
	select {
	case email := <-e.Pipe():
		e.AddMail(email)
	default:
		//do nothing
	}
}

func (e *emailViewer) Exit() {
}

type email interface {
	Title() string
	Body() string
	Do(index int)
}

func (e *emailViewer) AddMail(n email) {
	m := message{
		//sender:  n.Origin().Name() + "@" + n.Origin().Org(),
		content: n.Body(),
		options: e.parseOpts(n.Options),
	}
	e.mail[&m] = true

}
func (e *emailViewer) ornamentTitle(title string, depth int) string {
	subject := ""
	for i := 0; i <= depth; i++ {
		subject += "Re:"
	}
	return subject + title
}

func (e *emailViewer) addTitle() {
}

type contact interface {
	Name() string
	Org() string
}

func (e *emailViewer) Send(sent string) {
}

type option interface {
	Title() string
	Index() int
}

func (e *emailViewer) parseOpts(opts *events.Options) []*el.TextButton {
	buttons := make([]*el.TextButton, 0)
	h := 1
	w := e.titles.W()
	for _, option := range *opts {
		b := el.NewTextButton(e.root, h, w, option.Title, func() string {
			for _, button := range buttons {
				button.Deactivate()
			}
			option.Do(e.player)

			return "sent" + option.Title
		})
		buttons = append(buttons, b)
	}
	return buttons

}

type message struct {
	subject string
	sender  string
	content string
	depth   int
	options []*el.TextButton
}
