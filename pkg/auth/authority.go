package auth

import (
	"privacy-ex/pkg/ent/user"
)

type Role int

const (
	_Role = 1 << iota
	Admin
	Author
	Guest
)

type Authority interface {
	IsAdmin() bool
	IsAuthor() bool
	IsGuest() bool
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
	} else if role == user.RoleAuthor {
		return UserAuthority{
			role: Author,
		}
	} else {
		return UserAuthority{
			role: Guest,
		}
	}
}

func (u UserAuthority) IsAdmin() bool {
	return u.role == Admin
}

func (u UserAuthority) IsAuthor() bool {
	return u.role == Author
}

func (u UserAuthority) IsGuest() bool {
	return u.role == Guest
}

func (u UserAuthority) HasRole(role Role) bool {
	return u.role&role != 0
}
