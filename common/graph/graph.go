package graph

type graph struct {
	vertexes []vertex
}

func New() *graph {
	g := new(graph)
	g.vertexes = make([]vertex, 0)
	return g
}

func (g *graph) Vertex(label string) {
	g.vertexes = append(g.vertexes, vertex{label, make([]edge, 0)})
}

//implements sort
type vertex struct {
	label string
	edges []edge
}

func (v vertex) Edge(gate uint32, weight int, to int) {
	v.edges = append(v.edges, edge{gate, weight, to})
}
func (v vertex) Len() int           { return len(v.edges) }
func (v vertex) Less(i, j int) bool { return v.edges[i].weight <= v.edges[j].weight }
func (v vertex) Swap(i, j int) {
	temp := v.edges[i]
	v.edges[i] = v.edges[j]
	v.edges[j] = temp
}

type edge struct {
	gate   uint32
	weight int
	to     int
}

func (e *edge) Gate() uint32 { return e.gate }

func (e *edge) Weight() int { return e.weight }

type graphError struct{ err string }

func (g graphError) Error() string { return g.err }
