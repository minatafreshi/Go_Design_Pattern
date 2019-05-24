package Behavioral

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type Event struct {
	id string
}

type EventObserver interface {
	OnNotify(event Event)
}

type observer struct {
	name string
}

func NewEventObserver(name string) EventObserver {
	return &observer{name}
}

func (o *observer) OnNotify(event Event) {
	fmt.Fprintf(outputWriter, "observer '%s' received event '%s'\n", o.name, event.id)
}

type EventNotifier interface {
	Register(obs EventObserver)
	Deregister(obs EventObserver)
	Notify(event Event)
}

type eventNotifer struct {
	observers []EventObserver
}

// NewEventNotifier returns a new instance of an EventNotifier.
func NewEventNotifier() EventNotifier {
	return &eventNotifer{}
}

// Register registers a new observer for notifying on.
func (e *eventNotifer) Register(obs EventObserver) {
	e.observers = append(e.observers, obs)
}

// Deregister de-registers an observer for notifying on.
func (e *eventNotifer) Deregister(obs EventObserver) {
	for i := 0; i < len(e.observers); i++ {
		if obs == e.observers[i] {
			e.observers = append(e.observers[:i], e.observers[i+1:]...)
		}
	}
}

// Notify notifies all observers on an event.
func (e *eventNotifer) Notify(event Event) {
	for i := 0; i < len(e.observers); i++ {
		e.observers[i].OnNotify(event)
	}
}
