package main

import "github.com/nowakf/tview"

func Email() *email {
	return &email{}
}

type email struct{}

func (e *email) UI(nextMode func()) (title string, content tview.Primitive) {
	email := tview.NewTextView().
		SetText("email")
	return "Email", email
}
func (e *email) Update() {}

func (e *email) Count() int {
	return 0
}
func (e *email) Get(s string) (handle string, content map[string]string, keyvalue map[string]int) {
	return "", nil, nil
}
