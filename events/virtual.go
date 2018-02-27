package events

//virtual is opposed to concrete, in that it encompasses the raw text of events,
//actors, and actions, prior to their parsing into a 'concrete', specific form.

type virtual struct {
	abiliies map[string]ability
	actors   map[string]actor
	Events   map[string]Event
	groups   map[string]group
}

func newVirtual() *virtual {
	v := new(virtual)
	v.abiliies = make(map[string]ability, 0)
	v.actors = make(map[string]actor, 0)
	v.Events = make(map[string]Event, 0)
	v.groups = make(map[string]group, 0)
	return v
}
