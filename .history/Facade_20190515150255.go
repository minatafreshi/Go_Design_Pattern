package Structural

import (
	"fmt"
)

type CarModel struct {
}


func NewCarModel() *CarModel {
	return &CarModel{}
}


func (c *CarModel) SetModel() {
	fmt.Fprintf(outputWriter, " CarModel - SetModel\n")
}


type CarEngine struct {
}


func NewCarEngine() *CarEngine {
	return &CarEngine{}
}


func (c *CarEngine) SetEngine() {
	fmt.Fprintf(outputWriter, " CarEngine - SetEngine\n")
}


type CarBody struct {
}


func NewCarBody() *CarBody {
	return &CarBody{}
}


func (c *CarBody) SetBody() {
	fmt.Fprintf(outputWriter, " CarBody - SetBody\n")
}

// CarAccessories is the fourth car subsystem which describes the car accessories.
type CarAccessories struct {
}

// NewCarAccessories creates new car accessories.
func NewCarAccessories() *CarAccessories {
	return &CarAccessories{}
}

// SetAccessories sets the car accessories and logs.
func (c *CarAccessories) SetAccessories() {
	fmt.Fprintf(outputWriter, " CarAccessories - SetAccessories\n")
}

// CarFacade describes the car facade which provides a simplified interface to create a car.
type CarFacade struct {
	accessories *CarAccessories
	body        *CarBody
	engine      *CarEngine
	model       *CarModel
}

// NewCarFacade creates a new CarFacade.
func NewCarFacade() *CarFacade {
	return &CarFacade{
		accessories: NewCarAccessories(),
		body:        NewCarBody(),
		engine:      NewCarEngine(),
		model:       NewCarModel(),
	}
}

// CreateCompleteCar creates a new complete car.
func (c *CarFacade) CreateCompleteCar() {
	fmt.Fprintf(outputWriter, "******** Creating a Car **********\n")
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetBody()
	c.accessories.SetAccessories()
	fmt.Fprintf(outputWriter, "******** Car creation is completed. **********\n")
}