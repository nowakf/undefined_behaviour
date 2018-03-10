// Package events manages events in the game-world
package world

import (
	"ub/data"
)

const (
	pop_number = 256
)

//world is the world as an entity amongst others
type world struct {
	madness  int
	doomsday int
	model    *virtual
	people   [pop_number]person
	groups   []group
}

func NewWorld(con *WorldConfig) *world {
	w := new(world)
	w.model = newVirtual()
	return w
}
func LoadWorld(file *data.Save) *world {
	return new(world)
}

func (w *world) Generate(conf *WorldConfig) {
}
