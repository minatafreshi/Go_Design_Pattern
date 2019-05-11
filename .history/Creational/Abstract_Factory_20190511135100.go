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

type rectangle struct {
}

func (r *rectangle) Draw() {
		fmt.Fprintf(outputWriter, "I am rectanle")
}

type ShapeFactory interface {
	CreateCurvelShape() Shape
	CreateStrightShape() Shape
}

type simpleShapeFactory struct {
}

func NewSimpleShapeFactory() ShapeFactory {
	return &simpleShapeFactory{}
}

func (s *simpleShapeFactory) create

