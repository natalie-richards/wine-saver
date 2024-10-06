package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"

	"github.com/natalie-richards/wine-app/graph/model"
)

// AddBookmark is the resolver for the addBookmark field.
func (r *mutationResolver) AddBookmark(ctx context.Context, input model.AddBookmarkRequest) (*model.Bookmark, error) {
	err := r.App.AddBookmark(ctx, &input)
	if err != nil {
		return nil, err
	}
	bookmark := model.Bookmark(input)
	return &bookmark, nil
}

// ListBookmarks is the resolver for the listBookmarks field.
func (r *queryResolver) ListBookmarks(ctx context.Context) ([]*model.Bookmark, error) {
	bookmarks, err := r.App.GetBookmarks(ctx)
	if err != nil {
		return nil, err
	}
	return bookmarks, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
