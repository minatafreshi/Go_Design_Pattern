package Creational

import (
		"testing"
		"github.com/stretchr/testify/assert"
)

func TestGetInstance_ReturnsSingleton(t *testing.T) {
	singleton := GetInstance()
	assert.NotNil(t, singleton)
}

func TestGetInstance_MultipleCallsToGetInstance_ReturnsSingleton(t *testing.T) {
	singleton := GetInstance
}