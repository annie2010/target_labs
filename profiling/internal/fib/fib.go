// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package fib

// Compute computes a Fibonacci number.
func Compute(n int) int {
	if n < 2 {
		return n
	}

	return Compute(n-2) + Compute(n-1)
}

// ComputeIter computes Fibonacci number in line.
func ComputeIter(n int) int {
	if n < 2 {
		return n
	}
	p1, p2 := 0, 1
	for i := 0; i < n-1; i++ {
		p1, p2 = p2, p1+p2
	}
	return p2
}
