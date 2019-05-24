package Behavioral

import (
	"fmt"
)

type WorkerInterface interface {
	GetUp()
	EatBreakfast()
	GoToWork()
	Work()
	ReturnHome()
	Relax()
	Sleep()
}

type Worker struct {
	WorkerInterface
}

func NewWorker(w WorkerInterface) *Worker {
	return &Worker{w}
}

func (w *Worker) DailyRoutine() {
	w.GetUp()
	w.EatBreakfast()
	w.GoToWork()
	w.Work()
	w.ReturnHome()
	w.Relax()
	w.Sleep()
}

// PostMan is a worker.
type PostMan struct {
}

// GetUp prints what the postman does to get up.
func (w *PostMan) GetUp() {
	fmt.Fprintf(outputWriter, "Getting up\n")
}

// EatBreakfast prints what the postman does to eat breakfast.
func (w *PostMan) EatBreakfast() {
	fmt.Fprintf(outputWriter, "Eating pop tarts\n")
}

// GoToWork prints what the postman does to get to work.
func (w *PostMan) GoToWork() {
	fmt.Fprintf(outputWriter, "Cycle to work\n")
}

// Work prints what the postman does to work.
func (w *PostMan) Work() {
	fmt.Fprintf(outputWriter, "Post letters\n")
}

// ReturnHome prints what the postman does to get home.
func (w *PostMan) ReturnHome() {
	fmt.Fprintf(outputWriter, "Cycle home\n")
}

// Relax prints what the postman does to relax.
func (w *PostMan) Relax() {
	fmt.Fprintf(outputWriter, "Collect stamps\n")
}

// Sleep prints what the postman does to sleep.
func (w *PostMan) Sleep() {
	fmt.Fprintf(outputWriter, "Zzzzzzz\n")
}
