package ui

import (
	_ "strings"
	"ub/events"
	el "ub/ui/elements"
)

type newsViewer struct {
	*linker
	*viewer
	*notifier
	player *events.Player
	name
}

func NewNewsViewer(v *viewer, l *linker, p *events.Player) *newsViewer {
	a := new(newsViewer)
	a.viewer = v
	a.player = p
	a.linker = l
	a.notifier = NewNotifier(a.player.News)
	a.name = "news"
	return a

}

func (n *newsViewer) Start() {

	n.Table.WriteToCell(0, 0, el.NewSpacer(n.Node, 1, n.W()/5))
	a := el.NewTextbox(n.Node, 1, n.W()/3, "This is the newsviewer.")
	n.Table.WriteToCell(1, 1, a)
	b := el.NewTextButton(n.Node, 1, n.W()/3, "We're done here.", func() string {
		n.Next(s_menu)
		return "fabulous"
	})
	n.Table.WriteToCell(1, 2, b)
}

func (n *newsViewer) Update() {}

func (n *newsViewer) Exit() {}
