package world

import (
	"sync"
	o "ub/events/world/object"
)

type person struct {
	m sync.Mutex
	o.Identity
	o.Body
	o.Mind
	o.Social
}

func NewPerson() {}

func (p *person) Strength() *slow {
	return &slow{&p.m, p.Body.Get(), o.STR}
}

type slow struct {
	*sync.Mutex
	o.Value
	index int
}

func (s *slow) Get() int {
	s.Lock()
	defer s.Unlock()
	return s.Value.Get(s.index)
}

func (s *slow) Set(nu int) {
	s.Lock()
	s.Value.Set(s.index, nu)
	s.Unlock()
}
