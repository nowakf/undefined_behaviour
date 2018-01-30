package ui

import (
	el "cthu3/ui/elements"
)

type Setup struct {
	*el.Table
	h, w int
}

func NewSetup(h, w int) *Setup {
	s := new(Setup)
	cont := make([][]el.UiElement, 3)
	s.h, s.w = h, w
	empty := []el.UiElement{el.NewSpacer(h, w/4)}
	cont[0] = empty
	cont[1] = s.questionaire()
	cont[2] = empty
	s.Table = el.NewTable(s.h, s.w, cont)
	return s
}

func (s *Setup) questionaire() []el.UiElement {
	q := make([]el.UiElement, 6)
	println(s.w)
	q[0] = el.NewSpacer(4, s.w/2)
	q[1] = el.NewTextBox(s.w/2, "So many times, they went back and forth, combing the data for irregularities, shifting between one theorem and the next - ")

	column1 := []el.UiElement{
		el.NewTextButton(s.w/2/5, "A button!", func() string {
			return "a top left button!"
		}),
		el.NewTextButton(s.w/2/5, "A button!", func() string {
			return "a middle left button!"
		}),
		el.NewTextButton(s.w/2/5, "A button!", func() string {
			return "a bottom left button!"
		}),
	}
	column2 := []el.UiElement{el.NewSpacer(3, s.w/2/6)}
	column3 := []el.UiElement{
		el.NewTextButton(s.w/2/5, "A button!", func() string {
			return "a top middle button!"
		}),
		el.NewTextButton(s.w/2/5, "A button!", func() string {
			return "a middle middle button!"
		}),
		el.NewTextButton(s.w/2/5, "A button!", func() string {
			return "a bottom middle button!"
		}),
	}
	column4 := []el.UiElement{el.NewSpacer(3, s.w/2/6)}
	column5 := []el.UiElement{
		el.NewTextButton(s.w/2/5, "A button!", func() string {
			return "a top right button!"
		}),
		el.NewTextButton(s.w/2/5, "A button!", func() string {
			return "a middle right button!"
		}),
		el.NewTextButton(s.w/2/5, "A button!", func() string {
			return "a bottom right button!"
		}),
	}
	somebuttons := [][]el.UiElement{column1, column2, column3, column4, column5}
	q[2] = el.NewSpacer(2, s.w/2)

	q[3] = el.NewTable(3, s.w/2, somebuttons)

	q[4] = el.NewSpacer(3, s.w/2)
	q[5] = el.NewTextButton(s.w/2, "a new button", func() string {
		return "a button"
	})
	return q

}
