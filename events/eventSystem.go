package events

import (
	"bytes"
	"fmt"
	"sync"
	"text/template"
	"time"
)

type etype string

const (
	email etype = "email"
	news        = "news"
)

//EventSystem keeps track of Events, and produces accounts of them.
type EventSystem struct {
	mu         sync.Mutex
	w          *world
	complete   *virtual
	instants   []Event
	historical []Event
}

func NewEventSystem(w *world) *EventSystem {
	n := new(EventSystem)
	n.complete = newVirtual()
	n.w = w
	println(load()[0].Content)
	println(load()[1].Content)
	n.instants = n.startingInstants(n.w)
	n.historical = make([]Event, 0)

	return n
}

func (e *EventSystem) startingInstants(w *world) []Event {
	return make([]Event, 0)
}

func (e *EventSystem) Tick() {
	e.historical = append(e.historical, e.instants...)
	//do Events,

}

func (e *EventSystem) GetCurrent(t etype) []Event {
	evs := make([]Event, 0)
	for _, instant := range e.instants {
		if instant.Etype == t {

		}
	}
	return evs
}

func (e *EventSystem) GetStats() {

}

func (e *EventSystem) Test() {
}

//adds an Event to be played in the future
func (e *EventSystem) addEvent(Event_url string, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	Event, ok := e.complete.Events[Event_url]
	if ok {
		instance := e.instantiateEvent(&Event)
		e.mu.Lock()
		defer e.mu.Unlock()
		e.instants = append(e.instants, instance)
	} else {
		fmt.Println("no Event written for", Event_url)
	}

}

func (e *EventSystem) instantiateEvent(input *Event) Event {

	p, err := template.New("p").Parse(input.Content)

	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)

	err = p.Execute(buf, e.w)

	if err != nil {
		panic(err)
	}

	output := new(Event)
	output.Event_url = input.Event_url
	output.Etype = input.Etype
	output.Content = buf.String()

	return *output
}

// Event email, article,
// next []Event, delay
// effects []action

// simply using 'string.string.string' as the key title,
// it then returns an Event modified by the global variables

//Because hash functions convert keys to locations or page identifiers, the size of a hash structure is not related to the size of a key. The only effect of key size on a hash structure is that the hash function takes slightly longer to execute on a long key

// Events should be in goroutines....
