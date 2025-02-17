package auth

import (
	"context"
	"entgo.io/ent"
	"privacy-ex/pkg/ent/privacy"
)

func AllowIfAdmin() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, mutation ent.Mutation) error {
		return nil
	})
}

func AllowIfOwner() privacy.MutationRule {
	return nil
}
