package main

import "github.com/nowakf/tview"

func News() *news {
	return &news{}
}

type news struct{}

func (n *news) UI(nextMode func()) (title string, content tview.Primitive) {
	news := tview.NewTextView().
		SetText("news")
	return "News", news
}
func (n *news) Update() {}

func (n *news) Count() int {
	return 0
}
func (n *news) Get(s string) (handle string, content map[string]string, keyvalue map[string]int) {
	return "", nil, nil
}
