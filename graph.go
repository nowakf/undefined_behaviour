package main

import "github.com/nowakf/tview"

func Graph() *graph {
	return &graph{}
}

type graph struct{}

func (g *graph) UI(nextMode func()) (title string, content tview.Primitive) {
	graph := tview.NewTextView().
		SetText("graph")
	return "Graph", graph
}
func (g *graph) Update() {}
func (g *graph) Count() int {
	return 0
}
func (g *graph) Get(s string) (handle string, content map[string]string, keyvalue map[string]int) {
	return "", nil, nil
}
