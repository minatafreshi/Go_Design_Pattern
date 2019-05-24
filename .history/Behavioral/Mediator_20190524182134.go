package Behavioral

import "fmt"

type WildStallion interface {
	SetMediator(mediator Mediator)
}

type Bill struct {
	mediator Mediator
}

func (b *Bill) SetMediator(mediator Mediator) {
	b.mediator = mediator
}

func (b *Bill) Respond() {
	fmt.Fprintf(outputWriter, "Bill: What?\n")
	b.mediator.Communicate("Bill")
}

type Ted struct {
	mediator Mediator
}

func (t *Ted) SetMediator(mediator Mediator) {
	t.mediator = mediator
}

func (t *Ted) Talk() {
	fmt.Fprintf(outputWriter, "Ted: Bill?\n")
	t.mediator.Communicate("Ted")
}

func (t *Ted) Respond() {
	fmt.Fprintf(outputWriter, "Ted: Strange things are afoot at the Circle K.\n")
}

type Mediator interface {
	Communicate(who string)
}

// ConcreateMediator describes a mediator between Bill and Ted.
type ConcreateMediator struct {
	Bill
	Ted
}

// NewMediator creates a new ConcreateMediator.
func NewMediator() *ConcreateMediator {
	mediator := &ConcreateMediator{}
	mediator.Bill.SetMediator(mediator)
	mediator.Ted.SetMediator(mediator)
	return mediator
}

// Communicate communicates between Bill and Ted.
func (m *ConcreateMediator) Communicate(who string) {
	if who == "Ted" {
		m.Bill.Respond()
	} else if who == "Bill" {
		m.Ted.Respond()
	}
}
