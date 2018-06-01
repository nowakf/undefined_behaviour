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
	m.Table.WriteToCell(0, 0, el.NewSpacer(m.Table.Node, 3, m.W()/3))

	b := el.NewTextButton(m.Table.Node, 1, m.W()/5, "NEWS", func() string {
		m.Next(s_news)
		return "fabulous"
	})
	m.Table.WriteToCell(1, 1, b)
	c := el.NewTextButton(m.Table.Node, 1, m.W()/5, "MAIL", func() string {
		m.Next(s_email)
		return "fabulous"
	})
	m.Table.WriteToCell(1, 2, c)
	d := el.NewTextButton(m.Table.Node, 1, m.W()/5, "SETUP1", func() string {
		m.Next(s_setup)
		return "fabulous"
	})
	m.Table.WriteToCell(1, 4, d)
	e := el.NewTextButton(m.Table.Node, 1, m.W()/5, "SETUP2", func() string {
		m.Next(s_setup)
		return "fabulous"
	})
	m.Table.WriteToCell(1, 5, e)
	f := el.NewTextButton(m.Table.Node, 1, m.W()/5, "SETUP3", func() string {
		m.Next(s_setup)
		return "fabulous"
	})
	m.Table.WriteToCell(1, 6, f)
}

func (m *menu) Update() {
}

func (m *menu) Exit() {
	//start time
}
