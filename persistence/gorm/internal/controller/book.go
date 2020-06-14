package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gopherland/target_labs/gorm/internal/model"
	"github.com/gorilla/mux"
)

// Book represents a book resource controller.
type Book struct {
	model BookCruder
}

// BookCruder represents a curdable book.
type BookCruder interface {
	// List returns all books.
	List(context.Context) ([]model.Book, error)

	// ByAuthor returns books by a given author last name.
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
