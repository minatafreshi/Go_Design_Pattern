package Creational 

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type circle struct {
}

func (c *circle)