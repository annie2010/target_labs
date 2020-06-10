// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0
package controller_test

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gopherland/db/internal/controller"
	"github.com/gopherland/db/internal/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAuthors(t *testing.T) {
	var (
		m    = mockAuthor{}
		c    = controller.NewAuthor(m)
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/api/v1/authors", nil)
	)

	mx := mux.NewRouter()
	mx.HandleFunc(`/api/v1/authors`, c.Index)
	mx.ServeHTTP(rr, r)

	assert.Equal(t, http.StatusOK, rr.Code)
	var aa []model.Author
	err := json.NewDecoder(rr.Body).Decode(&aa)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(aa))
}

// Helpers...

type mockAuthor struct{}

func (m mockAuthor) List(context.Context) ([]model.Author, error) {
	aa := make([]model.Author, 0, 10)
	for i := 0; i < 10; i++ {
		aa = append(aa, model.Author{
			ID:        i,
			FirstName: "Blee",
			LastName:  "Doh" + strconv.Itoa(i),
			Age:       int(rand.Int31n(80)),
		})
	}

	return aa, nil
}
