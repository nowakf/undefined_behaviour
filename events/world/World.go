// Package events manages events in the game-World
package world

import (
	"math/rand"
	"ub/common/graph"
	"ub/data"
)

const (
	max_friends       = 4
	max_enemies       = 3
	max_relationships = 12
	max_love          = 255
	max_hate          = 255
	pop_number        = 256
)

//World is the World as an entity amongst others
type World struct {
	Madness             int
	Doomsday            int
	model               *virtual
	People              [pop_number]Subject
	Groups              []Group
	ActionPlanningGraph graph.Graph
	SocialGraph         graph.Graph
}

func NewWorld() *World {
	return &World{}
}
func LoadWorld(file *data.Save) *World {
	return new(World)
}

func (w *World) Generate(conf *WorldConfig) {
	w.generateSocialGraph()
	w.assignGroups()

}

func (w *World) generateSocialGraph() {
	society := graph.New()
	for i, person := range w.People {
		err := society.AddVertex(person)

		var relationships [max_relationships]int

		for j := 0; j < max_relationships; j++ {
			relationships[j] = rand.Intn(pop_number)

		}

		numberOfFriends := w.numberOfFriends(person)
		numberOfEnemies := w.numberOfEnemies(person)
		numberOfFamilly := w.numberOfFamilly(person)

		j := max_relationships

		//add friends
		for j >= 0 && j > max_relationships-numberOfFriends {
			j--
			err = society.DoubleLink(
				i,
				relationships[j],
				w.relationship(i, j, rand.Intn(max_love)),
				nil,
			)
		}
		//add enemies
		for j >= 0 && j > max_relationships-numberOfFriends-numberOfEnemies {
			j--

			err = society.DoubleLink(
				i,
				relationships[j],
				w.relationship(i, j, -rand.Intn(max_hate)),
				nil,
			)

		}

		for j >= 0 && j > max_relationships-numberOfFriends-numberOfEnemies-numberOfFamilly {
			j--
			err = society.DoubleLink(
				i,
				relationships[j],
				w.relationship(i, j, 0),
				nil,
			)
		}
		if err != nil {
			panic(err)
		}

		//add family
	}
	w.SocialGraph = society
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
