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

	tab.WriteToCell(0, 0, el.NewSpacer(1, w/4))
	tab.WriteToCell(5, 0, el.NewSpacer(1, w/4))

	column1 := []el.UiElement{
		el.NewTextButton(w/2/3, "A button!", func() string {
			return "a top left button!"
		}),
		el.NewTextButton(w/2/3, "A button!", func() string {
			return "a middle left button!"
		}),
		el.NewTextButton(w/2/3, "A button!", func() string {
			return "a bottom left button!"
		}),
	}

	for y, element := range column1 {
		tab.WriteToCell(1, y+2, element)
	}

	column2 := []el.UiElement{
		el.NewTextButton(w/2/3, "A button!", func() string {
			return "a top middle button!"
		}),
		el.NewTextButton(w/2/3, "A button!", func() string {
			return "a middle middle button!"
		}),
		el.NewTextButton(w/2/3, "Button four", func() string {
			return "a bottom middle button!"
		}),
	}
	for y, element := range column2 {
		tab.WriteToCell(2, y+2, element)
	}
	column3 := []el.UiElement{
		el.NewTextButton(w/2/3, "A button!", func() string {
			return "a top right button!"
		}),
		el.NewTextButton(w/2/3, "A button!", func() string {
			return "a middle right button!"
		}),
		el.NewTextButton(w/2/3, "A button!", func() string {
			return "a bottom right button!"
		}),
	}

	for y, element := range column3 {
		tab.WriteToCell(3, y+2, element)
	}
	return tab
}
