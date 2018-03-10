package world

//groups offer tasks to their members.
type group struct {
	emailRoot string
	members   []*person
}

func (g *group) GetActive() []*person {
	return g.members
}
