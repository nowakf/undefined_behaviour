package world

//groups offer tasks to their members.
type group struct {
	emailRoot string
	members   []*Person
}

func (g *group) GetActive() []*Person {
	return g.members
}
