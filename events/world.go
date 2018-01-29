// Package events manages events in the game-world
package events

import (
	"cthu3/data"
)

type world struct {

	//actors
	//player
}

func NewWorld(co *worldConfig, t *data.TextData) *world {
	w := new(world)
	return w
}

func (w *world) Generate(co *worldConfig, t *data.TextData) {
}
