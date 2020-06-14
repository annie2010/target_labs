// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0
package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gopherland/target_labs/sql/internal/model"
	"github.com/gorilla/mux"
)

// Book represents a book resource controller.
type Book struct {
	model BookCruder
}

// BookCruder represents available book operations.
type BookCruder interface {
	// List lists out all books.
	List(context.Context) ([]model.Book, error)

	// ByAuthor list books given an author last name.
	ByAuthor(context.Context, string) ([]model.Book, error)
}

// NewBook returns a new instance.
func NewBook(m BookCruder) Book {
	return Book{model: m}
}

// Index returns a collection of books
func (b Book) Index(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), defaultTimeout)
	defer cancel()

	books, err := b.model.List(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	raw, err := json.Marshal(&books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(raw); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ByAuthor returns a collection of books from a given author.
func (b Book) ByAuthor(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), defaultTimeout)
	defer cancel()

	params := mux.Vars(r)

	books, err := b.model.ByAuthor(ctx, params["author"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	raw, err := json.Marshal(&books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(raw); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
