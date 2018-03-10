package graph

import p "ub/events/person"

type graph struct {
	edges    [][]edge //first slice represents 'from', second 'to' vertex
	vertexes []vertex //indexed to the first slice above
}

func New() *graph {
	g := new(graph)
	g.edges = make([][]edge, 0)
	g.vertexes = make([]vertex, 0)
	return g
}
func (g *graph) AddVertex(name string, description string, does func(*p.Person) error) {
	g.vertexes = append(g.vertexes, vertex{name, description, does})
	g.edges = append(g.edges, make([]edge, 0))
}
func (g *graph) AddEdge(from, to, cost int, gate uint64) error {
	err := g.checkVerts(from, to)
	if err != nil {
		return err
	} else {
		g.edges[from] = append(g.edges[from], edge{gate, cost, to})
		return nil
	}
}
func (g *graph) Path(from vertex, to vertex, walker *p.Person) []vertex {
	//some kind of pathfinding
	return g.vertexes
}

type vertError struct {
	from, to, vertexes int
}

func (v vertError) Error() string {
	return "NoVertex: from index: " + string(v.from) + "to index: " + string(v.to) + "vertex length: " + string(v.vertexes)
}
func (g *graph) checkVerts(from int, to int) error {
	if from >= len(g.vertexes) || to >= len(g.vertexes) {
		return vertError{from, to, len(g.vertexes)}
	} else {
		return nil
	}
}

type edge struct {
	gate uint64
	cost int
	to   int
}
type vertex struct {
	name, description string
	transform         func(*p.Person) error
}
