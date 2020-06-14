package internal_test

import (
	"fmt"
	"testing"

	"github.com/gopherland/target_labs/gql/internal"
	"github.com/gopherland/target_labs/gql/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateAuthor(t *testing.T) {
	uu := map[string]struct {
		id     string
		eb, ea int
	}{
		"no-exist": {id: "isan-100", eb: 5, ea: 6},
		"exist":    {id: "isan-0", eb: 5, ea: 5},
	}

	for k := range uu {
		r := internal.NewResolver()
		seed(r)
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			a, err := r.CreateAuthor(model.AuthorInput{
				ID:    u.id,
				First: "Fred",
				Last:  "Blee",
			})

			if a == nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, u.id, a.ID)
			}
			assert.Equal(t, u.ea, len(r.Authors))
		})
	}
}

func TestAllBooks(t *testing.T) {
	r := internal.NewResolver()
	r.Books["isbn-0"] = model.Book{
		ID:       "isbn-0",
		Title:    "Blee",
		Category: model.BookCategoryRomance,
	}

	bb := r.AllBooks()
	assert.Equal(t, 1, len(bb))
}

func TestAllAuthors(t *testing.T) {
	r := internal.NewResolver()
	r.Authors["isan-0"] = model.Author{
		ID:    "isan-0",
		First: "Fred",
		Last:  "Blee",
	}

	aa := r.AllAuthors()
	assert.Equal(t, 1, len(aa))
}

func TestDeleteAuthor(t *testing.T) {
	uu := map[string]struct {
		id     string
		eb, ea int
	}{
		"hit":  {id: "isan-0", eb: 4, ea: 4},
		"miss": {id: "fred", eb: 5, ea: 5},
	}

	for k := range uu {
		r := internal.NewResolver()
		seed(r)
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			a, err := r.DeleteAuthor(u.id)
			if a == nil {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, u.id, a.ID)
			}
			assert.Equal(t, u.ea, len(r.Authors))
			assert.Equal(t, u.eb, len(r.Books))
		})
	}
}

func TestDeleteBook(t *testing.T) {
	uu := map[string]struct {
		id     string
		eb, ea int
	}{
		"hit":  {id: "isbn-0", eb: 4, ea: 5},
		"miss": {id: "fred", eb: 5, ea: 5},
	}

	for k := range uu {
		r := internal.NewResolver()
		seed(r)
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			b, err := r.DeleteBook(u.id)
			if b == nil {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, u.id, b.ID)
			}
			assert.Equal(t, u.eb, len(r.Books))
			assert.Equal(t, u.ea, len(r.Authors))
		})
	}
}

func TestBooksByAuthor(t *testing.T) {
	uu := map[string]struct {
		aid, bid string
		e        int
	}{
		"hit":  {aid: "isan-0", bid: "isbn-0", e: 1},
		"miss": {aid: "fred", e: 0},
	}

	for k := range uu {
		r := internal.NewResolver()
		seed(r)
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			bb, err := r.BooksByAuthor(u.aid)
			if bb == nil {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, u.e, len(bb))
				assert.Equal(t, u.bid, bb[0].ID)
			}
		})
	}
}

func TestDeleteBooksByAuthor(t *testing.T) {
	uu := map[string]struct {
		aid, bid string
		eb, ea   int
	}{
		"hit":  {aid: "isan-0", bid: "isbn-0", eb: 4, ea: 5},
		"miss": {aid: "fred", eb: 5, ea: 5},
	}

	for k := range uu {
		r := internal.NewResolver()
		seed(r)
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			bb, err := r.DeleteBooksByAuthor(u.aid)
			if bb == nil {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, u.bid, bb[0].ID)
			}
			assert.Equal(t, u.eb, len(r.Books))
			assert.Equal(t, u.ea, len(r.Authors))
		})
	}
}

// Helpers...

func seed(r *internal.Resolver) {
	for i := 0; i < 5; i++ {
		isan := fmt.Sprintf("isan-%d", i)
		r.Authors[isan] = model.Author{
			ID:    isan,
			First: "Fred",
			Last:  fmt.Sprintf("Duh%d", i),
		}
	}
	for i := 0; i < 5; i++ {
		isbn := fmt.Sprintf("isbn-%d", i)
		r.Books[isbn] = model.Book{
			ID:       isbn,
			Title:    fmt.Sprintf("Duh%d", i),
			Category: model.BookCategoryRomance,
			Authors: []model.Author{
				r.Authors[fmt.Sprintf("isan-%d", i)],
			},
		}
	}
}
