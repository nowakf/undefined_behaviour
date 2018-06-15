package events

import (
	w "ub/events/world"
	//"bytes"
	//	"fmt"
	"math"
	"sync"
	//"text/template"
	"time"
)

//eventSystem keeps track of Events, and produces accounts of them.
type EventSystem struct {
	tickRate chan float64
	w        *w.World

	//this sends moods to actors
	sync.Mutex
}

func NewEventSystem(w *w.World) *EventSystem {
	e := new(EventSystem)
	e.w = w
	e.tickRate = make(chan float64)
	return e
}
func (e *EventSystem) Start() {

}

//tickrate sets the time taken before each loop of the event system
func (e *EventSystem) TickRate(newTickRate float64) {
	if newTickRate == 0 {
		e.tickRate <- math.MaxFloat64
	} else {
		e.tickRate <- 100 / newTickRate
	}
}

func (e *EventSystem) Loop(stop chan struct{}) {

	tick := time.NewTicker(time.Duration(<-e.tickRate) * time.Millisecond)
	for {
		select {
		case <-tick.C:
			e.update()
		case newTickRate := <-e.tickRate:
			tick.Stop()
			tick = time.NewTicker(time.Duration(newTickRate) * time.Millisecond)
		case <-stop:
			tick.Stop()
			break
		}
	}
	//clean up...
	//do finishing stuff here
}
func (e *EventSystem) update() {
	for _, person := range e.w.People {
		//read from graph
		person.Observe(e.w)

		//write to graph
		person.Act(e.w)
	}

}

func (e *EventSystem) instantiateActor(input *actor) (*actor, error) {
	return input, nil
}

//this should check an event against the World, and fill in the particulars.
func (e *EventSystem) instantiateRecord(input *w.Record) *w.Record {

	output := new(w.Record)

	return output
}

// Event email, article,
// next []Event, delay
// effects []action

// simply using 'string.string.string' as the key title,
// it then returns an Event modified by the global variables

//Because hash functions convert keys to locations or page identifiers, the size of a hash structure is not related to the size of a key. The only effect of key size on a hash structure is that the hash function takes slightly longer to execute on a long key

// Events should be in goroutines....
