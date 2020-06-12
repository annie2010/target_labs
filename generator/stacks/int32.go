
package stacks

import (
	"errors"
	"fmt"
	"strings"
)

// int32 represents a stack data structure.
type Int32 struct {
	data []int32
}

// Top fetch the top of the stack or zero value if empty.
func (s *Int32) Top() int32 {
	var v int32
	if s.Empty() {
		return v
	}
	return s.data[len(s.data)-1]
}

// Push adds a new element to the stack.
func (s *Int32) Push(v int32) {
	s.data = append(s.data, v)
}

// Pop removes an element from the top of that stack
func (s *Int32) Pop() (int32, error) {
	if s.Empty() {
		var v int32
		return v, errors.New("stack is empty")
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v, nil
}

// Empty returns true if stack is empty. False otherwise.
func (s *Int32) Empty() bool {
	return len(s.data) == 0
}

// String dump stack as string
func (s *Int32) String() string {
	ss := make([]string, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		ss = append(ss, fmt.Sprintf("%v", s.data[i]))
	}
	return strings.Join(ss, ",")
}
