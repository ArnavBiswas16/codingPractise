package main

import (
	"fmt"
)

type (
	Stack struct {
		top    *node
		length int
		name   string
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func New(name string) *Stack {
	return &Stack{nil, 0, name}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

func display(s *Stack) {

	for s.Len() > 0 {
		val := s.Pop()
		println(val.(int))
	}
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}

func main() {
	n := 20
	s := New("source")
	d := New("destination")
	h := New("helper")

	s.populate(n)
	calc(n, s, d, h)
	fmt.Println(c)

	// display(d)

}

func (s *Stack) populate(n int) {
	for i := n; i >= 1; i-- {
		s.Push(i)
	}
}

var c int

func calc(n int, source *Stack, dest *Stack, help *Stack) {
	c++
	if n == 1 {
		ring := source.Pop()
		dest.Push(ring)
		fmt.Println(ring, "from ", source.name, "to", dest.name)
		return
	}

	calc(n-1, source, help, dest)
	ring := source.Pop()
	dest.Push(ring)
	fmt.Println(ring, "from ", source.name, "to", dest.name)
	calc(n-1, help, dest, source)

}
