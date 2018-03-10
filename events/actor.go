package events

import (
	w "ub/events/world"
	p "ub/events/world/person"
)

//actor is a person in the world
type actor struct {
	me *p.Person
	//mailbox
	*goal
	*w.world
}
type goal struct {
}

func (a *actor) whatever() {
	a.me.Stats.Madness()
}
