package ui

import (
	"ub/events"
	el "ub/ui/elements"
)

type menu struct {
	name
	*viewer
	*linker
	ev *events.EventSystem
}

func NewMenu(v *viewer, l *linker, ev *events.EventSystem) *menu {
	m := new(menu)
	m.viewer = v
	m.name = "menu"
	m.linker = l
	m.ev = ev
	return m
}

func (m *menu) Start() {
	//stop time
	println("start happens")
	println(m.W(), m.H(), "mH, mW")
	m.Table.WriteToCell(0, 0, el.NewSpacer(m.Table.Node, 1, 1))

	b := el.NewTextButton(m.Table.Node, 1, m.W()/5, "We're done here.", func() string {
		m.Next(s_news)
		return "fabulous"
	})
	m.Table.WriteToCell(1, 1, b)
	c := el.NewTextButton(m.Table.Node, 1, m.W()/5, "We're done here.", func() string {
		m.Next(s_email)
		return "fabulous"
	})
	m.Table.WriteToCell(1, 2, c)
}

func (m *menu) Update() {
}

func (m *menu) Exit() {
	//start time
}
