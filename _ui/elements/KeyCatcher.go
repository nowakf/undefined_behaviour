package elements

import (
	"github.com/faiface/pixel/pixelgl"
)

//this interface covers anything that captures keypresses -
//only one of these should be registering input events at
//any given time
type KeyCatcher interface {
	Escape() KeyCatcher
	DoIfValid(typed string, win *pixelgl.Window) (KeyCatcher, bool)
}
type typer struct {
	parent   KeyCatcher
	selected bool
	toType   func(string)
}

var escapeKeys = []pixelgl.Button{
	pixelgl.KeyEscape,
}

func newTyper(parent KeyCatcher, toType func(string)) *typer {
	t := new(typer)
	t.parent = parent
	t.toType = toType
	return t
}

func (t *typer) DoIfValid(typed string, win *pixelgl.Window) (KeyCatcher, bool) {
	switch {
	case t.escapeBox(win):
		return t.Escape(), true
	case typed != "":
		t.toType(typed)
		return t, true
	default:
		return t, true
	}
}
func (t *typer) Escape() KeyCatcher {
	return t.parent
}

func (t *typer) escapeBox(win *pixelgl.Window) bool {
	for _, key := range escapeKeys {
		if win.JustReleased(key) {
			return true
		}
	}
	return false
}

type keyChecker struct {
	parent    KeyCatcher
	buttons   []pixelgl.Button
	validKeys map[pixelgl.Button]func()
}

func GlobalKeyChecker() *keyChecker {
	k := new(keyChecker)
	k.buttons = make([]pixelgl.Button, 0)
	k.validKeys = make(map[pixelgl.Button]func())
	return k
}
func NewKeyChecker(parent KeyCatcher) *keyChecker {
	k := new(keyChecker)
	k.parent = parent
	k.buttons = make([]pixelgl.Button, 0)
	k.validKeys = make(map[pixelgl.Button]func())
	return k
}
func (k *keyChecker) Escape() KeyCatcher {
	return k.parent
}

func (k *keyChecker) Add(key pixelgl.Button, does func()) {
	k.buttons = append(k.buttons, key)
	_, ok := k.validKeys[key]
	if ok {
		panic("key already defined")
	}
	k.validKeys[key] = does

}

func (k *keyChecker) DoIfValid(typed string, win *pixelgl.Window) (KeyCatcher, bool) {
	switch {
	case k.keyPress(win):
		return k, true
	case k.Escape() != nil:
		return k.Escape().DoIfValid(typed, win)
	default:
		return k, false
	}
}

func (k *keyChecker) keyPress(win *pixelgl.Window) bool {
	for _, button := range k.buttons {
		if win.JustReleased(button) {
			funct, ok := k.validKeys[button]
			if ok {
				funct()
				return true
			} else {
				panic("button is part of set but not part of validkeys")
			}
		}
	}
	return false
}
