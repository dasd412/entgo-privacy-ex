package auth

import (
	"context"
	"entgo.io/ent"
	"privacy-ex/pkg/ent/privacy"
)

func AllowMutateIfAdmin() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, mutation ent.Mutation) error {
		return nil
	})
}

func AllowMutateIfOwner() privacy.MutationRule {
	return nil
}

func DenyQueryIfNotAdmin() privacy.QueryRule {
	return privacy.QueryRuleFunc(func(ctx context.Context, query ent.Query) error {
		return nil
	})
}
