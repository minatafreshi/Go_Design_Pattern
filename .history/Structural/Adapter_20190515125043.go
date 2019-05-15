
package Structural

import (
	"fmt"
)

type Target interface {
	Execute()
}

// Adaptee defines the old interface which needs to be adapted.
type Adaptee struct {
}

func (a *Adaptee) SpecificExecute() {
	fmt.Fprint(outputWriter, "Executing SpecificExecute.")
}

type Adapter struct {
	*Adaptee
}

// Execute adapts from the new interface Execute to the old SpecificExecute.
func (a *Adapter) Execute() {
	a.SpecificExecute()
}