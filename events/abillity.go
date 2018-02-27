package events

import "math/rand"

type ability struct {
	title        string
	requirements map[*ability]int
	tags         []string //this corresponds to |categories| in the eventHeader type
	variance     int
}

//Checks if the action can be done.
func (a *ability) Attempt(o *actor, t *actor) (bool, string) {
	if len(a.requirements) == 0 {
		return true, a.title + a.doneString()
	}
	for skill, req := range a.requirements {

		roll := rand.Intn(a.variance)

		oskill, ok := o.abilities.normal[skill]

		if !ok {
			oskill = -10
		}
		tskill, ok := t.abilities.normal[skill]

		if !ok {
			tskill = -10
		}
		if oskill+roll-req-tskill < 0 {
			return false, a.title + a.failString(oskill, tskill, req, roll)
		}
	}
	return true, a.title + a.doneString()
}

func (a *ability) GetTags() []string {
	return a.tags
}

func (a *ability) doneString() string {
	return " done!"
}

func (a *ability) failString(oskill, tskill, req, roll int) string {
	return " failed!"
}
