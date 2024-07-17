package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/samber/lo"
	"github.com/uptrace/bun"
)

type GetUserRoleRepository interface {
	GetUserRole(ctx context.Context, teamID uuid.UUID, userID string) (*entities.Role, error)
}

type getUserRoleRepositoryImpl struct {
	db bun.IDB
}

func (r *getUserRoleRepositoryImpl) GetUserRole(ctx context.Context, teamID uuid.UUID, userID string) (*entities.Role, error) {
	team := new(entities.Team)

	err := r.db.NewSelect().Model(team).
		Column("owner_id").
		Where("id = ?", teamID).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrTeamNotFound
		}

		return nil, err
	}

	if team.OwnerID == userID {
		return lo.ToPtr(entities.RoleOwner), nil
	}

	member := new(entities.TeamMember)

	err = r.db.NewSelect().Model(member).
		Column("role").
		Where("team_id = ?", teamID).
		Where("user_id = ?", userID).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrMemberNotFound
		}

		return nil, err
	}

	switch member.Role {
	case entities.MemberRoleAdmin:
		return lo.ToPtr(entities.RoleAdmin), nil
	case entities.MemberRoleMember:
		return lo.ToPtr(entities.RoleMember), nil
	default:
		return nil, fmt.Errorf("unknown role: %v", member.Role)
	}
}

func NewGetUserRoleRepository(db bun.IDB) GetUserRoleRepository {
	return &getUserRoleRepositoryImpl{
		db: db,
	}
}
