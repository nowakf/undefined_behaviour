package ui

import el "ub/ui/elements"

type state interface {
	el.UiElement
	HasNew() bool
}
