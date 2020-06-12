// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package grep_test

import (
	"testing"

	"github.com/gopherland/target_labs/grep"
	"github.com/stretchr/testify/assert"
)

func TestWordCountV1(t *testing.T) {
	<<!!YOUR_CODE!!>> -- Test V1. Use a table test with subtests and make sure your implementation is valid
}

func TestWordCountV2(t *testing.T) {
	<<!!YOUR_CODE!!>> -- Test V2. Use a table test with subtests and make sure your implementation is valid
}

func BenchmarkWordCountV1(b *testing.B) {
	<<!!YOUR_CODE!!>> -- Implement the v1 benchmark

}

func BenchmarkWordCountV2(b *testing.B) {
	<<!!YOUR_CODE!!>> -- Implement the v2 benchmark
}

// Helpers...

func genSamples(t testing.TB) []string {
	t.Helper()

	return []string{
		"*** START OF THIS PROJECT GUTENBERG EBOOK MOBY DICK; OR THE WHALE ***",
		"MOBY-DICK;",
		`“Moby Dick?” shouted Ahab. “Do ye know the white whale then, Tash?”`,
		"seen—Moby Dick—Moby Dick!”",
	}
}
