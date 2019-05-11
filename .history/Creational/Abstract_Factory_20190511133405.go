package Creational 

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type circle struct {
}

func (c *circle) Draw() {
		fmt.Fprintf(outputWriter, "I am circle")
}

type square struct {
}

func (s *square) Draw() {
		fmt.fpr
}