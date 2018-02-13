package events

import "strings"

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
	Consequences []string `yaml:"consequences"`
	//this is the events consequences, including 'magic events' that effect game state...
}

func (e Event) Subject() string {
	subject := ""
	for i := 0; i <= e.Depth; i++ {
		subject += "Re:"
	}
	return subject + e.Title
}

func (e Event) Headline() string {
	return strings.ToUpper(e.Title)
}

func (e Event) Sender() string {
	return e.Instigator.Name + "@" + e.Instigator.Org.EmailRoot
}
func (e Event) Body() string {
	return e.Content
}
