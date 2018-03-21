package events

import (
	w "ub/events/world"
)

//actor is a person in the world
type actor struct {
	*w.Person
	*goal
	*w.World
}

func (a *actor) whatever() {
	a.Data.Get(a.Body().Strength())
}

type goal struct {
}
