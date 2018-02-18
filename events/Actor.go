package events

type Actor struct {
	Name          string
	Org           *Organization
	Abilities     map[*Action]int
	Relationships map[*Actor]int
	history       history
}

//GenerateActor is for when you want to randomly generate a person, for
//whom no content exists.
func GenerateActor(w *World) Actor {
	return Actor{
		Name:          "THX780",
		Org:           &Organization{},
		Abilities:     make(map[*Action]int),
		Relationships: make(map[*Actor]int),
	}

}

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

func (a *Actor) History() *history {
	return &a.history
}

//checks for mail - returns number of mails in mailbox
