package service

import (
	"context"
	"privacy-ex/internal/repository"
	"privacy-ex/pkg/ent"
	"privacy-ex/pkg/ent/user"
	"privacy-ex/pkg/graph/gen/graphqlmodel"
)

type (
	userService struct {
		userRepository repository.UserRepository
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
		) (*graphqlmodel.AuthPayload, error)
		SignIn(
			ctx context.Context,
			client *ent.Client,
			username string,
			password string,
		) (*graphqlmodel.AuthPayload, error)
	}
)

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) FindUser(
	ctx context.Context,
	client *ent.Client,
	id int,
) (*ent.User, error) {
	return s.userRepository.FindOne(
		ctx, client, func(query *ent.UserQuery) {
			query.Where(user.ID(id))
		},
	)
}

func (s *userService) Signup(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateUserInput,
) (*graphqlmodel.AuthPayload, error) {
	return nil, nil
}

func (s *userService) SignIn(
	ctx context.Context,
	client *ent.Client,
	username string,
	password string,
) (*graphqlmodel.AuthPayload, error) {
	return nil, nil
}
