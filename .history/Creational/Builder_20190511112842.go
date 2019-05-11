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
	fmt.Println("Mustang Name:", m.name)
	fmt.Println("Mustang Fuel:", m.fuelType)
	fmt.Println("Mustang Cost:", m.cost)
	fmt.Println("Mustang Power:", m.power)
}

type MustangBuilder interface {
	creatMustang() MyMustang
}

type ShelbyGT500Builder struct {

}

func (a ShelbyGT500Builder) creatMustang() MyMustang {
	mustang := MyMustang{}
	mustang.name = "Shelby"
	mustang.fuelType = "gas"
	mustang.cost = 30000
	mustang.power = 700
}