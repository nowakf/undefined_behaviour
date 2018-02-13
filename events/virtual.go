package events

//virtual is opposed to concrete, in that it encompasses the raw text of events,
//actors, and actions, prior to their parsing into a 'concrete', specific form.

type virtual struct {
	Actions map[string]Action
	Actors  map[string]Actor
	Events  map[string]Event
}

func newVirtual() *virtual {
	v := new(virtual)
	v.Actions = make(map[string]Action)
	v.Actors = make(map[string]Actor)
	v.Events = make(map[string]Event)

	return v
}
