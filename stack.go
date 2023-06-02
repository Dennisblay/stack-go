package main

import "fmt"

// Stack represents a stack data structure
type Stack struct {
	items []interface{}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek returns the top element from the stack without removing it
func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}

	return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

func main___() {
	stack := Stack{}

	// Pushing elements onto the stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Checking if the stack is empty
	fmt.Println("Is empty:", stack.IsEmpty())

	// Getting the top element
	fmt.Println("Top element:", stack.Peek())

	// Popping elements from the stack
	fmt.Println("Popped:", stack.Pop())
	fmt.Println("Popped:", stack.Pop())

	// Getting the updated top element
	fmt.Println("Top element:", stack.Peek())

	// Checking the size of the stack
	fmt.Println("Size:", stack.Size())
}
