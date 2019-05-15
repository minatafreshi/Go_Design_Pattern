package Creational

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type StoogeType int

const (
	Larry StoogeType = iota
	Moe
	Curly
)

type Stooge interface {
	SlapStick()
}

type larry struct {
}

func (s *larry) SlapStick() {
	fmt.Fprint(outputWriter, "Larry: Poke eyes\n")
}

type moe struct {
}

func (s *moe) SlapStick() {
	fmt.Fprint(outputWriter, "Moe: Slap head\n")
}

type curly struct {
}

func (s *curly) SlapStick() {
	fmt.Fprint(outputWriter, "Curly: Suffer abuse\n")
}

func NewStooge(stooge StoogeType) Stooge {
	if stooge == Larry {
		return &larry{}
	} else if stooge == Moe {
		return &moe{}
	} else if stooge == Curly {
		return &curly{}
	}
	return nil
}