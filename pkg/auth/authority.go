package auth

import (
	"privacy-ex/pkg/ent/user"
)

type Role int

const (
	_Role = 1 << iota
	Admin
	Viewer
)

type Authority interface {
	IsAdmin() bool
	IsViewer() bool
	HasRole(role Role) bool
}

type UserAuthority struct {
	role Role
}

func NewAuthority(role user.Role) UserAuthority {
	if role == user.RoleAdmin {
		return UserAuthority{
			role: Admin,
		}
	} else {
		return UserAuthority{
			role: Viewer,
		}
	}
}

func (u UserAuthority) IsAdmin() bool {
	return u.role == Admin
}

func (u UserAuthority) IsViewer() bool {
	return u.role == Viewer
}

func (u UserAuthority) HasRole(role Role) bool {
	return u.role&role != 0
}
