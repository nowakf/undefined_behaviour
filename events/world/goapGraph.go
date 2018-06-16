package world

import "container/heap"

type graph []vertex

type vertex struct {
	minRequiredCharacter  uint64
	minRequiredPossesions uint64
	neighbours            []int
	cost                  int
}

func (v vertex) Required() Agent {
	return Agent{v.minRequiredPossesions, v.minRequiredCharacter}
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

func (g *graph) unspool(start node, searcher Agent) (out path) {
	current := start
	for {
		out.path = append(out.path, current.self)
		out.pathCost += current.cost
		out.resultingAgent.Union((*g)[current.self].Required())
		if current.parent == nil {
			break
		}
		current = *current.parent
	}
	return
}

type Agent struct {
	possesions uint64
	character  uint64
}

func (a *Agent) Union(b Agent) *Agent        { return a }
func (a *Agent) Intersection(b Agent) *Agent { return a }
func (a *Agent) Difference(b Agent) *Agent   { return a }

func (g graph) Search(searcher Agent, goal int) (paths []path) {

	frontier := frontier{node{self: goal, parent: nil, cost: 0}}
	//add the goal node

	heap.Init(&frontier)
	//initialize the heap

	visited := make(map[int]bool)

	//break when we have all the characteristics we need to reach the goal - or if there's no way,
	//break anyway
	for len(frontier) > 0 {

		location := heap.Pop(&frontier).(node)

		visited[location.self] = true

		hasRequiredStats, hasRequiredThings := g.Test(location.self)

		if hasRequiredStats(searcher.character) && hasRequiredThings(searcher.possesions) {
			//we've found a path!
			path := g.unspool(location, searcher) //unspool
			paths = append(paths, path)

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
func (g graph) SortSearched(paths []path) (completeRoute []int) {
	//sort paths by cost
	//then check each in turn to see if the path moves the initial agent towards the goal
	//then, if the agent is at the goal, return the complete route there
	//if the agent cannot reach the goal, return nil
	return nil
}

type path struct {
	pathCost       int
	resultingAgent Agent
	path           []int
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
