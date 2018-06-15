package simpleGraph

import "container/heap"

type graph []vertex

type vertex struct {
	minRequiredCharacter  uint64
	minRequiredPossesions uint64
	neighbours            []int
	cost                  int
}

var predefinedVertexes = []vertex{}

const (
	maxChar       = 0xFFFFFFFFFFF
	maxPossession = 0xFFFFFFFFFFF
)

func Graph() *graph {
	g := new(graph)
	refs, passChars, passPossessions, costs := g.ParseTests(maxChar, maxPossession)
	for i := range refs {
		g.AddVertex(passChars[i], passPossessions[i], costs[i])
	}

	return g
}

func (g *graph) AddVertex(minRequiredPossesions uint64, minRequiredCharacter uint64, cost int) {
	neighbours, _, _, _ := g.ParseTests(minRequiredCharacter, minRequiredPossesions)

	*g = append(*g, vertex{
		minRequiredCharacter:  minRequiredCharacter,
		minRequiredPossesions: minRequiredPossesions,
		neighbours:            neighbours,
		cost:                  cost,
	})

}
func (g *graph) ParseTests(minRequiredCharacter uint64, minRequiredPossesions uint64) (ref []int, passChar []uint64, passPossession []uint64, cost []int) {
	refs := make([]int, 0)
	passChars := make([]uint64, 0)
	passPossessions := make([]uint64, 0)
	costs := make([]int, 0)
	//for bool condition of test
	//for constanttests {if true, append i}
	//for scalar condition of test
	//for scalareconstanttests{if true, append i}
	//append i
	//for predefinedVertexes{if true, append i}
	return refs, passChars, passPossessions, costs

}

func (g *graph) Test(location int) (hasRequiredStats func(uint64) bool, hasRequiredThings func(uint64) bool) {
	hasRequiredStats = func(stats uint64) bool {
		return true
	}
	hasRequiredThings = func(things uint64) bool {
		return true
	}
	return
}

func (g *graph) unspool(final node) (path []int) {
	current := final
	for {
		path = append(path, current.self)
		if current.parent == nil {
			break
		}
		current = *current.parent
	}
	return
}

func (g graph) Search(possesions uint64, character uint64, goal int) (path []int) {

	frontier := frontier{node{self: goal, parent: nil, cost: 0}}
	heap.Init(&frontier)

	visited := make(map[int]bool)

	for len(frontier) > 0 {
		location := heap.Pop(&frontier).(node)

		visited[location.self] = true

		hasRequiredStats, hasRequiredThings := g.Test(location.self)
		if hasRequiredStats(character) && hasRequiredThings(possesions) {
			//we pass the tests, so we're done.
			path = g.unspool(location) //unspool
			return
		}

		for _, next := range g[location.self].neighbours {
			if !visited[next] {
				heap.Push(&frontier, node{
					self:   next,
					parent: &location,
					cost:   location.cost + g[next].cost})
			}
		}

	}
	return nil

}

type frontier []node

func (f frontier) Len() int {
	return len(f)
}

func (f frontier) Less(i int, j int) bool {
	return f[i].cost > f[j].cost
}

func (f frontier) Swap(i int, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *frontier) Push(x interface{}) {
	*f = append(*f, x.(node))
}

func (f *frontier) Pop() interface{} {
	old := *f
	n := len(old)
	x := old[n-1]
	*f = old[0 : n-1]
	return x
}

type node struct {
	self   int
	parent *node
	cost   int
}
