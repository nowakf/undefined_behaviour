package graph

import "testing"
import "math/rand"
import "time"
import "fmt"

const graphSize = 10

type gate uint32

func randGraph(vertexCount int) *graph {
	rand.Seed(time.Now().UnixNano())
	g := New()
	for i := 0; i < vertexCount; i++ {
		g.AddVertex("whatever")
	}
	for i := 0; i < vertexCount*2; i++ {
		g.Link(rand.Intn(vertexCount), rand.Intn(vertexCount), rand.Int(), gate(rand.Uint32()))
	}
	return g
}

func TestBiDirectionalSearch(t *testing.T) {
	g := randGraph(graphSize)
	for i := 0; i < 20; i++ {
		check := func(g interface{}) bool {
			return g.(gate) < gate(rand.Uint32())

		}
		path, _ := g.BiDirectionalSearch(rand.Intn(graphSize), rand.Intn(graphSize), check)
		fmt.Printf("path %v is %s,\n", i, path.String())
	}
}
