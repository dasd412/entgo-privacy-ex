package auth

import (
	"privacy-ex/pkg/ent/user"
)

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
	role Role
}

func NewAuthority(role user.Role) UserAuthority {
	if role == user.RoleAdmin {
		return UserAuthority{
			role: Admin,
		}
	} else if role == user.RoleOwner {
		return UserAuthority{
			role: Owner,
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

func (u UserAuthority) IsOwner() bool {
	return u.role == Owner
}

func (u UserAuthority) IsViewer() bool {
	return u.role == Viewer
}
