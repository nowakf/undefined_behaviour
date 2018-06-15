package world

//groups offer tasks to their members.
type Group struct {
	Name        string     `yaml:"name"`
	Tags        string     `yaml:"tags"`
	Description string     `yaml:"description"`
	Goals       []goal     `yaml:"goals"`
	EmailRoot   string     `yaml:"emailRoot"`
	members     []*Subject `yaml:"members"`
}

func (g *Group) AddMember(p *Subject) {
	g.members = append(g.members, p)
}
