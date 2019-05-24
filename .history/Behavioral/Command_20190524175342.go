package behavioral

import "fmt"

// Command is a command which when executed on a person will call a given method.
type Command struct {
	person *Person
	method func()
}

func NewCommand(person *Person, method func()) Command {
	return Command{person, method}
}

func (c *Command) Execute() {
	c.method()
}

type Person struct {
	name string
	cmd  Command
}

func NewPerson(name string, cmd Command) Person {
	return Person{name, cmd}
}

func (p *Person) Talk() {
	fmt.Fprintf(outputWriter, "%s is talking.\n", p.name)
	p.cmd.Execute()
}

func (p *Person) PassOn() {
	fmt.Fprintf(outputWriter, "%s is passing on.\n", p.name)
	p.cmd.Execute()
}

func (p *Person) Gossip() {
	fmt.Fprintf(outputWriter, "%s is gossiping.\n", p.name)
	p.cmd.Execute()
}

func (p *Person) Listen() {
	fmt.Fprintf(outputWriter, "%s is listening.\n", p.name)
}