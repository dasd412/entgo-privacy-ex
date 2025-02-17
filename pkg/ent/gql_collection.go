// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"privacy-ex/pkg/ent/post"
	"privacy-ex/pkg/ent/user"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (po *PostQuery) CollectFields(ctx context.Context, satisfies ...string) (*PostQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return po, nil
	}
	if err := po.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return po, nil
}

func (po *PostQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(post.Columns))
		selectedFields = []string{post.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "author":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: po.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			po.withAuthor = query
			if _, ok := fieldSeen[post.FieldAuthorID]; !ok {
				selectedFields = append(selectedFields, post.FieldAuthorID)
				fieldSeen[post.FieldAuthorID] = struct{}{}
			}
		case "title":
			if _, ok := fieldSeen[post.FieldTitle]; !ok {
				selectedFields = append(selectedFields, post.FieldTitle)
				fieldSeen[post.FieldTitle] = struct{}{}
			}
		case "content":
			if _, ok := fieldSeen[post.FieldContent]; !ok {
				selectedFields = append(selectedFields, post.FieldContent)
				fieldSeen[post.FieldContent] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[post.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, post.FieldCreatedAt)
				fieldSeen[post.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[post.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, post.FieldUpdatedAt)
				fieldSeen[post.FieldUpdatedAt] = struct{}{}
			}
		case "authorID":
			if _, ok := fieldSeen[post.FieldAuthorID]; !ok {
				selectedFields = append(selectedFields, post.FieldAuthorID)
				fieldSeen[post.FieldAuthorID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		po.Select(selectedFields...)
	}
	return nil
}

type postPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PostPaginateOption
}

func newPostPaginateArgs(rv map[string]any) *postPaginateArgs {
	args := &postPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*PostWhereInput); ok {
		args.opts = append(args.opts, WithPostFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "posts":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PostClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, postImplementors)...); err != nil {
				return err
			}
			u.withPosts = query
		case "email":
			if _, ok := fieldSeen[user.FieldEmail]; !ok {
				selectedFields = append(selectedFields, user.FieldEmail)
				fieldSeen[user.FieldEmail] = struct{}{}
			}
		case "password":
			if _, ok := fieldSeen[user.FieldPassword]; !ok {
				selectedFields = append(selectedFields, user.FieldPassword)
				fieldSeen[user.FieldPassword] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[user.FieldName]; !ok {
				selectedFields = append(selectedFields, user.FieldName)
				fieldSeen[user.FieldName] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[user.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldCreatedAt)
				fieldSeen[user.FieldCreatedAt] = struct{}{}
			}
		case "role":
			if _, ok := fieldSeen[user.FieldRole]; !ok {
				selectedFields = append(selectedFields, user.FieldRole)
				fieldSeen[user.FieldRole] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
