// Package events manages events in the game-World
package world

import (
	"github.com/nowakf/undefined_behaviour/data"
	"math/rand"
)

const (
	max_friends       = 4
	max_enemies       = 3
	max_relationships = 8
	max_love          = 256
	max_hate          = 256
	pop_number        = 256
)

//World is the World as an entity amongst others
type World struct {
	Madness  int
	Doomsday int
	People   [pop_number]Subject
	Groups   []Group
	graph
}

func NewWorld() *World {
	return &World{}
}
func LoadWorld(file *data.Save) *World {
	return new(World)
}

func (w *World) Generate(conf *WorldConfig) {
	w.generateRelationships()
	w.assignGroups()

}

func (w *World) generateRelationships() {
}
func (w *World) assignGroups() {
	for _, person := range w.People {
		i := rand.Intn(len(w.Groups) * 2)
		if i < len(w.Groups) {
			w.Groups[i].AddMember(&person)
		}
		//else is in no group
	}
}
func (w *World) relationship(a int, b int, bias int) int {
	return rand.Intn(bias)
}

//this is so if necessary I can make this personality-dependent
func (w *World) numberOfFriends(p Subject) int {
	return rand.Intn(max_friends)
}

func (w *World) numberOfEnemies(p Subject) int {
	return rand.Intn(max_enemies)
}
func (w *World) numberOfFamilly(p Subject) int {
	return rand.Intn(max_enemies)
}

type relationship struct {
}
