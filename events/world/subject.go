package world

import (
	"math/rand"
	o "ub/events/world/object"
)

type Subject struct {
	o.Object
	goal
	location
}

func (s *Subject) Generate(ID int) {
	s.Object.
		//stats
		Set(o.STR, rand.Intn(16)).
		Set(o.CHA, rand.Intn(16)).
		Set(o.WIS, rand.Intn(16)).
		Set(o.INT, rand.Intn(16)).
		Set(o.WIL, rand.Intn(16)).
		Set(o.SAN, rand.Intn(16)).
		//bools
		Set(o.LawAbiding, rand.Intn(2)).
		//Stress
		Set(o.STRESS, rand.Intn(50)).
		Set(o.ID, ID)

}

func (s *Subject) Observe(w *World) {
	switch stress := s.Get(o.STRESS); {
	case stress > 0:

	case stress > 100:

	case stress > 250:

	}
}
func (s *Subject) Act(w *World) {

}

type goal struct {
}

type location struct {
	APGraph     int
	SocialGraph int
}
