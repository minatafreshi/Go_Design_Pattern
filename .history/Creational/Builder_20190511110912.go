package Creational

import (
		"fmt"
)

type MyMustang struct {
	name string
	fuelType string
	cost int
	power int
}

func (m MyMustang) showMustang() {
	fmt.Println("Mustang Name: ")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}