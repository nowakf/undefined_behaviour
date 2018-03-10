package world

//virtual is opposed to concrete, in that it encompasses the raw text of events,
//actors, and actions, prior to their parsing into a 'concrete', specific form.
type virtual struct {
	people  map[string]person
	records map[string]Record
	groups  map[string]group
}

func newVirtual() *virtual {
	v := new(virtual)
	v.people = make(map[string]person, 0)
	v.records = make(map[string]Record, 0)
	v.groups = make(map[string]group, 0)
	return v
}
