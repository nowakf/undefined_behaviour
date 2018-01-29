package ui_test

import (
	ui "cthu3/ui"
	"testing"
)

func TestNewSetup(t *testing.T) {
	if ui.State.Draw(ui.NewSetup(3, 4), 3, 4) == nil {
		t.Fatalf("doesn't make a state!")
	}
}
