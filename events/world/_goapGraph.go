package world

import (
	"container/heap"

	"github.com/nowakf/undefined_behaviour/events/world/items"
	"github.com/nowakf/undefined_behaviour/events/world/stats"
)

type graph []vertex

type vertex struct {
	required bitmap
	delta    bitmap
	upstream []int
	cost     int
}

func Graph() *graph {
	g := new(graph)

	return g
}
func (g *graph) AddVertex(required, delta bitmap) {

	*g = append(*g, vertex{required: required, delta: delta})

	ref := len(*g) - 1

	for i := range *g {
		if (*g)[i].required&delta != 0 {
			//it intersects, so add it to the nodes upstream list
			(*g)[i].upstream = append((*g)[i].upstream)
		}
		if (*g)[i].delta&required != 0 {
			//it intersects, so it's upstream - add it to this vertexes'
			(*g)[ref].upstream = append((*g)[ref].upstream, i)
		}
	}

}

func (g *graph) unspool(start node) (out path) {
	current := start
	for {
		out.path = append(out.path, current.self)
		out.pathCost += current.cost
		if current.parent == nil {
			break
		}
		current = *current.parent
	}
	return
}

type Agent struct {
	items.Items
	stats.Stats
}

func (a Agent) Superset(b Agent) bool {
	return a.Items.Superset(b.Items) && a.Stats.Superset(b.Stats)
}

func (g graph) Search(searcher Agent, goal int) path {

	frontier := frontier{node{self: goal, parent: nil, cost: 0}}
	//add the goal node

	var check bitmap

	heap.Init(&frontier)
	//initialize the heap

	visited := make(map[int]int)

	for len(frontier) > 0 {

		location := heap.Pop(&frontier).(node)

		visited[location.self]++

		if (*g)[location.self]&^check == 0 {
			//we have all the required fields, so here's a start point
			path := g.unspool(location) //unspool

		}
		//return all the nodes you pass checks for,
		//and if the node produces a delta, a reference to itself x the diff between
		//character + value of gate
		for _, next := range g[location.self].upstream {
			intersection = g[location.self].delta & g[next].required

			if intersection != 0 {
				cost := location.cost
				for _, v := range intersection.set {
					cost += location.cost
					heap.Push(&frontier, node{
						self:   self,
						parent: &location,
						cost:   cost,
					})

				}

			}
			heap.Push(&frontier, node{
				self:   next,
				parent: &location,
				cost:   location.cost + g[next].cost,
			})
		}

	}
	return nil

}
func (g graph) Test(searcher Agent, vert int) []int {
	//check the bools
	//extract the scalars
	//check the scalars
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
