package main

import "fmt"

func main() {
	s := Stack{}
	for n := 0; n < 20; n++ {
		s.Push(n)
	}
	for n := 0; n < 9; n++ {
		fmt.Println(s.Pop())
	}
	//s.Push(a)

	fmt.Println(s.items)
}

type Stack struct {
	items []interface{}
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}
func (s *Stack) Pop() interface{} {
	if s.isEmpty() {
		return nil
	}
	lastItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return lastItem
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) TopOrPeek() interface{} {
	lastItem := s.items[len(s.items)-1]
	return lastItem
}
