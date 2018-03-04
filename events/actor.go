package events

type actor struct {
	name string
	org  *group
	//stats		//64_uint?
	*motivations //64_uint
	*abilities   //64_uint
	*relationships
	history history
}

func NewActor() *actor {
	a := new(actor)
	a.history = history{mail: make([]Record, 0)}
	return a
}
func (a *actor) History() *history {
	return &a.history
}

func (a *actor) Name() string { return "" }
func (a *actor) Org() string  { return "" }
func (a *actor) SetOrg()      {}

type history struct {
	mail []Record
}

func (h *history) NewMail() func() (*Record, int) {
	index := 0
	return func() (*Record, int) {
		length := len(h.mail)
		if index < length {
			return &h.mail[index], length - index
		} else {
			return nil, length - index
		}
	}
}

//checks for mail - returns number of mails in mailbox

// has to be good for random lookup,
// has to be sorted...
