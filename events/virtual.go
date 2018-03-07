package events

//virtual is opposed to concrete, in that it encompasses the raw text of events,
//actors, and actions, prior to their parsing into a 'concrete', specific form.

type virtual struct {
	actors  map[string]actor
	records map[string]Record
	groups  map[string]group
}

func newVirtual() *virtual {
	v := new(virtual)
	v.actors = make(map[string]actor, 0)
	v.records = make(map[string]Record, 0)
	v.groups = make(map[string]group, 0)
	return v
}
