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
	q := make([]el.UiElement, 4)
	println(s.w)
	q[0] = el.NewTextBox(s.w/2, "So many times, they went back and forth, combing the data for irregularities, shifting between one theorem and the next - ")

	column1 := []el.UiElement{
		el.NewTextButton(s.w/2/3, "A button!", func() string {
			return "a button!"
		}),
		el.NewTextButton(s.w/2/3, "A button!", func() string {
			return "a button!"
		}),
		el.NewTextButton(s.w/2/3, "A button!", func() string {
			return "a button!"
		}),
	}

	somebuttons := [][]el.UiElement{column1, column1, column1}

	q[1] = el.NewTable(3, s.w/2, somebuttons)

	q[2] = el.NewSpacer(3, s.w/2)
	q[3] = el.NewTextButton(s.w/2, "a new button", func() string {
		return "a button"
	})
	return q

}
