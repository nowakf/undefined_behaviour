package ui

import (
	el "ub/ui/elements"
)

type Setup struct {
	*el.Table
}

func NewSetup(h, w int) *Setup {
	s := new(Setup)
	s.Table = el.NewTable(h, w)
	s.questionaire()
	return s
}

func (s *Setup) questionaire() {

	s.WriteToCell(0, 0, el.NewSpacer(1, s.W()/4))
	s.WriteToCell(5, 0, el.NewSpacer(1, s.W()/4))

	column1 := []el.UiElement{
		el.NewTextButton(s.W()/2/3, "A button!", func() string {
			return "a top left button!"
		}),
		el.NewTextButton(s.W()/2/3, "A button!", func() string {
			return "a middle left button!"
		}),
		el.NewTextButton(s.W()/2/3, "A button!", func() string {
			return "a bottom left button!"
		}),
	}

	for y, element := range column1 {
		s.WriteToCell(1, y+2, element)
	}

	column2 := []el.UiElement{
		el.NewTextButton(s.W()/2/3, "A button!", func() string {
			return "a top middle button!"
		}),
		el.NewTextButton(s.W()/2/3, "A button!", func() string {
			return "a middle middle button!"
		}),
		el.NewTextButton(s.W()/2/3, "Button four", func() string {
			return "a bottom middle button!"
		}),
	}
	for y, element := range column2 {
		s.WriteToCell(2, y+2, element)
	}
	column3 := []el.UiElement{
		el.NewTextButton(s.W()/2/3, "A button!", func() string {
			return "a top right button!"
		}),
		el.NewTextButton(s.W()/2/3, "A button!", func() string {
			return "a middle right button!"
		}),
		el.NewTextButton(s.W()/2/3, "A button!", func() string {
			return "a bottom right button!"
		}),
	}

	for y, element := range column3 {
		s.WriteToCell(3, y+2, element)
	}

}
