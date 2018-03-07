package events

import (
	p "ub/events/person"
)

//actor is a person in the world
type actor struct {
	me *p.Person
	*goal
	*world
}

func (a *actor) whatever() {
	a.me.Stats.Madness()
}
