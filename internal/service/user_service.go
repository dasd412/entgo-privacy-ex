package service

import (
	"context"
	"privacy-ex/pkg/ent"
	"privacy-ex/pkg/ent/user"
)

type (
	userService struct {
	}

	UserService interface {
		FindUser(
			ctx context.Context,
			client *ent.Client,
			id int,
		) (*ent.User, error)
		Signup(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateUserInput,
		) (*ent.User, error)
		SignIn(
			ctx context.Context,
			client *ent.Client,
			username string,
			password string,
		) (string, error)
	}
)

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) FindUser(
	ctx context.Context,
	client *ent.Client,
	id int,
) (*ent.User, error) {
	return client.User.Query().
		Where(user.ID(id)).
		Only(ctx)
}

func (s *userService) Signup(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateUserInput,
) (*ent.User, error) {
	return client.User.
		Create().
		SetInput(input).
		Save(ctx)
}

func (s *userService) SignIn(
	ctx context.Context,
	client *ent.Client,
	username string,
	password string,
) (string, error) {
	return "", nil
}
