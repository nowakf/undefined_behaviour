package events

type relationships struct {
	relations map[*actor]int
}
type relationship struct {
	r1, r2               *actor
	r1feeling, r2feeling int
}

func (r *relationships) AddRelationship(relationship) {
}

func (r *relationships) setRelationship(object *actor, delta int) {}
