package events

import w "ub/events/world"

type Player struct {
	newsFeed []*w.Record
	mailBox  []*w.Record
	IsNew    bool
	*w.Person
}

func LoadPlayer() *Player {
	return new(Player)
}
func NewPlayer(conf *PlayerConfig) *Player {
	p := new(Player)
	p.mailBox = make([]*w.Record, 0)
	p.IsNew = true
	return p
}

func (p *Player) feed(viewer chan *w.Record, list []*w.Record) func() int {
	unviewed := 0
	return func() int {
		for i := unviewed; i < len(list); i++ {
			viewer <- list[i]
			unviewed++
		}
		return unviewed
	}
}

func (p *Player) News(viewer chan *w.Record) func() int {
	return p.feed(viewer, p.mailBox)
}
func (p *Player) Mail(viewer chan *w.Record) func() int {
	return p.feed(viewer, p.newsFeed)
}

type PlayerConfig struct{}
