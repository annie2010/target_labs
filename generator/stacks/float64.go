
package stacks

import (
	"errors"
	"fmt"
	"strings"
)

// float64 represents a stack data structure.
type Float64 struct {
	data []float64
}

// Top fetch the top of the stack or zero value if empty.
func (s *Float64) Top() float64 {
	var v float64
	if s.Empty() {
		return v
	}
	return s.data[len(s.data)-1]
}

// Push adds a new element to the stack.
func (s *Float64) Push(v float64) {
	s.data = append(s.data, v)
}

// Pop removes an element from the top of that stack
func (s *Float64) Pop() (float64, error) {
	if s.Empty() {
		var v float64
		return v, errors.New("stack is empty")
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v, nil
}

// Empty returns true if stack is empty. False otherwise.
func (s *Float64) Empty() bool {
	return len(s.data) == 0
}

// String dump stack as string
func (s *Float64) String() string {
	ss := make([]string, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		ss = append(ss, fmt.Sprintf("%v", s.data[i]))
	}
	return strings.Join(ss, ",")
}
