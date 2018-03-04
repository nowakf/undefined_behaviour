package events

type event struct {
	subject *actor
	object  *actor
	ability *ability
	world   *world
	success bool
	tags    []string //e.g "theft", "arson", "uncanny" etc
}

func (a *event) checkIfHasReccord() {}