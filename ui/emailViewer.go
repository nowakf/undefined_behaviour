package ui

import (
	"ub/events"
	el "ub/ui/elements"
)

type emailViewer struct {
	*viewer
	*linker
	*notifier
	player *events.Player
	name
	titles   *el.Table
	body     *el.Table
	contacts []contact
	mail     map[*message]bool
}

func NewEmailViewer(v *viewer, l *linker, p *events.Player) *emailViewer {
	e := new(emailViewer)
	e.viewer = v
	e.player = p
	e.linker = l
	e.notifier = NewNotifier(e.player.Mail)
	return e
}

func (e *emailViewer) Start() {

	e.titles = el.NewTable(e.Node, e.H()-5, e.W()/5)

	e.Table.WriteToCell(0, 1, e.titles)

	e.body = el.NewTable(e.Node, e.H()-5, e.W()/5*4)
	e.Table.WriteToCell(1, 1, e.body)

	a := el.NewTextbox(e.Table.Node, 1, e.W()/5, "This is the emailviewer.")
	e.Table.WriteToCell(0, 2, a)

	b := el.NewTextbox(e.Table.Node, 1, e.W()/5, "here's some text, just to be clear")
	e.Table.WriteToCell(1, 2, b)

	c := el.NewTextButton(e.Table.Node, 1, e.W()/5, "We're done here.", func() string {

		println("you're certainly clicking it")

		e.Next(s_menu)

		return "fabulous"
	})

	e.Table.WriteToCell(0, 3, c)

	dummy := make(events.Options, 0)
	dummy.Insert("one", "one")
	dummy.Insert("two", "one")
	dummy.Insert("three", "one")
	dummy.Insert("four", "one")
	for i, button := range e.parseOpts(&dummy) {
		e.titles.WriteToCell(0, i, button)
	}
}
func (e *emailViewer) Update() {
	select {
	case email := <-e.Pipe():
		e.AddMail(email)
	default:
		//do nothing
	}
	i := 0
	for message, isRead := range e.mail {
		i++
		b := el.NewTextButton(e.Node, 1, e.W()/5, message.subject, func() string {
			e.body.WriteToCell(0, message.depth, el.NewTextbox(e.body.Node, 1, e.body.W(), message.content))
			isRead = true
			return "x"
		})
		e.titles.WriteToCell(0, i, b)
	}
}

func (e *emailViewer) Exit() {
}

func (e *emailViewer) AddMail(n *events.Event) {
	m := message{
		subject: e.ornamentTitle(n.Title(), n.Depth),
		sender:  n.Origin().Name() + "@" + n.Origin().Org(),
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
func (e *emailViewer) Send(sent string) {
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

type contact interface {
	Name() string
	Org() string
}

type message struct {
	subject string
	sender  string
	content string
	depth   int
	options []*el.TextButton
}
