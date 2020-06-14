// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package fib_test

import (
	"testing"

	"github.com/gopherland/labs2/profiling/internal/fib"
	"github.com/stretchr/testify/assert"
)

var uu = []struct {
	n, e int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
}

func TestCompute(t *testing.T) {
	for _, u := range uu {
		assert.Equal(t, u.e, fib.Compute(u.n))
	}
}

<<!!YOUR_CODE!!>> -- Ensure your new implementation produces the same results!


func BenchmarkCompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib.Compute(20)
	}
}

<<!!YOUR_CODE!!>> -- Benchmark your new implementation!
