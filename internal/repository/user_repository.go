package repository

import (
	"context"
	"privacy-ex/pkg/ent"
)

type (
	userRepository struct {
	}

	UserRepository interface {
		FindOne(
			ctx context.Context,
			client *ent.Client,
			opts ...func(query *ent.UserQuery),
		) (*ent.User, error)
		CreateOne(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateUserInput,
		) (*ent.User, error)
	}
)

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u userRepository) FindOne(
	ctx context.Context,
	client *ent.Client,
	opts ...func(query *ent.UserQuery),
) (*ent.User, error) {
	query := client.User.
		Query()

	for _, opt := range opts {
		opt(query)
	}

	return query.Only(ctx)
}

func (u userRepository) CreateOne(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateUserInput,
) (*ent.User, error) {
	return client.User.
		Create().
		SetInput(input).
		Save(ctx)
}
