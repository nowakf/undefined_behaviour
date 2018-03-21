package ui

import (
	"ub/events"
	el "ub/ui/elements"
)

type setup struct {
	name
	*viewer
	*linker
	//*events.PlayerConfig
	//*events.WorldConfig
}

func NewSetup(v *viewer, l *linker, p *events.Player) *setup {
	s := new(setup)
	s.viewer = v
	s.name = "setup"
	s.linker = l
	return s
}

//makes a starting questionaire
func (s *setup) Start() {

	//stop time

	h, w := s.H(), s.W()
	s.Table = el.NewTable(s.root, h, w)
	s.Table.WriteToCell(0, 0, el.NewSpacer(s.Table.Node, 1, w/4))
	a := el.NewTextbox(s.Table.Node, 5, w/3, "Sometimes, I question what we did, not on moral grounds - but on practical ones. There wasn't much to be said, after all was said and done.")
	s.Table.WriteToCell(1, 1, a)
	b := el.NewTextButton(s.Table.Node, 1, w/3, "A new button. Fabulous.", func() string {
		return "fabulous"
	})
	s.Table.WriteToCell(1, 2, el.NewSpacer(s.Table.Node, 1, w/4))

	s.Table.WriteToCell(1, 3, b)

	c := el.NewTextButton(s.Table.Node, 1, w/3, "We're done here.", func() string {
		s.Next(s_email)
		return "fabulous"
	})
	s.Table.WriteToCell(1, 4, c)
}

func (s *setup) Update() {}

func (s *setup) Exit() {
	//start time
}
