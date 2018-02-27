package events

type action struct {
	subject *actor
	object  *actor
	ability *ability
	world   *world
	success bool
	tags    []string //e.g "theft", "arson", "uncanny" etc
}

func (a *action) checkIfEvent() {}
