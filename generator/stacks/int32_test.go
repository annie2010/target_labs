
package stacks

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestInt32Stack(t *testing.T) {
	var data int32

	s := Int32{}
	s.Push(data)
	assert.Equal(t, data, s.Top())
	assert.Equal(t, fmt.Sprintf("%v", data), s.String())
	s.Pop()
	assert.Assert(t, s.Empty())
}
