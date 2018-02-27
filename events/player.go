package events

type Player struct {
	newsFeed []*Event
	mailBox  []*Event
	IsNew    bool
	*actor
}

func LoadPlayer() *Player {
	return new(Player)
}
func NewPlayer(conf *PlayerConfig) *Player {
	p := new(Player)
	p.actor = NewActor()
	p.mailBox = make([]*Event, 0)
	p.IsNew = true
	return p
}

func (p *Player) feed(viewer chan *Event, list []*Event) func() int {
	unviewed := 0
	return func() int {
		for i := unviewed; i < len(list); i++ {
			viewer <- list[i]
			unviewed++
		}
		return unviewed
	}
}

func (p *Player) News(viewer chan *Event) func() int {
	return p.feed(viewer, p.mailBox)
}
func (p *Player) Mail(viewer chan *Event) func() int {
	return p.feed(viewer, p.newsFeed)
}

type PlayerConfig struct{}
