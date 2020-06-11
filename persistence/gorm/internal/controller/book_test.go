package controller_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gopherland/target_labs/gorm/internal/controller"
	"github.com/gopherland/target_labs/gorm/internal/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestBooks(t *testing.T) {
	var (
		m    = mockBook{}
		c    = controller.NewBook(m)
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/api/v1/books", nil)
	)

	mx := mux.NewRouter()
	mx.HandleFunc(`/api/v1/books`, c.Index)
	mx.ServeHTTP(rr, r)
	assert.Equal(t, http.StatusOK, rr.Code)

	var bb []model.Book
	err := json.NewDecoder(rr.Body).Decode(&bb)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(bb))
}

func TestBooksByAuthor(t *testing.T) {
	var (
		m    = mockBook{}
		c    = controller.NewBook(m)
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/api/v1/books/fred", nil)
	)

	mx := mux.NewRouter()
	mx.HandleFunc(`/api/v1/books/{author}`, c.ByAuthor)
	mx.ServeHTTP(rr, r)
	assert.Equal(t, http.StatusOK, rr.Code)

	var bb []model.Book
	err := json.NewDecoder(rr.Body).Decode(&bb)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(bb))
}

// Helpers...

type mockBook struct{}

func (m mockBook) List(context.Context) ([]model.Book, error) {
	aa := make([]model.Book, 0, 10)
	for i := 0; i < 10; i++ {
		aa = append(aa, model.Book{
			Model: gorm.Model{
				ID: uint(i),
			},
			ISBN:        strconv.Itoa(i),
			Title:       "Doh" + strconv.Itoa(i),
			PublishedOn: time.Now(),
		})
	}

	return aa, nil
}

func (m mockBook) ByAuthor(context.Context, string) ([]model.Book, error) {
	aa := make([]model.Book, 0, 10)
	for i := 0; i < 1; i++ {
		aa = append(aa, model.Book{
			Model: gorm.Model{
				ID: uint(i),
			},
			ISBN:        strconv.Itoa(i),
			Title:       "Doh" + strconv.Itoa(i),
			PublishedOn: time.Now(),
		})
	}

	return aa, nil
}
