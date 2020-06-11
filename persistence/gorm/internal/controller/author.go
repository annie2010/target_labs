package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gopherland/target_labs/gorm/internal/model"
)

const defaultTimeout = 5 * time.Second

// Author represents a author resource controller.
type Author struct {
	model AuthorCruder
}

// AuthorCruder represents CRUD operations on author model.
type AuthorCruder interface {
	List(context.Context) ([]model.Author, error)
}

// NewAuthor returns a new instance.
func NewAuthor(m AuthorCruder) Author {
	return Author{model: m}
}

// Index returns a collection of authors
func (a Author) Index(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), defaultTimeout)
	defer cancel()

	authors, err := a.model.List(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	raw, err := json.Marshal(&authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(raw); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
