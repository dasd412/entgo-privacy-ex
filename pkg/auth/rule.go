package auth

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

func DenyIfNoAuthority() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		role := UserAuthorityFromContext(ctx)

		if role == nil {
			return privacy.Denyf("no role found in context")
		}
		return privacy.Skip
	})
}

func AllowIfSignupOrLogin() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, mutation ent.Mutation) error {

		return nil
	})
}

func AllowIfAdminOrAuthor() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		role := UserAuthorityFromContext(ctx)

		if role.IsAdmin() || role.IsAuthor() {
			return privacy.Allow
		}

		return privacy.Deny
	})
}
