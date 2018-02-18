package events

import "math/rand"

type Action struct {
	Title        string
	Requirements map[*Action]int
	Variance     int
}

//Checks if the action can be done.
func (a *Action) Attempt(o *Actor, t *Actor) (bool, string) {
	if len(a.Requirements) == 0 {
		return true, a.Title + a.doneString()
	}
	for skill, req := range a.Requirements {

		roll := rand.Intn(a.Variance)

		oskill, ok := o.Abilities[skill]

		if !ok {
			oskill = -10
		}
		tskill, ok := t.Abilities[skill]

		if !ok {
			tskill = -10
		}
		if oskill+roll-req-tskill < 0 {
			return false, a.Title + a.failString(oskill, tskill, req, roll)
		}
	}
	return true, a.Title + a.doneString()
}

func (a *Action) Apply(o *Actor, t *Actor, w *World) {
}

func (a *Action) doneString() string {
	return " done!"
}

func (a *Action) failString(oskill, tskill, req, roll int) string {
	return " failed!"
}
