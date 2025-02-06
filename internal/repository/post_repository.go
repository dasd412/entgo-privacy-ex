package repository

import (
	"context"
	"privacy-ex/pkg/ent"
)

type (
	postRepository struct {
	}

	PostRepository interface {
		FindOne(
			ctx context.Context,
			client *ent.Client,
			opts ...func(query *ent.PostQuery),
		) (*ent.Post, error)
		CreateOne(
			ctx context.Context,
			client *ent.Client,
			input ent.CreatePostInput,
		) (*ent.Post, error)
		UpdateOne(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdatePostInput,
		) (*ent.Post, error)
		DeleteOne(ctx context.Context, client *ent.Client, id int) error
	}
)

func NewPostRepository() PostRepository {
	return &postRepository{}
}
func (p *postRepository) FindOne(
	ctx context.Context,
	client *ent.Client,
	opts ...func(query *ent.PostQuery),
) (*ent.Post, error) {
	query := client.Post.
		Query()

	for _, opt := range opts {
		opt(query)
	}

	return query.Only(ctx)
}

func (p *postRepository) CreateOne(
	ctx context.Context,
	client *ent.Client,
	input ent.CreatePostInput,
) (*ent.Post, error) {
	return client.Post.
		Create().
		SetInput(input).
		Save(ctx)
}

func (p *postRepository) UpdateOne(
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

func (p *postRepository) DeleteOne(
	ctx context.Context,
	client *ent.Client,
	id int,
) error {
	return client.Post.
		DeleteOneID(id).
		Exec(ctx)
}
