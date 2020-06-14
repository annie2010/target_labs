package internal

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gopherland/target_labs/gql/internal/generated"
	"github.com/gopherland/target_labs/gql/internal/model"
)

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.AuthorInput) (*model.Author, error) {
	panic("NYI")
}

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*model.Author, error) {
	panic("NYI")
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*model.Book, error) {
	panic("NYI")
}

func (r *mutationResolver) DeleteBooksByAuthor(ctx context.Context, id string) ([]model.Book, error) {
	panic("NYI")
}

func (r *queryResolver) AllBooks(ctx context.Context) ([]model.Book, error) {
	panic("NYI")
}

func (r *queryResolver) AllAuthors(ctx context.Context) ([]model.Author, error) {
	panic("NYI")
}

func (r *queryResolver) BooksByAuthor(ctx context.Context, id string) ([]model.Book, error) {
	panic("NYI")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
