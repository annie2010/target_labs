package internal_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopherland/labs2/profiling/internal"
	"github.com/stretchr/testify/assert"
)

func TestFibHandler(t *testing.T) {
	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/fib?n=20", nil)
	)

	internal.FibHandler(rr, r)
	assert.Equal(t, http.StatusOK, rr.Code)

	var res internal.Results
	err := json.NewDecoder(rr.Body).Decode(&res)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(res))
	for i, f := range []int{0, 1, 1, 2} {
		assert.Equal(t, f, res[i].Fibonacci)
	}
}

func TestFibHandlerIter(t *testing.T) {
	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/fib?n=20", nil)
	)

	internal.FibHandlerIter(rr, r)
	assert.Equal(t, http.StatusOK, rr.Code)

	var res internal.Results
	err := json.NewDecoder(rr.Body).Decode(&res)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(res))
	for i, f := range []int{0, 1, 1, 2} {
		assert.Equal(t, f, res[i].Fibonacci)
	}
}

func BenchmarkFibHandler(b *testing.B) {
	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/fib?n=20", nil)
	)

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		internal.FibHandler(rr, r)
	}
}

func BenchmarkFibHandlerIter(b *testing.B) {
	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/fib?n=20", nil)
	)

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		internal.FibHandlerIter(rr, r)
	}
}
