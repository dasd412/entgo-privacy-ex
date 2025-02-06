package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"privacy-ex/internal/repository"
	"privacy-ex/internal/service"
	"privacy-ex/pkg/ent"
	"privacy-ex/pkg/graph/gen"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	entClient   *ent.Client
	userService service.UserService
	postService service.PostService
}

func NewSchema(entClient *ent.Client) graphql.ExecutableSchema {
	return gen.NewExecutableSchema(
		gen.Config{
			Resolvers: &Resolver{
				entClient:   entClient,
				userService: service.NewUserService(repository.NewUserRepository()),
				postService: service.NewPostService(repository.NewPostRepository()),
			},
		},
	)
}
