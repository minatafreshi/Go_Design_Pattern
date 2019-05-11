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
		fmt.Fprintf(outputWriter, "I am ")
}