package main

// Stack represents a stack data structure.
type Stack struct {
	items []interface{}
}

// Push adds a new element to the top of the stack.
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack.
func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	topIndex := len(s.items) - 1
	topItem := s.items[topIndex]
	s.items = s.items[:topIndex]
	return topItem
}

// Peek returns the top element from the stack without removing it.
func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return len(s.items)
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
