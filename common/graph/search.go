package graph

import _ "container/heap"
import "sort"
import "math/rand"
import "fmt"

//flow returns a large flow-field for a graph
//the idea is to have this as pathfinding for agents
type flow struct {
	g    *graph
	flow [][]edge
	//[start index][goal index]neighborindex
}

const timeOut = 20

type flowError struct{ err string }

func (f flowError) Error() string { return f.err }

func Flow(g *graph) (*flow, error) {
	f := new(flow)

	f.g = g //so the graph is fixed at the point where the flow is generated

	f.flow = make([][]edge, len(g.vertexes))

	var err error
	//sort all the verts
	for _, vert := range g.vertexes {
		sort.Sort(vert)
	}
	//initialize index and flow arrays
	for start, vert := range g.vertexes {
		f.flow[start] = f.dijkastra(&vert)
	}
	//run dijkastra on each vertex for all neighbors
	//this must be done after the index is built

	return f, err
}

func (f *flow) dijkastra(start *vertex) []edge {

	return []edge{}

}

func (f *flow) Step(actor uint32, start int, goal int) (newpos int, pathfinding error) {
	e := f.flow[start][goal]
	if f.gate(actor, e.Gate()) {
		return e.to, nil
	} else {
		return f.randomStep(actor, start)
	}
}
func (f *flow) randomStep(actor uint32, start int) (newpos int, pathfinding error) {
	length := len(f.flow[start])
	for i := 0; i < timeOut; i++ {
		next := f.flow[start][rand.Intn(length)]
		if f.gate(actor, next.Gate()) {
			return next.to, nil
		}
	}
	return start, pathfindingError{fmt.Sprintf("could not find path from %v", start)}
}
func (f *flow) gate(actor uint32, gate uint32) bool {
	return true
}

type pathfindingError struct{ err string }

func (p pathfindingError) Error() string { return p.err }
