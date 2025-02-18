package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"fmt"
	"privacy-ex/pkg/ent"
	"privacy-ex/pkg/graph/gen"
	"privacy-ex/pkg/graph/gen/graphqlmodel"
)

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, input ent.CreateUserInput) (*graphqlmodel.AuthPayload, error) {
	entClient := ent.FromContext(ctx)
	return r.userService.Signup(ctx, entClient, input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*graphqlmodel.AuthPayload, error) {
	entClient := ent.FromContext(ctx)
	return r.userService.Login(ctx, entClient, email, password)
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	/*
		JWTMiddleware가 인터셉트해서 리프레시토큰을 발급함. 이 API는 지우면 안 됨.
	*/
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	entClient := ent.FromContext(ctx)
	return r.userService.UpdateUser(ctx, entClient, id, input)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (bool, error) {
	entClient := ent.FromContext(ctx)
	return r.userService.DeleteUser(ctx, entClient, id)
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput) (*ent.Post, error) {
	// 기본적으로 모든 사용자가 생성 가능
	entClient := ent.FromContext(ctx)
	return r.postService.CreatePost(ctx, entClient, input)
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input ent.UpdatePostInput) (*ent.Post, error) {
	//작성자만 수정 가능
	entClient := ent.FromContext(ctx)
	return r.postService.UpdatePost(ctx, entClient, id, input)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id int) (bool, error) {
	//관리자, 작성자만 삭제 가능
	entClient := ent.FromContext(ctx)
	return r.postService.DeletePost(ctx, entClient, id)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*ent.User, error) {
	return r.userService.FindUser(ctx, r.entClient, id)
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id int) (*ent.Post, error) {
	//  기본적으로 모든 사용자가 조회 가능
	return r.postService.FindPost(ctx, r.entClient, id)
}

// Mutation returns gen.MutationResolver implementation.
func (r *Resolver) Mutation() gen.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
