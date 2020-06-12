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
	samples := genSamples(t)

	uu := map[string]struct {
		text string
		e    int64
	}{
		"semi-cols": {
			text: samples[0],
			e:    1,
		},
		"dash": {
			text: samples[1],
			e:    1,
		},
		"quotes": {
			text: samples[2],
			e:    1,
		},
		"special-dash": {
			text: samples[3],
			e:    2,
		},
	}

	t.Parallel()
	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, grep.WordCount("moby", u.text))
		})
	}
}

func TestWordCountV2(t *testing.T) {
	samples := genSamples(t)

	uu := map[string]struct {
		text string
		e    int64
	}{
		"semi-cols": {
			text: samples[0],
			e:    1,
		},
		"dash": {
			text: samples[1],
			e:    1,
		},
		"quotes": {
			text: samples[2],
			e:    1,
		},
		"special-dash": {
			text: samples[3],
			e:    2,
		},
	}

	t.Parallel()
	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, grep.WordCountBytes("moby", u.text))
		})
	}
}

func BenchmarkWordCountV1(b *testing.B) {
	const sample = `“Moby Dick?” shouted Ahab. “Do ye know the white whale then, Tash?”`

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		grep.WordCount("moby", sample)
	}
}

func BenchmarkWordCountV2(b *testing.B) {
	sample := genSamples(b)[2]

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		grep.WordCountBytes("moby", sample)
	}
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
