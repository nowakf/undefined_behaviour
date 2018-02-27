// ui manages user input, the abstract display of objects, and
// sends messages to the event system, when input changes game state.
package ui

import (
	_ "fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	c "ub/common"
	"ub/events"
	el "ub/ui/elements"
)

// a ui displays the stuff that's in view, manages changes between states,
// and sends game-state-altering input to the event system
type ui struct {
	h, w int

	win *pixelgl.Window

	player *events.Player

	evSys *events.EventSystem

	*linker
	states          map[stateEnum]state
	monitoredStates map[stateEnum]monitor

	notifier map[string]int

	*mouse
	*keyboard

	focused el.KeyCatcher
}

type stateEnum int

const (
	s_email stateEnum = iota
	s_news
	s_net
	s_book
	s_setup
	s_menu
)

//creates a new UI, returns a pointer
func NewUI(h, w int, win *pixelgl.Window, e *events.EventSystem) *ui {

	u := new(ui)
	u.mouse = &mouse{win: win}
	u.keyboard = &keyboard{win: win}
	//take the first state, setup
	u.win = win
	u.keyboard.context = u.inputRules()
	return u
}

func (u *ui) inputRules() el.KeyCatcher {
	k := el.GlobalKeyChecker()
	k.Add(pixelgl.KeyEscape, func() {
		u.linker.Next(s_menu)
	})
	return k
}

func (u *ui) Start(h, w int) {
	//generate the states
	//start on the menu

	v := NewViewer(h, w)

	defaultConfig := events.PlayerConfig{}
	u.player = events.NewPlayer(&defaultConfig)

	l := linker{links: &u.states, current: u.states[s_menu]}

	u.linker = &l

	u.states = map[stateEnum]state{
		s_menu:  NewMenu(v, u.linker, u.evSys),
		s_setup: NewSetup(v, u.linker, u.player),
		s_email: NewEmailViewer(v, u.linker, u.player),
		s_news:  NewNewsViewer(v, u.linker, u.player),
	}

	u.monitoredStates = map[stateEnum]monitor{
		s_email: u.states[s_email].(monitor),
		s_news:  u.states[s_news].(monitor),
	}

	for _, state := range u.monitoredStates {
		go state.Listener()
	}

	u.linker.Next(s_menu)

}

//resize calls resize on all of the states...
func (u *ui) Resize(h, w int) {
	u.current.Resize(h, w)
}

func (u *ui) HasNew(display map[string]int, states map[stateEnum]monitor) map[string]int {
	for _, s := range states {
		select {
		case count := <-s.Monitor():
			display[s.(state).Name()] += count
		default:
		}

	}
	return display

}

func (u *ui) checkColor() func(color pixel.RGBA) int {

	colorMap := make(map[pixel.RGBA]int)
	index := 0

	return func(color pixel.RGBA) int {

		_, ok := colorMap[color]
		if !ok {
			colorMap[color] = index
			index++
		}
		return colorMap[color]
	}

}

func (u *ui) Update() {
	u.current.Update()
}

//Draw produces the ui state as a bunch of layers
func (u *ui) Draw() []Layer {
	diff := u.current.Draw(0, 0)
	//at the moment, this just has a fresh draw, offset by the X, Y coords.
	//in the future, it should only return changed cells, and it should offset
	//those by the degree of scroll.

	onscreen := u.crop(u.current.H(), u.current.W(), diff)
	//this returns everything that fits in the view.
	cells := make([]c.Cell, len(onscreen))
	i := 0
	for _, cell := range onscreen {
		cells[i] = cell
		i++
	}

	return u.toLayers(cells)
}
func (u *ui) toLayers(cells []c.Cell) []Layer {
	//make two stacks:
	fstack := make([]Layer, 0)
	bstack := make([]Layer, 0)

	//index functions:
	checkBackground := u.checkColor()
	//index = 0
	checkForeground := u.checkColor()

	// index starts at length of colors

	for _, cell := range cells {

		index := checkForeground(cell.Foreground)
		if len(fstack) == index {
			fstack = append(fstack, Layer{cell.Foreground, make([]c.Cell, 0)})
		}
		fstack[index].content = append(fstack[index].content, cell)

		index = checkBackground(cell.Background)
		if len(bstack) == index {
			bstack = append(bstack, Layer{cell.Background, make([]c.Cell, 0)})
		}
		cell.Letter = '█'
		bstack[index].content = append(bstack[index].content, cell)
	}
	return append(bstack, fstack...)
}

type coord struct {
	X, Y int
}

func (u *ui) crop(h, w int, diff []c.Cell) map[coord]c.Cell {

	view := make(map[coord]c.Cell, 0)

	for _, cell := range diff {
		if cell.X < w && cell.Y < h {
			view[coord{cell.X, cell.Y}] = cell
		}
	}

	return view

}

//this will check what input there is, then return true if it exists
func (u *ui) Event() bool {

	mouseEvent := u.mouse.Event(u.current.H(), u.current.W(), u.current)

	switch {
	case u.keyboard.Event():
		return true
	case mouseEvent:
		return true
	case len(u.HasNew(u.notifier, u.monitoredStates)) > 0:
		return true
	default:
	}
	return false

}

type Layer struct {
	color   pixel.RGBA
	content []c.Cell
}

func (l *Layer) Color() pixel.RGBA {
	return l.color
}

func (l *Layer) Content() []c.Cell {
	return l.content
}
