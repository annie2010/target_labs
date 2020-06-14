package internal

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gopherland/target_labs/gql/internal/generated"
	model1 "github.com/gopherland/target_labs/gql/internal/generated/model"
)

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model1.AuthorInput) (*model1.Author, error) {
	return r.Resolver.CreateAuthor(input)
}

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*model1.Author, error) {
	return r.Resolver.DeleteAuthor(id)
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*model1.Book, error) {
	return r.Resolver.DeleteBook(id)
}

func (r *mutationResolver) DeleteBooksByAuthor(ctx context.Context, id string) ([]model1.Book, error) {
	return r.Resolver.DeleteBooksByAuthor(id)
}

func (r *queryResolver) AllBooks(ctx context.Context) ([]model1.Book, error) {
	return r.Resolver.AllBooks(), nil
}

func (r *queryResolver) AllAuthors(ctx context.Context) ([]model1.Author, error) {
	return r.Resolver.AllAuthors(), nil
}

func (r *queryResolver) BooksByAuthor(ctx context.Context, id string) ([]model1.Book, error) {
	return r.Resolver.BooksByAuthor(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
