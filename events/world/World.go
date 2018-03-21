// Package events manages events in the game-World
package world

import (
	"ub/data"
)

const (
	pop_number = 256
)

//World is the World as an entity amongst others
type World struct {
	madness  int
	doomsday int
	model    *virtual
	people   [pop_number]Person
	groups   []group
}

func NewWorld(con *WorldConfig) *World {
	w := new(World)
	w.model = newVirtual()
	return w
}
func LoadWorld(file *data.Save) *World {
	return new(World)
}

func (w *World) Generate(conf *WorldConfig) {
}
