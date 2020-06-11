package controller_test

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gopherland/target_labs/gorm/internal/controller"
	"github.com/gopherland/target_labs/gorm/internal/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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
			Model: gorm.Model{
				ID: uint(i),
			},
			FirstName: "Blee",
			LastName:  "Doh" + strconv.Itoa(i),
			Age:       int(rand.Int31n(80)),
		})
	}

	return aa, nil
}
