package Behavioral

import "fmt"

type Element interface {
	Accept(v Visitor)
}

type This struct {
}

func (t *This) This() string {
	return "This"
}

func (t *This) Accept(v Visitor) {
	v.VisitThis(t)
}

type That struct {
}

func (t *That) That() string {
	return "That"
}

func (t *That) Accept(v Visitor) {
	v.VisitThat(t)
}

type TheOther struct {
}

func (t *TheOther) TheOther() string {
	return "TheOther"
}

func (t *TheOther) Accept(v Visitor) {
	v.VisitTheOther(t)
}

type Visitor interface {
	VisitThis(e *This)
	VisitThat(e *That)
	VisitTheOther(e *TheOther)
}

// UpVisitor defines an up visitor.
type UpVisitor struct {
}

// VisitThis visits this.
func (v *UpVisitor) VisitThis(e *This) {
	fmt.Printf("do Up on %v\n", e.This())
}

// VisitThat visits that.
func (v *UpVisitor) VisitThat(e *That) {
	fmt.Printf("do Up on %v\n", e.That())
}

// VisitTheOther visits the other.
func (v *UpVisitor) VisitTheOther(e *TheOther) {
	fmt.Printf("do Up on %v\n", e.TheOther())
}

// DownVisitor defines a down visitor.
type DownVisitor struct {
}

// VisitThis visits this.
func (v *DownVisitor) VisitThis(e *This) {
	fmt.Printf("do Down on %v\n", e.This())
}

// VisitThat visits that.
func (v *DownVisitor) VisitThat(e *That) {
	fmt.Printf("do Down on %v\n", e.That())
}

// VisitTheOther visits the other.
func (v *DownVisitor) VisitTheOther(e *TheOther) {
	fmt.Printf("do Down on %v\n", e.TheOther())
}
