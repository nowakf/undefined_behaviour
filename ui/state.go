package ui

import (
	el "ub/ui/elements"
)

type state interface {
	el.UiElement
	Start()
	Update()
	Exit()
	Next(stateEnum)
	Flush()
	Name() string
}

type monitor interface {
	Monitor() chan int
	Listener()
}
