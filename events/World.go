// Package events manages events in the game-World
package events

//World contains concrete actors
type World struct {

	//actors
	//player
}

func NewWorld(co *WorldConfig) *World {
	w := new(World)
	return w
}

func (w *World) Generate(co *WorldConfig) {
}
