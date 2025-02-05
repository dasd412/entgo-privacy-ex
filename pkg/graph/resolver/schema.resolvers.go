package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"fmt"
	"privacy-ex/pkg/ent"
	"privacy-ex/pkg/graph/gen"
)

// RequestAccessToken is the resolver for the requestAccessToken field.
func (r *mutationResolver) RequestAccessToken(ctx context.Context, username string, password string) (string, error) {
	panic(fmt.Errorf("not implemented: RequestAccessToken - requestAccessToken"))
}

// VerifyAccessToken is the resolver for the verifyAccessToken field.
func (r *mutationResolver) VerifyAccessToken(ctx context.Context, token string) (bool, error) {
	panic(fmt.Errorf("not implemented: VerifyAccessToken - verifyAccessToken"))
}

// RequestRefreshToken is the resolver for the requestRefreshToken field.
func (r *mutationResolver) RequestRefreshToken(ctx context.Context, id int) (string, error) {
	panic(fmt.Errorf("not implemented: RequestRefreshToken - requestRefreshToken"))
}

// VerifyRefreshToken is the resolver for the verifyRefreshToken field.
func (r *mutationResolver) VerifyRefreshToken(ctx context.Context, token string) (bool, error) {
	panic(fmt.Errorf("not implemented: VerifyRefreshToken - verifyRefreshToken"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns gen.MutationResolver implementation.
func (r *Resolver) Mutation() gen.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
