package events

type actor struct {
	name string
	org  *group
	*motivations
	*abilities
	*relationships
	history history
}

func NewActor() *actor {
	a := new(actor)
	a.history = history{mail: make([]Event, 0)}
	return a
}
func (a *actor) History() *history {
	return &a.history
}

func (a *actor) Name() string { return "" }
func (a *actor) Org() string  { return "" }
func (a *actor) SetOrg()      {}

type history struct {
	mail []Event
}

func (h *history) NewMail() func() (*Event, int) {
	index := 0
	return func() (*Event, int) {
		length := len(h.mail)
		if index < length {
			return &h.mail[index], length - index
		} else {
			return nil, length - index
		}
	}
}

//checks for mail - returns number of mails in mailbox
