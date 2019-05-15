package Structural

import "fmt"

type Component interface {
	Traverse()
}

type Leaf struct {
	value int
}

func NewLeaf(value int) *Leaf {
	return &Leaf{value}
}

func (l *Leaf) Traverse() {
	fmt.Printf("%v  ", l.value)
}

type Composite struct {
	children []Component
}

func NewComposite() *Composite {
	return &Composite{make([]Component, 0)}
}

// Add adds a new component to the composite.
func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

// Traverse traverses the composites children.
func (c *Composite) Traverse() {
	for i := 0; i < len(c.children); i++ {
		c.children[i].Traverse()
	}
}