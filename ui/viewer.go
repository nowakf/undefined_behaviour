package ui

import (
	"ub/events"
	el "ub/ui/elements"
)

type viewer struct {
	root *el.Node
	*el.Table
}

func NewViewer(h, w int) *viewer {
	v := new(viewer)
	v.root = el.Root()
	v.Table = el.NewTable(v.root, h, w)
	return v
}
func (v *viewer) Flush() {
	h, w := v.H(), v.W()
	node := v.Node
	v.Table = el.NewTable(node, h, w)
}

type name string

func (n *name) Name() string {
	return string(*n)
}

type linker struct {
	current state
	links   *map[stateEnum]state
}

func (l *linker) Next(next stateEnum) {
	if l.current != nil {
		l.current.Exit()
	}
	l.current = (*l.links)[next]
	l.current.Flush()
	l.current.Start()
}

type notifier struct {
	old     int
	counter chan int
	pipe    chan *events.Event
	source  func() int
}

func NewNotifier(source func(chan *events.Event) func() int) *notifier {
	n := new(notifier)
	n.counter = make(chan int)
	n.pipe = make(chan *events.Event)
	n.source = source(n.pipe)
	return n
}

func (n *notifier) Monitor() chan int {
	return n.counter
}
func (n *notifier) Listener() {
	delta := n.old - n.source()
	if delta > 0 {
		n.counter <- delta
		n.old += delta
	}
}
func (n *notifier) Pipe() chan *events.Event {
	return n.pipe
}
