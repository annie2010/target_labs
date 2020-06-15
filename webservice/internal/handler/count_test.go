// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopherland/target_labs/webservice/internal/handler"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCountHandler(t *testing.T) {
	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/v1/wc/3lpigs/pig", nil)
	)

	h := handler.NewBook()
	mx := mux.NewRouter()
	mx.HandleFunc(`/v1/wc/{book:[\w]+}/{word:[\w]+}`, h.Count)
	mx.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Code)
	var resp handler.Response
	err := json.NewDecoder(rr.Body).Decode(&resp)
	assert.Nil(t, err)
	assert.Equal(t, int64(26), resp.Occurrences)
	assert.Equal(t, "3lpigs", resp.Book)
	assert.Equal(t, "pig", resp.Word)
}

func BenchmarkCountHandler(b *testing.B) {
	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/v1/grep/3lpigs/pig", nil)
		h    = handler.NewBook()
	)

	mx := mux.NewRouter()
	mx.HandleFunc(`/v1/grep/{book:[\w]+}/{word:[\w]+}`, h.Count)

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		mx.ServeHTTP(rr, r)
	}
}
