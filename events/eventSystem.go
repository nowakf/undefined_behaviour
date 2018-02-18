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
	mu         sync.Mutex
	w          *World
	complete   *virtual
	Player     *Actor
	instants   []Event
	historical []Event
}

func NewEventSystem(w *World) *EventSystem {
	n := new(EventSystem)
	n.complete = newVirtual()
	n.w = w

	player, exists := n.complete.Actors["player"]
	if !exists {
		player = GenerateActor(w)
		println("no default player in virtual set")
		println(player.Name)
	}

	//p, err := n.instantiateActor(&player)

	//if err != nil {
	//	panic(err)
	//}
	n.Player = &player

	println(n.Player.Name)

	n.instants = n.startingInstants(n.w)
	n.historical = make([]Event, 0)

	return n
}

func (e *EventSystem) startingInstants(w *World) []Event {
	return make([]Event, 0)
}

func (e *EventSystem) Tick() []Event {
	e.historical = append(e.historical, e.instants...)
	//do Events,
	return e.instants

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

type ActorCreationError struct {
	cause string
}

func (a *ActorCreationError) Error() string {
	return a.cause
}

func (e *EventSystem) instantiateActor(input *Actor) (*Actor, error) {
	return input, nil
}

//this should check an event against the World, and fill in the particulars.
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
