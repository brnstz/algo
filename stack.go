package algo

import (
	"errors"
)

var EmptyStack = errors.New("stack is empty")

// A LIFO stack implementation using linked lists.
type Stack struct {
	tail *item
	size int
}

// Private structure to maintain linked list of items
type item struct {
	value interface{}
	prev  *item
}

func NewStack() *Stack {

	// Return a stack with a nil pointer for tail, and size = 0
	return &Stack{}
}

func (s *Stack) Push(val interface{}) {
	// Create a new item for the tail with current val, and a next pointer
	// to the current tail value (possibly nil)
	it := &item{value: val, prev: s.tail}

	// Overwrite current tail with our new item
	s.tail = it

	// Increment size of stack
	s.size++

	return
}

func (s *Stack) Pop() (interface{}, error) {
	// Return nil value and empty error if stack is empty
	if s.IsEmpty() {
		return nil, EmptyStack
	}

	// Otherwise, actually pop from the list
	it := s.tail

	// Set new tail to prev value
	s.tail = it.prev

	// Decrement size of stack
	s.size--

	return it.value, nil
}

// Peek on stack without popping.
func (s *Stack) Peek() interface{} {
	return s.tail.value
}

// How many items are on the stack?
func (s *Stack) Size() int {
	return s.size
}

// Is this stack empty?
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// Implement AddDel interface for test code
func (s *Stack) Add(val interface{}) {
	s.Push(val)
}
func (s *Stack) Del() (interface{}, error) {
	return s.Pop()
}
