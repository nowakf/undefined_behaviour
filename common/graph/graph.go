//implements a graph - by default directed. Each edge can have a gate
package graph

import (
	"container/heap"
	"fmt"
	"math"
)

type graph struct {
	vertexes []interface{}
	edges    map[int]map[int]edge
}

func New() *graph {
	g := new(graph)
	g.vertexes = make([]interface{}, 0)
	g.edges = make(map[int]map[int]edge)
	return g
}

//vertex adds a vertex. A vertex can store whatever
func (g *graph) AddVertex(v interface{}) error {
	g.vertexes = append(g.vertexes, v)
	return nil
}

//link creates a one-way link between vertices. The gate
//can be whatever
func (g *graph) Link(a int, b int, weight int, gate interface{}) error {
	_, ok := g.edges[a]
	if !ok {
		g.edges[a] = make(map[int]edge)
	}
	g.edges[a][b] = edge{weight, gate}
	return nil
}

//DoubleLink creates a two-way link between vertices
func (g *graph) DoubleLink(a int, b int, weight int, gate interface{}) error {
	_, ok := g.edges[a]
	if !ok {
		g.edges[a] = make(map[int]edge)
	}
	_, ok = g.edges[b]
	if !ok {
		g.edges[b] = make(map[int]edge)
	}
	e := edge{weight, gate}
	g.edges[a][b], g.edges[b][a] = e, e
	return nil
}

//BiDirectionalSearch returns the lowest cost path between two points. It's probably buggy
func (g *graph) BiDirectionalSearch(sourceIndex int, sinkIndex int, checkFunc func(interface{}) bool) (path, error) {
	sourceNode, sinkNode := node{source, -1, sourceIndex, 0}, node{sink, -1, sinkIndex, 0}

	frontier := frontier{&sourceNode, &sinkNode}

	heap.Init(&frontier)

	history := make(map[int]node, 0)
	history[sourceIndex], history[sinkIndex] = sourceNode, sinkNode

	lastValues := make([]int, 2)
	cost := math.MaxInt64

	for frontier.Len() > 0 {

		n := *(heap.Pop(&frontier).(*node))

		v := n.visitedBy

		for destination, edge := range g.edges[n.value] {
			if checkFunc(edge.gate) {
				last, ok := history[destination]
				lastValues[v] = last.value
				if !ok {
					nn := node{v, n.value, destination, n.priority + edge.weight}
					heap.Push(&frontier, &nn)
					history[destination] = nn
				} else {
					if last.visitedBy != v && last.priority < cost {
						cost = last.priority
					}
				}
			}

		}
	}
	if cost < math.MaxInt64 {
		return g.unspool(history, lastValues[source], lastValues[sink])
	} else {
		return path{}, nil
	}

}
func (g *graph) unspool(history map[int]node, lastSource int, lastSink int) (path, error) {

	spooler := func(current int) []int {
		spool := make(path, 0)
		for {
			next, ok := history[current]
			if !ok {
				break
			}
			delete(history, current)
			spool = append(spool, next.value)
			current = next.predecessor

		}
		return spool
	}

	return append(spooler(lastSource), spooler(lastSink)...), nil

}

type path []int

func (p path) String() string {
	s := ""
	for _, n := range p {
		s += fmt.Sprintf("-(%v)-", n)
	}
	return s
}

type visitor int

const (
	sink visitor = iota
	source
)

type node struct {
	visitedBy   visitor
	predecessor int
	value       int
	priority    int
}

//frontier is a heap of nodes to be visited - it is a priority queue
type frontier []*node

func (f frontier) Len() int            { return len(f) }
func (f frontier) Less(i, j int) bool  { return f[i].priority > f[j].priority }
func (f frontier) Swap(i, j int)       { f[i], f[j] = f[j], f[i] }
func (f *frontier) Push(n interface{}) { *f = append(*f, n.(*node)) }
func (f *frontier) Pop() interface{} {
	old := *f
	n := len(old)
	x := old[n-1]
	*f = old[0 : n-1]
	return x
}

//check tests if an actor can pass a gate

type vertex struct {
}

type edge struct {
	weight int
	gate   interface{}
}
