package Creational

import (
		"testing"
		"github.com/stretchr/testify/assert"
)

func TestGetInstance_ReturnsSingleton(t *testing.T) {
	s	ingleton := GetInstance()
	assert.NotNil(t, singleton)
}