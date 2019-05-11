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
	mustang = MyMustang{}
	mustang.name = "Shelby GT500"
	mustang.fuelType = "gas"
	mustang.cost = 30000
	mustang.power = 700
}

func (a ShelbyCobra427Builder) creatMustang() MyMustang {
	mustang = MyMustang{}
	mustang.name = "Shelby Cobra427"
	mustang.fuelType = "gas"
	mustang.cost = 750000
	mustang.power = 427
}
type ProductionLine struct {
	builder MustangBuilder
}

func (a ProductionLine)CreateMustang() MyMustang {
	return a.builder.createMustang()
}

func main() {
	p := ProductionLine{Builder{}}
	res := p.CreateMyMustang()
	res.showMustang()
	fmt.Println("=============================================================================")
	p = ProductionLine{ABuilder{}}
	res = p.CreateMyMustang()
	res.showAudi()

}