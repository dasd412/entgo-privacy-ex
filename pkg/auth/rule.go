package auth

import (
	"context"
	"entgo.io/ent"
	"privacy-ex/pkg/ent/privacy"
)

func DenyIfNoAuthority() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		return nil
	})
}

func AllowIfAdminOrAuthor() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, mutation ent.Mutation) error {
		return nil
	})
}
