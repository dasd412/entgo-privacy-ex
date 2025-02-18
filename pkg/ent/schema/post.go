package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/privacy"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"privacy-ex/pkg/auth"
	"time"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields defines the fields for the Post entity.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty().
			Comment("게시물 제목"),

		field.Text("content").
			Comment("게시물 내용"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("게시물 작성 시간"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("게시물 수정 시간"),
	}
}

// Edges defines the relationships of the Post entity.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("posts").
			Unique().
			Required().
			Comment("게시물 작성자와의 관계"),
	}
}

func (Post) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			auth.DenyIfNoAuthority(),
			auth.AllowIfAdminOrAuthor(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

func (Post) Annotations() []schema.Annotation {
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
