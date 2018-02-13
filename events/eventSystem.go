package events

import (
	//"bytes"
	"fmt"
	"sync"
	//"text/template"
	"time"
)

//EventSystem keeps track of Events, and produces accounts of them.
type EventSystem struct {
	w          *world
	complete   *virtual
	instants   []Event
	historical []Event
	mailPipe   chan (Event)
}

func NewEventSystem(w *world) *EventSystem {
	n := new(EventSystem)
	n.complete = newVirtual()
	n.w = w
	//n.mailPipe = make(chan (Event)) //just so it doesn't freak when there's no mailsystem
	println(load()[0].Email.Content)
	if load()[0].Article == (Article{}) {
		println("we have no article")
	} else {
		println("we have something here")
	}
	n.instants = n.startingInstants(n.w)
	n.historical = make([]Event, 0)

	return n
}

func (e *EventSystem) startingInstants(w *world) []Event {
	return make([]Event, 0)
}

func (e *EventSystem) Tick() []Event {
	e.historical = append(e.historical, e.instants...)
	//do Events,
	return e.instants

}

func (e *EventSystem) MailHookup(mails chan (Event)) {
	e.mailPipe = mails
}

func (e *EventSystem) GetCurrent() []Event {
	return *new([]Event)
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

	output := new(Event)

	return *output
}

// Event email, article,
// next []Event, delay
// effects []action

// simply using 'string.string.string' as the key title,
// it then returns an Event modified by the global variables

//Because hash functions convert keys to locations or page identifiers, the size of a hash structure is not related to the size of a key. The only effect of key size on a hash structure is that the hash function takes slightly longer to execute on a long key

// Events should be in goroutines....
