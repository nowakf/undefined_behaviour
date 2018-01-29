package events

import (
	"bytes"
	//c "cthu3/common"
	"cthu3/data"
	"fmt"
	"sync"
	"text/template"
	"time"
)

//EventSystem keeps track of events, and produces accounts of them.
type EventSystem struct {
	mu         sync.Mutex
	data       *data.TextData
	w          *world
	instants   []event
	historical []event
}

func NewEventSystem(data *data.TextData, w *world) *EventSystem {
	n := new(EventSystem)
	n.data = data
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

func (e *EventSystem) GetStats() {

}

func (e *EventSystem) Test() {
}

func (e *EventSystem) addEvent(event_url string, delay int) {
	time.Sleep(time.Duration(delay) * time.Second) //whew
	event, ok := e.data.EventsData[event_url]
	if ok {
		instance := e.instantiateEvent(&event)
		e.mu.Lock()
		defer e.mu.Unlock()
		e.instants = append(e.instants, instance)
	} else {
		fmt.Println("no event written for", event_url)
	}

}

func (e *EventSystem) instantiateEvent(input *data.DeadEvent) event {

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
