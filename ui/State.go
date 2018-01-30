package ui

import c "cthu3/common"

type State interface {
	Draw(x, y int) []c.Cell
	OnMouse(x, y int, pressed bool, released bool) func() string
}
