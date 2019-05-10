package Creational

type Prototype interface {
	Position() string 
	Clone() Prototype
}

type ConcretePrototype struct {

	position string
}

func (p *ConcretePrototype) Position() string {
	return p.position
}

func (p *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{p.position}
}