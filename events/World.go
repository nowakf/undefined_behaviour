// Package events manages events in the game-world
package events

import "ub/data"

//world contains concrete actors, organizations, and their states
type world struct {
	madness  int
	doomsday int
	ev       *EventSystem
	model    *virtual
	actors   [256]actor
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

func (w *world) setDoomsday(delta int) {
	w.doomsday += delta
}
func (w *world) setMadness(delta int) {
	w.doomsday += delta
}

func (w *world) GetDoomsday() int {
	return w.doomsday
}
func (w *world) GetActiveActors() []*actor {
	active := make([]*actor, 0)
	for _, group := range w.groups {
		active = append(active, group.GetActive()...)
	}
	return active
}
func (w *world) Generate(conf *WorldConfig) {
}
