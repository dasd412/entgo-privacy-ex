// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"privacy-ex/pkg/ent/post"
	"privacy-ex/pkg/ent/predicate"
	"privacy-ex/pkg/ent/user"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePost = "Post"
	TypeUser = "User"
)

// PostMutation represents an operation that mutates the Post nodes in the graph.
type PostMutation struct {
	config
	op            Op
	typ           string
	id            *int
	title         *string
	content       *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	author        *int
	clearedauthor bool
	done          bool
	oldValue      func(context.Context) (*Post, error)
	predicates    []predicate.Post
}

var _ ent.Mutation = (*PostMutation)(nil)

// postOption allows management of the mutation configuration using functional options.
type postOption func(*PostMutation)

// newPostMutation creates new mutation for the Post entity.
func newPostMutation(c config, op Op, opts ...postOption) *PostMutation {
	m := &PostMutation{
		config:        c,
		op:            op,
		typ:           TypePost,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPostID sets the ID field of the mutation.
func withPostID(id int) postOption {
	return func(m *PostMutation) {
		var (
			err   error
			once  sync.Once
			value *Post
		)
		m.oldValue = func(ctx context.Context) (*Post, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Post.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPost sets the old Post of the mutation.
func withPost(node *Post) postOption {
	return func(m *PostMutation) {
		m.oldValue = func(context.Context) (*Post, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PostMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an httperror otherwise.
func (m PostMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PostMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PostMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Post.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *PostMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *PostMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An httperror is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *PostMutation) ResetTitle() {
	m.title = nil
}

// SetContent sets the "content" field.
func (m *PostMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *PostMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An httperror is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *PostMutation) ResetContent() {
	m.content = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *PostMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *PostMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An httperror is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *PostMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *PostMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *PostMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An httperror is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *PostMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetAuthorID sets the "author_id" field.
func (m *PostMutation) SetAuthorID(i int) {
	m.author = &i
}

// AuthorID returns the value of the "author_id" field in the mutation.
func (m *PostMutation) AuthorID() (r int, exists bool) {
	v := m.author
	if v == nil {
		return
	}
	return *v, true
}

// OldAuthorID returns the old "author_id" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An httperror is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldAuthorID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAuthorID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAuthorID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAuthorID: %w", err)
	}
	return oldValue.AuthorID, nil
}

// ResetAuthorID resets all changes to the "author_id" field.
func (m *PostMutation) ResetAuthorID() {
	m.author = nil
}

// ClearAuthor clears the "author" edge to the User entity.
func (m *PostMutation) ClearAuthor() {
	m.clearedauthor = true
	m.clearedFields[post.FieldAuthorID] = struct{}{}
}

// AuthorCleared reports if the "author" edge to the User entity was cleared.
func (m *PostMutation) AuthorCleared() bool {
	return m.clearedauthor
}

// AuthorIDs returns the "author" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// AuthorID instead. It exists only for internal usage by the builders.
func (m *PostMutation) AuthorIDs() (ids []int) {
	if id := m.author; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetAuthor resets all changes to the "author" edge.
func (m *PostMutation) ResetAuthor() {
	m.author = nil
	m.clearedauthor = false
}

// Where appends a list predicates to the PostMutation builder.
func (m *PostMutation) Where(ps ...predicate.Post) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the PostMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *PostMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Post, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *PostMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *PostMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Post).
func (m *PostMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PostMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.title != nil {
		fields = append(fields, post.FieldTitle)
	}
	if m.content != nil {
		fields = append(fields, post.FieldContent)
	}
	if m.created_at != nil {
		fields = append(fields, post.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, post.FieldUpdatedAt)
	}
	if m.author != nil {
		fields = append(fields, post.FieldAuthorID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PostMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case post.FieldTitle:
		return m.Title()
	case post.FieldContent:
		return m.Content()
	case post.FieldCreatedAt:
		return m.CreatedAt()
	case post.FieldUpdatedAt:
		return m.UpdatedAt()
	case post.FieldAuthorID:
		return m.AuthorID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An httperror is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PostMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case post.FieldTitle:
		return m.OldTitle(ctx)
	case post.FieldContent:
		return m.OldContent(ctx)
	case post.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case post.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case post.FieldAuthorID:
		return m.OldAuthorID(ctx)
	}
	return nil, fmt.Errorf("unknown Post field %s", name)
}

// SetField sets the value of a field with the given name. It returns an httperror if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PostMutation) SetField(name string, value ent.Value) error {
	switch name {
	case post.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case post.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case post.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case post.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case post.FieldAuthorID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAuthorID(v)
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PostMutation) AddedFields() []string {
	var fields []string
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PostMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an httperror if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PostMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Post numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PostMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PostMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// httperror if the field is not defined in the schema.
func (m *PostMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Post nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an httperror if the field is not defined in the schema.
func (m *PostMutation) ResetField(name string) error {
	switch name {
	case post.FieldTitle:
		m.ResetTitle()
		return nil
	case post.FieldContent:
		m.ResetContent()
		return nil
	case post.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case post.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case post.FieldAuthorID:
		m.ResetAuthorID()
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PostMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.author != nil {
		edges = append(edges, post.EdgeAuthor)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PostMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case post.EdgeAuthor:
		if id := m.author; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PostMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PostMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PostMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedauthor {
		edges = append(edges, post.EdgeAuthor)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PostMutation) EdgeCleared(name string) bool {
	switch name {
	case post.EdgeAuthor:
		return m.clearedauthor
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an httperror
// if that edge is not defined in the schema.
func (m *PostMutation) ClearEdge(name string) error {
	switch name {
	case post.EdgeAuthor:
		m.ClearAuthor()
		return nil
	}
	return fmt.Errorf("unknown Post unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an httperror if the edge is not defined in the schema.
func (m *PostMutation) ResetEdge(name string) error {
	switch name {
	case post.EdgeAuthor:
		m.ResetAuthor()
		return nil
	}
	return fmt.Errorf("unknown Post edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	email         *string
	password      *string
	name          *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	posts         *int
	clearedposts  bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an httperror otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An httperror is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An httperror is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An httperror is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An httperror is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetPostsID sets the "posts" edge to the Post entity by id.
func (m *UserMutation) SetPostsID(id int) {
	m.posts = &id
}

// ClearPosts clears the "posts" edge to the Post entity.
func (m *UserMutation) ClearPosts() {
	m.clearedposts = true
}

// PostsCleared reports if the "posts" edge to the Post entity was cleared.
func (m *UserMutation) PostsCleared() bool {
	return m.clearedposts
}

// PostsID returns the "posts" edge ID in the mutation.
func (m *UserMutation) PostsID() (id int, exists bool) {
	if m.posts != nil {
		return *m.posts, true
	}
	return
}

// PostsIDs returns the "posts" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// PostsID instead. It exists only for internal usage by the builders.
func (m *UserMutation) PostsIDs() (ids []int) {
	if id := m.posts; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetPosts resets all changes to the "posts" edge.
func (m *UserMutation) ResetPosts() {
	m.posts = nil
	m.clearedposts = false
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldEmail:
		return m.Email()
	case user.FieldPassword:
		return m.Password()
	case user.FieldName:
		return m.Name()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An httperror is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an httperror if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an httperror if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// httperror if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an httperror if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.posts != nil {
		edges = append(edges, user.EdgePosts)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgePosts:
		if id := m.posts; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedposts {
		edges = append(edges, user.EdgePosts)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgePosts:
		return m.clearedposts
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an httperror
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgePosts:
		m.ClearPosts()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an httperror if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgePosts:
		m.ResetPosts()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
