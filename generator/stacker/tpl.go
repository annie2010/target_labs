package main

func goFile() string {
	return `
package {{ .Package }}

import (
	"errors"
	"fmt"
	"strings"
)

// {{ .Type }} represents a stack data structure.
type {{ asType .Type }} struct {
	data []{{ .Type }}
}

// Top fetch the top of the stack or zero value if empty.
func (s *{{ asType .Type }}) Top() {{ .Type }} {
	var v {{ .Type }}
	if s.Empty() {
		return v
	}
	return s.data[len(s.data)-1]
}

// Push adds a new element to the stack.
func (s *{{ asType .Type }}) Push(v {{ .Type }}) {
	s.data = append(s.data, v)
}

// Pop removes an element from the top of that stack
func (s *{{ asType .Type }}) Pop() ({{ .Type }}, error) {
	if s.Empty() {
		var v {{ .Type }}
		return v, errors.New("stack is empty")
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v, nil
}

// Empty returns true if stack is empty. False otherwise.
func (s *{{ asType .Type }}) Empty() bool {
	return len(s.data) == 0
}

// String dump stack as string
func (s *{{ asType .Type }}) String() string {
	ss := make([]string, 0, len(s.data))
	for i := len(s.data) - 1; i >= 0; i-- {
		ss = append(ss, fmt.Sprintf("%v", s.data[i]))
	}
	return strings.Join(ss, ",")
}
`
}

func goTestFile() string {
	return `
package {{ .Package }}

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func Test{{ asType .Type}}Stack(t *testing.T) {
	var data {{ .Type }}

	s := {{ asType .Type }}{}
	s.Push(data)
	assert.Equal(t, data, s.Top())
	assert.Equal(t, fmt.Sprintf("%v", data), s.String())
	s.Pop()
	assert.Assert(t, s.Empty())
}
`
}
