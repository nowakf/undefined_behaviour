package events

type virtual struct {
	Actions map[string]action
	Actors  map[string]actor
	Events  map[string]event
}

func newVirtual() *virtual {
	v := new(virtual)
	v.Actions = make(map[string]action)
	v.Actors = make(map[string]actor)
	v.Events = make(map[string]event)

	return v
}
