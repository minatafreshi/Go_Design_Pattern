package Behavioral

import "fmt"

type Strategy interface {
	Execute()
}

type strategyA struct {
}

func NewStrategyA() Strategy {
	return &strategyA{}
}

func (s *strategyA) Execute() {
	fmt.Fprintf(outputWriter, "executing strategy A\n")
}

type strategyB struct {
}

func NewStrategyB() Strategy {
	return &strategyB{}
}

func (s *strategyB) Execute() {
	fmt.Fprintf(outputWriter, "executing strategy B\n")
}

type Context struct {
	strategy Strategy
}

func NewContext() *Context {
	return &Context{}
}

// SetStrategy sets the strategy to execute for this context.
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// Execute executes the strategy.
func (c *Context) Execute() {
	c.strategy.Execute()
}
