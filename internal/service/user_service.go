package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"privacy-ex/internal/repository"
	"privacy-ex/pkg/auth"
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
	hashedPassword, err := s.hashPassword(input.Password)

	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	input.Password = hashedPassword

	created, err := s.userRepository.CreateOne(ctx, client, input)

	if err != nil {
		return nil, err
	}

	jwtTokenPair, err := auth.GenerateTokenPair(created.ID)

	if err != nil {
		return nil, err
	}

	return &graphqlmodel.AuthPayload{
		User:         created,
		AccessToken:  jwtTokenPair.AccessToken,
		RefreshToken: jwtTokenPair.RefreshToken,
	}, nil
}

func (s *userService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *userService) SignIn(
	ctx context.Context,
	client *ent.Client,
	username string,
	password string,
) (*graphqlmodel.AuthPayload, error) {
	return nil, nil
}
