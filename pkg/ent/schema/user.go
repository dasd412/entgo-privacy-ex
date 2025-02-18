package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/privacy"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"privacy-ex/pkg/auth"
	"time"
)

type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Comment("이메일").
			Unique(),
		field.String("password").
			Comment("해시화된 비밀 번호"),
		field.String("name").
			Comment("이름"),
		field.Time("created_at").
			Comment("생성 날짜").
			Default(time.Now).
			Immutable(),
		field.Enum("role").
			Values("admin", "author", "guest").
			Comment("인가 권한"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).
			Unique().
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
	}
}

func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// 회원 가입, 로그인 같은 경우는 전부 풀어놔야 함.
			// 다른 경우에는 전부 사용자 이상만 권한이 있어야 함.
			auth.AllowIfSignupOrLogin(),
			auth.AllowIfAdminOrAuthor(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			auth.AllowIfSignupOrLogin(),
			auth.AllowIfAdminOrAuthor(),
			privacy.AlwaysDenyRule(),
		},
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.MultiOrder(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
