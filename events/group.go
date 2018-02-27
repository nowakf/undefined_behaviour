package events

//groups offer tasks to their members.
type group struct {
	emailRoot string
	members   []*actor
}

func (g *group) GetActive() []*actor {
	return g.members
}
