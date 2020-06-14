package internal_test

import (
	"context"
	"testing"

	"github.com/gopherland/target_labs/gql/internal"
	"github.com/gopherland/target_labs/gql/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestResolverCreateAuthor(t *testing.T) {
	uu := map[string]struct {
		id     string
		eb, ea int
	}{
		"no-exist": {id: "isan-100", eb: 5, ea: 6},
		"exist":    {id: "isan-0", eb: 5, ea: 5},
	}

	ctx := context.Background()
	for k := range uu {
		r := internal.NewResolver()
		seed(r)
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			a, err := r.Mutation().CreateAuthor(ctx, model.AuthorInput{
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

func TestResolverAllBooks(t *testing.T) {
	r := internal.NewResolver()
	r.Books["isbn-0"] = model.Book{
		ID:       "isbn-0",
		Title:    "Blee",
		Category: model.BookCategoryRomance,
	}

	ctx := context.Background()
	bb, err := r.Query().AllBooks(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(bb))
}

func TestResolverAllAuthors(t *testing.T) {
	r := internal.NewResolver()
	r.Authors["isan-0"] = model.Author{
		ID:    "isan-0",
		First: "Fred",
		Last:  "Blee",
	}

	ctx := context.Background()
	aa, err := r.Query().AllAuthors(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(aa))
}

func TestResolverDeleteAuthor(t *testing.T) {
	uu := map[string]struct {
		id     string
		eb, ea int
	}{
		"hit":  {id: "isan-0", eb: 4, ea: 4},
		"miss": {id: "fred", eb: 5, ea: 5},
	}

	ctx := context.Background()
	for k := range uu {
		r := internal.NewResolver()
		seed(r)
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			a, err := r.Mutation().DeleteAuthor(ctx, u.id)
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

func TestResolverDeleteBook(t *testing.T) {
	uu := map[string]struct {
		id     string
		eb, ea int
	}{
		"hit":  {id: "isbn-0", eb: 4, ea: 5},
		"miss": {id: "fred", eb: 5, ea: 5},
	}

	ctx := context.Background()
	for k := range uu {
		r := internal.NewResolver()
		seed(r)
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			b, err := r.Mutation().DeleteBook(ctx, u.id)
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

func TestResolverBookByAuthor(t *testing.T) {
	uu := map[string]struct {
		aid, bid string
		e        int
	}{
		"hit":  {aid: "isan-0", bid: "isbn-0", e: 1},
		"miss": {aid: "fred", e: 0},
	}

	ctx := context.Background()
	for k := range uu {
		r := internal.NewResolver()
		seed(r)
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			bb, err := r.Query().BooksByAuthor(ctx, u.aid)
			if bb == nil {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, u.e, len(bb))
				assert.Equal(t, u.bid, bb[0].ID)
			}
		})
	}
}

func TestResolverDeleteBookByAuthor(t *testing.T) {
	uu := map[string]struct {
		aid, bid string
		eb, ea   int
	}{
		"hit":  {aid: "isan-0", bid: "isbn-0", eb: 4, ea: 5},
		"miss": {aid: "fred", eb: 5, ea: 5},
	}

	ctx := context.Background()
	for k := range uu {
		r := internal.NewResolver()
		seed(r)
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			bb, err := r.Mutation().DeleteBooksByAuthor(ctx, u.aid)
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
