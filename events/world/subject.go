package world

import (
	"math/rand"

	"github.com/nowakf/undefined_behaviour/events/world/items"
	"github.com/nowakf/undefined_behaviour/events/world/stats"
)

//since the pop no is 256, an uint8 is fine
type ID uint8

type Subject struct {
	items   items.Items
	stats   stats.Stats
	friends [max_relationships]ID
	goal    int
}

type generationError string

func (g generationError) Error() string {
	return string(g)
}

func (s *Subject) Generate(ID ID) error {
	if ID == 0 {
		return generationError("ID 0 is reserved for the null identifier")
	}
	//set character
	s.stats.Set(stats.LawAbiding, rand.Intn(2)).
		Set(stats.Depression, rand.Intn(2)).
		Set(stats.STR, rand.Intn(16)).
		Set(stats.CHA, rand.Intn(16)).
		Set(stats.WIS, rand.Intn(16)).
		Set(stats.INT, rand.Intn(16)).
		Set(stats.WIL, rand.Intn(16)).
		Set(stats.SAN, rand.Intn(5)).
		Set(stats.ID, int(ID)).
		Set(stats.Stress, rand.Intn(50))

	//set items
	s.items.Set(items.Money, rand.Intn(100))
	return nil

}

func (s *Subject) Observe(w *World) {
	stress := w.Madness
	switch {
	case stress > 0:

	case stress > 100:

	case stress > 250:

	}
}
func (s *Subject) Act(w *World) {

}
