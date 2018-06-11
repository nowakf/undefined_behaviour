package main

import (
	"fmt"
	"strconv"
	c "ub/common"

	"github.com/nowakf/pixel"
	"github.com/nowakf/pixel/pixelgl"
	"github.com/nowakf/tview"
)

func run() {
	//--------------------------------------------------------------//
	//	CONFIGURATION						//
	//--------------------------------------------------------------//
	windowConfig := pixelgl.WindowConfig{
		Resizable: true,
		Bounds:    pixel.R(0, 0, 1024, 1024),
		Monitor:   pixelgl.PrimaryMonitor(),
	}

	app, err := tview.NewApplication(
		&tview.Config{
			FontSize:     18,
			FontPath:     "./assets/fonts/DejaVuSansMono.ttf",
			AdjustX:      -2,
			AdjustY:      -2,
			DPI:          72,
			WindowConfig: windowConfig,
		})
	//--------------------------------------------------------------//
	//								//
	//--------------------------------------------------------------//

	c.Check(err)

	//generate the player here

	//these are the major modes of the program
	majorModes := []MajorMode{
		Book(),
		News(),
		Email(),
		Graph(),
	}
	//the bottom has some global info
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)

	currentMode := 0
	info.Highlight(strconv.Itoa(currentMode))
	pages := tview.NewPages()

	previousMode := func() {
		currentMode = (currentMode - 1 + len(majorModes)) % len(majorModes)
		//wraps around
		if majorModes[currentMode].Count() > 0 {
			majorModes[currentMode].Update()
		}
		info.Highlight(strconv.Itoa(currentMode))
		pages.SwitchToPage(strconv.Itoa(currentMode))

	}
	nextMode := func() {
		if majorModes[currentMode].Count() > 0 {
			majorModes[currentMode].Update()
		}
		currentMode = (currentMode + 1) % len(majorModes)
		info.Highlight(strconv.Itoa(currentMode))
		pages.SwitchToPage(strconv.Itoa(currentMode))
	}

	for index, mode := range majorModes {
		title, primitive := mode.UI(nextMode)
		pages.AddPage(strconv.Itoa(index),
			primitive, true, index == currentMode)
		fmt.Fprintf(info, `%d	["%d"][lightgrey] %s [""]	[""]`, index+1, index, title)
		//print the title on the info?
	}
	majorModes[currentMode].Update()
	//create the main layout
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	app.SetInputCapture(func(event *pixelgl.KeyEv) *pixelgl.KeyEv {
		if event.Key == pixelgl.KeyTab && event.M == pixelgl.ModShift {
			previousMode()
		} else if event.Key == pixelgl.KeyTab {
			nextMode()
		}
		return event

	})

	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}

}
func main() {
	pixelgl.Run(run)
}

type MajorMode interface {
	UI(nextMode func()) (title string, content tview.Primitive)
	Count() int
	Update()
}
