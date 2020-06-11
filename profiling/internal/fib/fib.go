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

<<!!YOUR_CODE!!>> -- improve on the provided Fib implementation.