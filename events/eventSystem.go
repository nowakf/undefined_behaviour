package events

import (
	"bytes"
	"fmt"
	"sync"
	"text/template"
	"time"
)

type etype int

const (
	Email etype = iota
	News
)

//EventSystem keeps track of events, and produces accounts of them.
type EventSystem struct {
	mu         sync.Mutex
	w          *world
	complete   *virtual
	instants   []event
	historical []event
}

func NewEventSystem(w *world) *EventSystem {
	n := new(EventSystem)
	n.complete = newVirtual()
	n.w = w
	n.instants = n.startingInstants(n.w)
	n.historical = make([]event, 0)
	return n
}

func (e *EventSystem) startingInstants(w *world) []event {
	return make([]event, 0)
}

func (e *EventSystem) Tick() {
	e.historical = append(e.historical, e.instants...)
	//do events,

}

func (e *EventSystem) GetCurrent(t etype) []event {
	evs := make([]event, 0)
	for _, instant := range e.instants {
		if instant.etype == t {

		}
	}
	return evs
}

func (e *EventSystem) GetStats() {

}

func (e *EventSystem) Test() {
}

//adds an event to be played in the future
func (e *EventSystem) addEvent(event_url string, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	event, ok := e.complete.Events[event_url]
	if ok {
		instance := e.instantiateEvent(&event)
		e.mu.Lock()
		defer e.mu.Unlock()
		e.instants = append(e.instants, instance)
	} else {
		fmt.Println("no event written for", event_url)
	}

}

func (e *EventSystem) instantiateEvent(input *event) event {

	p, err := template.New("p").Parse(input.Content())

	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)

	err = p.Execute(buf, e.w)

	if err != nil {
		panic(err)
	}

	output := new(event)
	output.event_url = input.Event_Url()
	output.etype = input.Etype()
	output.content = buf.String()

	return *output
}

// event email, article,
// next []event, delay
// effects []action

// simply using 'string.string.string' as the key title,
// it then returns an event modified by the global variables

//Because hash functions convert keys to locations or page identifiers, the size of a hash structure is not related to the size of a key. The only effect of key size on a hash structure is that the hash function takes slightly longer to execute on a long key

// events should be in goroutines....
