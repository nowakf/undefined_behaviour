package events

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
