package service

import (
	"context"
	"privacy-ex/pkg/ent"
)

type (
	postService struct {
	}

	PostService interface {
		FindPost(
			ctx context.Context,
			client *ent.Client,
			id int,
		) (*ent.Post, error)
		CreatePost(
			ctx context.Context,
			client *ent.Client,
			input ent.CreatePostInput,
		) (*ent.Post, error)
		UpdatePost(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdatePostInput,
		) (*ent.Post, error)
		DeletePost(ctx context.Context, client *ent.Client, id int) (
			bool,
			error,
		)
	}
)

func NewPostService() PostService {
	return &postService{}
}
func (p postService) FindPost(
	ctx context.Context,
	client *ent.Client,
	id int,
) (*ent.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p postService) CreatePost(
	ctx context.Context,
	client *ent.Client,
	input ent.CreatePostInput,
) (*ent.Post, error) {
	return client.Post.
		Create().
		SetInput(input).
		Save(ctx)
}

func (p postService) UpdatePost(
	ctx context.Context,
	client *ent.Client,
	id int,
	input ent.UpdatePostInput,
) (*ent.Post, error) {
	return client.Post.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (p postService) DeletePost(
	ctx context.Context,
	client *ent.Client,
	id int,
) (bool, error) {
	var success = false

	err := client.Post.
		DeleteOneID(id).
		Exec(ctx)

	if err == nil {
		success = true
	}

	return success, err
}
