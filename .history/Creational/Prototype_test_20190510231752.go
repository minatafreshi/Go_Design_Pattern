package Creational

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPosition_ReturnPosition(t *testing.T) {
	expectedPosition := "Prototype instance"

	proto := ConcretePrototype{expectedPosition}
	actualPosition := proto.Position()

	assert.Equal(t, expectedPosition, actualPosition)
}

func TestClone_ReturnsNil(t *testing.T) {
	position := "Prototype instance"

	proto := ConcretePrototype{position}
	newProto := proto.Clone()

	assert.NotNil(t, newProto)
}

func TestClone_ReturnDifferentInstance(t *testing.T) {
	position := "Proto instance"

	proto := ConcretePrototype{position}
	newProto:= proto.Clone()
	
	assert.NotEqual(t, proto, newProto)
}

func TestPosition_WhenPrototypeIsCloned_ReturnsPosition(t *testing.T) {
	expectedPosition := "prototype instance"

	proto := ConcretePrototype{expectedPosition}
	newProto := proto.Clone()
	actualPosition := newProto.Position()

	assert.Equal(t, expectedPosition, actualPosition)
}