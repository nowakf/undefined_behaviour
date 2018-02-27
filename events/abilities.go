package events

type abilities struct {
	normal map[*ability]int
	magic  map[*ability]int
}

func (a *abilities) Add(ability) {
}

func (a *abilities) setAbillity(*ability, int) {
}

func (a *abilities) Queue(ability) {
}

func (a *abilities) removeAbillity(*ability) {
}
