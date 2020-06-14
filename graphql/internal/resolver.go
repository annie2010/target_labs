package internal

import (
	"fmt"

	"github.com/gopherland/target_labs/gql/internal/generated/model"
)

// Books represents a collection of books.
type Books map[string]model.Book

// Authors represents a collection of authors.
type Authors map[string]model.Author

// Resolver represents a schema resolver.
type Resolver struct {
	Books   map[string]model.Book
	Authors map[string]model.Author
}

// NewResolver returns a new instance.
func NewResolver() *Resolver {
	return &Resolver{
		Books:   make(Books),
		Authors: make(Authors),
	}
}

// CreateAuthor creates a new author.
func (r *Resolver) CreateAuthor(f model.AuthorInput) (*model.Author, error) {
	if _, ok := r.Authors[f.ID]; ok {
		return nil, fmt.Errorf("Meow! Author `%s already exists", f.ID)
	}

	a := model.Author(f)
	r.Authors[f.ID] = a

	return &a, nil
}

// AllBooks returns all available books.
func (r *Resolver) AllBooks() []model.Book {
	bb := make([]model.Book, 0, len(r.Books))
	for _, b := range r.Books {
		bb = append(bb, b)
	}

	return bb
}

// AllAuthors returns all available authors.
func (r *Resolver) AllAuthors() []model.Author {
	aa := make([]model.Author, 0, len(r.Authors))
	for _, a := range r.Authors {
		aa = append(aa, a)
	}

	return aa
}

// DeleteAuthor nukes an author by ISAN and all related books.
// Returns an error is the author does not exists.
func (r *Resolver) DeleteAuthor(id string) (*model.Author, error) {
	if a, ok := r.Authors[id]; ok {
		delete(r.Authors, id)
		_, _ = r.DeleteBooksByAuthor(id)
		return &a, nil
	}

	return nil, fmt.Errorf("Meow! Author `%s not found", id)
}

// BooksByAuthor find all books given an author ISAN.
// Returns an error if the author does not exists.
func (r *Resolver) BooksByAuthor(id string) ([]model.Book, error) {
	if _, ok := r.Authors[id]; !ok {
		return nil, fmt.Errorf("Meow! Author `%s not found", id)
	}

	bb := make([]model.Book, 0, len(r.Books))
	for _, b := range r.Books {
		for _, a := range b.Authors {
			if a.ID == id {
				bb = append(bb, b)
			}
		}
	}

	return bb, nil
}

// DeleteBooksByAuthors deletes all books from a given author ISAN.
func (r *Resolver) DeleteBooksByAuthor(id string) ([]model.Book, error) {
	victims := make([]model.Book, 0, 1)
	for k, b := range r.Books {
		if hasAuthor(b.Authors, id) {
			victims = append(victims, b)
			delete(r.Books, k)
		}
	}

	if len(victims) > 0 {
		return victims, nil
	}

	return nil, fmt.Errorf("Meow! No books found for author `%s", id)
}

// DeleteBook deletes a book given an ISBN.
func (r *Resolver) DeleteBook(id string) (*model.Book, error) {
	if b, ok := r.Books[id]; ok {
		delete(r.Books, id)
		return &b, nil
	}

	return nil, fmt.Errorf("Meow! Book `%s not found", id)
}

// Helpers...

func hasAuthor(aa []model.Author, id string) bool {
	for _, a := range aa {
		if a.ID == id {
			return true
		}
	}
	return false
}
