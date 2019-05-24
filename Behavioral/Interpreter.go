package Behavioral

import (
	"strings"
)

type Expression interface {
	Interpret(variables map[string]Expression) int
}

type Integer struct {
	integer int
}

func (n *Integer) Interpret(variables map[string]Expression) int {
	return n.integer
}

type Plus struct {
	leftOperand  Expression
	rightOperand Expression
}

func (p *Plus) Interpret(variables map[string]Expression) int {
	return p.leftOperand.Interpret(variables) + p.rightOperand.Interpret(variables)
}

type Minus struct {
	leftOperand  Expression
	rightOperand Expression
}

func (m *Minus) Interpret(variables map[string]Expression) int {
	return m.leftOperand.Interpret(variables) - m.rightOperand.Interpret(variables)
}

type Variable struct {
	name string
}

func (v *Variable) Interpret(variables map[string]Expression) int {
	value, found := variables[v.name]
	if found == false {
		return 0
	}
	return value.Interpret(variables)
}

type Evaluator struct {
	syntaxTree Expression
}

func NewEvaluator(expression string) *Evaluator {
	expressionStack := new(Stack)
	for _, token := range strings.Split(expression, " ") {
		if token == "+" {
			right := expressionStack.Pop().(Expression)
			left := expressionStack.Pop().(Expression)
			subExpression := &Plus{left, right}
			expressionStack.Push(subExpression)
		} else if token == "-" {
			right := expressionStack.Pop().(Expression)
			left := expressionStack.Pop().(Expression)
			subExpression := &Minus{left, right}
			expressionStack.Push(subExpression)
		} else {
			expressionStack.Push(&Variable{token})
		}
	}
	syntaxTree := expressionStack.Pop().(Expression)
	return &Evaluator{syntaxTree}
}

func (e *Evaluator) Interpret(context map[string]Expression) int {
	return e.syntaxTree.Interpret(context)
}

type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Push(value interface{}) {
	s.top = &Node{value, s.top}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value
}