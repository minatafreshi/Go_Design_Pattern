package Structural

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type ITask interface {
	Execute(taskType string)
}

type Task struct {
	taskName string
}

func (t *Task) Execute(taskType string) {
	fmt.Fprint(outputWriter, "Performing task type: "+taskType)
}

// ProxyTask represents a proxy task with re-routes tasks.
type ProxyTask struct {
	task *Task
}

// NewProxyTask creates a new instance of a ProxyTask.
func NewProxyTask() *ProxyTask {
	return &ProxyTask{task: &Task{}}
}

func (t *ProxyTask) Execute(taskType string) {
	if taskType == "Run" {
		t.task.Execute(taskType)
	}
}