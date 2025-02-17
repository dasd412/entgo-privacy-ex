package auth

import "context"

type Role int

const (
	_Role = 1 << iota
	Admin
	Owner
	Viewer
)

type Authority interface {
	IsAdmin() bool
	IsOwner() bool
	IsViewer() bool
}

type UserAuthority struct {
	Role Role
}

func (u UserAuthority) IsAdmin() bool {
	return u.Role == Admin
}

func (u UserAuthority) IsOwner() bool {
	return u.Role == Owner
}

func (u UserAuthority) IsViewer() bool {
	return u.Role == Viewer
}

type ctxKey struct{}

func FromContext(ctx context.Context) Authority {
	v, _ := ctx.Value(ctxKey{}).(Authority)
	return v
}

func NewContext(ctx context.Context, v Authority) context.Context {
	return context.WithValue(ctx, ctxKey{}, v)
}
