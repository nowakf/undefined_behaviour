package events

import w "ub/events/world"

type event struct {
	subject *actor
	object  *actor
	world   *w.World
	success bool
	tags    []string //e.g "theft", "arson", "uncanny" etc
}

func (a *event) checkIfHasReccord() {}
