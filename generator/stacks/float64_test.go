
package stacks

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestFloat64Stack(t *testing.T) {
	var data float64

	s := Float64{}
	s.Push(data)
	assert.Equal(t, data, s.Top())
	assert.Equal(t, fmt.Sprintf("%v", data), s.String())
	s.Pop()
	assert.Assert(t, s.Empty())
}
