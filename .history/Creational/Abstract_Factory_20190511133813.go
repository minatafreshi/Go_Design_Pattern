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
		fmt.Fprintf(outputWriter, "I am square")
}

type ellipse struct {
}

func (e *ellipse) Draw() {
		fmt.Fprintf(outputWriter, "I am ellipse")
}

type rectangel struct {
}

func (r *rectangel) Draw() {
		fmt.Fprintf(outputWriter, "I am re")
}