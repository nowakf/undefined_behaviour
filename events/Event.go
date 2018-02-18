package events

//Event is a catch-all for any interaction between Actors
type Event struct {
	Event_url string `yaml:"event_url"`
	//this is the unique identifier of the event
	Instigator *Actor `yaml:"instigator"`
	//this is the doer
	Target *Actor `yaml:"target"`
	//this is the target - it can be the player,
	//or an institution
	Depth int `yaml:"depth"`
	//this is the depth from the entry point to the
	//event chain
	Title string `yaml:"title"`
	//this is the title - for newspapers, Events, etc
	Content string `yaml:"content"`
	//this is the event's textual content
	Options []Action `yaml:"options"`
	//this is the events options, each of which
	//holds a pointer to an event
	Consequences []Action `yaml:"consequences"`
	//this is the events consequences
}

func (e *Event) GetDepth() int {
	return e.Depth
}

func (e *Event) GetTitle() string {
	return e.Title
}

func (e *Event) GetInstigator() *Actor {
	return e.Instigator
}
func (e *Event) GetBody() string {
	return e.Content
}

func (e *Event) GetOptions() []*Action {
	return make([]*Action, 0)
}
