package Behavioral

import (
	"fmt"
)

type Machine struct {
	current State
}

func NewMachine() *Machine {
	fmt.Fprintf(outputWriter, "Machine is ready.\n")
	return &Machine{NewOFF()}
}

func (m *Machine) setCurrent(s State) {
	m.current = s
}

func (m *Machine) On() {
	m.current.On(m)
}

func (m *Machine) Off() {
	m.current.Off(m)
}

type State interface {
	On(m *Machine)
	Off(m *Machine)
}

type ON struct {
}

func NewON() State {
	return &ON{}
}

func (o *ON) On(m *Machine) {
	fmt.Fprintf(outputWriter, "   already ON\n")
}

func (o *ON) Off(m *Machine) {
	fmt.Fprintf(outputWriter, "   going from ON to OFF\n")
	m.setCurrent(NewOFF())
}

// OFF describes the off button state.
type OFF struct {
}

// NewOFF creates a new OFF state.
func NewOFF() State {
	return &OFF{}
}

// On switches the state from off to on.
func (o *OFF) On(m *Machine) {
	fmt.Fprintf(outputWriter, "   going from OFF to ON\n")
	m.setCurrent(NewON())
}

// Off does nothing.
func (o *OFF) Off(m *Machine) {
	fmt.Fprintf(outputWriter, "   already OFF\n")
}
