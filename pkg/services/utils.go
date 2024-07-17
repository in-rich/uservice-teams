package services

import (
	"errors"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/samber/lo"
)

func StringRoleToEntityRole(role string) (*entities.MemberRole, error) {
	switch role {
	case "admin":
		return lo.ToPtr(entities.MemberRoleAdmin), nil
	case "member":
		return lo.ToPtr(entities.MemberRoleMember), nil
	default:
		return nil, errors.New("invalid role")
	}
}
