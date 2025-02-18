package service

import (
	"context"
	"entgo.io/contrib/entgql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"privacy-ex/internal/repository"
	"privacy-ex/pkg/auth"
	"privacy-ex/pkg/ent"
	"privacy-ex/pkg/ent/user"
	"privacy-ex/pkg/graph/gen/graphqlmodel"
	"privacy-ex/pkg/httperror"
	"strconv"
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
		Paginate(
			ctx context.Context,
			client *ent.Client,
			after *entgql.Cursor[int],
			first *int,
			before *entgql.Cursor[int],
			last *int,
			where *ent.UserWhereInput,
		) (*ent.UserConnection, error)
		Signup(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateUserInput,
		) (*graphqlmodel.AuthPayload, error)
		Login(
			ctx context.Context,
			client *ent.Client,
			email string,
			password string,
		) (*graphqlmodel.AuthPayload, error)
		RefreshToken(ctx context.Context, client *ent.Client, refreshToken string) (*graphqlmodel.AuthPayload, error)
		UpdateUser(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdateUserInput,
		) (*ent.User, error)
		DeleteUser(ctx context.Context, client *ent.Client, id int) (
			bool,
			error,
		)
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

func (s *userService) Paginate(ctx context.Context, client *ent.Client, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return s.userRepository.Paginate(ctx, client, after, first, before, last, where)
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

	jwtTokenPair, err := auth.GenerateTokenPair(created.ID, created.Role)

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

func (s *userService) Login(
	ctx context.Context,
	client *ent.Client,
	email string,
	password string,
) (*graphqlmodel.AuthPayload, error) {
	found, err := s.userRepository.FindOne(
		ctx, client, func(query *ent.UserQuery) {
			query.Where(user.EmailEQ(email))
		},
	)

	if ent.IsNotFound(err) {
		return nil, &httperror.HTTPError{
			StatusCode: http.StatusUnauthorized,
			Message:    "failed to login",
		}
	}

	if err != nil {
		return nil, err
	}

	// 입력된 비밀번호화 해시화된 거 비교
	if err := s.checkPassword(found.Password, password); err != nil {
		return nil, &httperror.HTTPError{
			StatusCode: http.StatusUnauthorized,
			Message:    "failed to login",
		}
	}

	jwtTokenPair, err := auth.GenerateTokenPair(found.ID, found.Role)

	if err != nil {
		return nil, err
	}

	return &graphqlmodel.AuthPayload{
		User:         found,
		AccessToken:  jwtTokenPair.AccessToken,
		RefreshToken: jwtTokenPair.RefreshToken,
	}, nil
}

func (s *userService) checkPassword(
	hashedPassword,
	plainPassword string,
) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(plainPassword),
	)
}

func (s *userService) RefreshToken(ctx context.Context, client *ent.Client, refreshToken string) (*graphqlmodel.AuthPayload, error) {
	token, err := auth.ValidateJwt(refreshToken, true)

	if err != nil {
		return nil, &httperror.HTTPError{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid refresh token",
		}
	}

	// Claims 가져오기
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	// 사용자 ID 확인
	userIdStr, ok := claims["sub"].(string)

	if !ok {
		return nil, errors.New("invalid token subject")
	}

	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	found, err := s.userRepository.FindOne(
		ctx, client, func(query *ent.UserQuery) {
			query.Where(user.IDEQ(userId))
		},
	)

	if err != nil {
		return nil, err
	}

	jwtTokenPair, err := auth.GenerateTokenPair(found.ID, found.Role)

	if err != nil {
		return nil, err
	}

	return &graphqlmodel.AuthPayload{
		User:         found,
		AccessToken:  jwtTokenPair.AccessToken,
		RefreshToken: jwtTokenPair.RefreshToken,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, client *ent.Client, id int, input ent.UpdateUserInput) (*ent.User, error) {
	return s.userRepository.UpdateOne(ctx, client, id, input)
}

func (s *userService) DeleteUser(ctx context.Context, client *ent.Client, id int) (bool, error) {
	var success = false

	err := s.userRepository.DeleteOne(ctx, client, id)

	if err == nil {
		success = true
	}

	return success, err
}
