package ui

import (
	"testing"
)

var UI = new(ui)

func test_mousepos(t *testing.T) {
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			fx, fy := float64(x), float64(y)
			mx, my := UI.mousepos(fx, fy, 100, 100)
			if mx > 1 || my > 1 || mx < 0 || my < 0 {
				t.Fatal("mousepos is going out of bounds?")
			}

		}
	}
}

func test_floatToCellCoord(t *testing.T) {
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			fx, fy := float64(x), float64(y)

			outx, outy := UI.floatToCellCoord(fx/100, fy/100, 10, 10)
			if outx != x/10 || outy != y/10 {
				t.Fatalf("cellcoord is %v,%v, it should be %v, %v", outx, outy, x/10, y/10)
			}

		}
	}
}
