package Behavioral

type Memento struct {
	state int
}

func NewMemento(value int) *Memento {
	return &Memento{value}
}

type Number struct {
	value int
}

func NewNumber(value int) *Number {
	return &Number{value}
}

func (n *Number) Dubble() {
	n.value = 2 * n.value
}

func (n *Number) Half() {
	n.value = n.value / 2
}

func (n *Number) Value() int {
	return n.value
}

func (n *Number) CreateMemento() *Memento {
	return NewMemento(n.value)
}

// ReinstateMemento reinstates the value of the Number to the value of the memento.
func (n *Number) ReinstateMemento(memento *Memento) {
	n.value = memento.state
}
