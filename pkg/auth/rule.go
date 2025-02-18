package auth

import (
	"context"
	"entgo.io/ent"
	"privacy-ex/pkg/ent/privacy"
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

func DenyIfGuest() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		role := UserAuthorityFromContext(ctx)

		if role.IsGuest() {
			return privacy.Deny
		}

		return privacy.Skip
	})
}

func AllowIfAdminOrAuthor() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, mutation ent.Mutation) error {
		role := UserAuthorityFromContext(ctx)

		if role.IsAdmin() || role.IsAuthor() {
			return privacy.Allow
		}

		return privacy.Deny
	})
}
