package events

import w "/github.com/nowakf/undefined_behaviour/events/world"

type event struct {
	s/github.com/nowakf/undefined_behaviourject *actor
	Object  *actor
	world   *w.World
	success bool
	tags    []string //e.g "theft", "arson", "uncanny" etc
}

func (a *event) checkIfHasReccord() {}
