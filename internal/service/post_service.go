package service

import (
	"context"
	"privacy-ex/internal/repository"
	"privacy-ex/pkg/ent"
	"privacy-ex/pkg/ent/post"
)

type (
	postService struct {
		postRepository repository.PostRepository
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

func NewPostService(postRepository repository.PostRepository) PostService {
	return &postService{
		postRepository: postRepository,
	}
}
func (p *postService) FindPost(
	ctx context.Context,
	client *ent.Client,
	id int,
) (*ent.Post, error) {
	return p.postRepository.FindOne(
		ctx, client, func(query *ent.PostQuery) {
			query.Where(post.ID(id))
		},
	)
}

func (p *postService) CreatePost(
	ctx context.Context,
	client *ent.Client,
	input ent.CreatePostInput,
) (*ent.Post, error) {
	return p.postRepository.CreateOne(
		ctx, client, input,
	)
}

func (p *postService) UpdatePost(
	ctx context.Context,
	client *ent.Client,
	id int,
	input ent.UpdatePostInput,
) (*ent.Post, error) {
	return p.postRepository.UpdateOne(
		ctx, client, id, input,
	)
}

func (p *postService) DeletePost(
	ctx context.Context,
	client *ent.Client,
	id int,
) (bool, error) {
	var success = false

	err := p.postRepository.DeleteOne(ctx, client, id)

	if err == nil {
		success = true
	}

	return success, err
}
