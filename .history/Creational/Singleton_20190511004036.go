package Creational

import (
		"sync"
)

type Singleton struct {
}

var instance *Singleton
var once sync