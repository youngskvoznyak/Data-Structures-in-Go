package main

import "fmt"

type Stack struct {
	items []interface{}
}

// Push will add a value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end and return the removed value
func (s *Stack) Pop() interface{} {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func useStack() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(10)
	myStack.Push(1000)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

}
