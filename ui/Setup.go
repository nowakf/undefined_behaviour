package ui

import (
	"ub/events"
	el "ub/ui/elements"
)

type setup struct {
	*el.Table
	player *events.Actor
}

func NewSetup(h, w int, player *events.Actor) *setup {
	s := new(setup)
	s.Table = s.questionaire(h, w)
	s.player = player
	return s
}

func (s *setup) HasNew() bool {
	return false
}

//the purpose of this is to weed out input events that don't belong in the top level
//func (s *setup) OnMouse(x, y int, pressed, released bool) func() string {

//return func() string { return "foo" }
//}

//makes a starting questionaire
func (s *setup) questionaire(h, w int) *el.Table {
	tab := el.NewTable(h, w)

	tab.WriteToCell(0, 0, el.NewSpacer(h, w/4))

	question1 := el.NewTextbox(w/2, "What is your name?")
	tab.WriteToCell(1, 0, question1)

	options1 := el.NewTable(w/2, h/4)
	options1.WriteToCell(0, 0, el.NewTextButton(w/2/3, "frederick", func() string {
		return "frederick"
	}))
	options1.WriteToCell(1, 0, el.NewTextButton(w/2/3, "james", func() string {
		return "james"
	}))
	options1.WriteToCell(2, 0, el.NewTextButton(w/2/3, "ben", func() string {
		return "ben"
	}))
	tab.WriteToCell(1, 1, options1)

	return tab

}
