package ui

import (
	"ub/events"
	el "ub/ui/elements"
)

type setup struct {
	root *el.Node
	*el.Table
	player *events.Actor
}

func NewSetup(h, w int, player *events.Actor) *setup {
	s := new(setup)
	s.root = el.Root()
	s.player = player
	s.Table = s.questionaire(h, w)
	return s
}

func (s *setup) HasNew() bool {
	return false
}

//makes a starting questionaire
func (s *setup) questionaire(h, w int) *el.Table {
	tab := el.NewTable(s.root, h, w)
	tab.WriteToCell(0, 0, el.NewSpacer(tab.Node, 1, w/4))
	a := el.NewTextbox(tab.Node, 1, w/3, "Sometimes, I question what we did, not on moral grounds - but on practical ones. There wasn't much to be said, after all was said and done.")
	tab.WriteToCell(1, 1, a)
	b := el.NewTextButton(tab.Node, 1, w/3, "A new button. Fabulous.", func() string { return "fabulous" })
	tab.WriteToCell(1, 3, b)
	println(b.W(), a.W(), "width of b, then a")
	return tab
}
