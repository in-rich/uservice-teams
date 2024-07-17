package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/uptrace/bun"
)

type UpdateTeamMemberData struct {
	Role entities.MemberRole
}

type UpdateTeamMemberRepository interface {
	UpdateTeamMember(ctx context.Context, team uuid.UUID, userID string, data *UpdateTeamMemberData) (*entities.TeamMember, error)
}

type updateTeamMemberRepositoryImpl struct {
	db bun.IDB
}

func (r *updateTeamMemberRepositoryImpl) UpdateTeamMember(
	ctx context.Context, team uuid.UUID, userID string, data *UpdateTeamMemberData,
) (*entities.TeamMember, error) {
	member := &entities.TeamMember{
		Role: data.Role,
	}

	res, err := r.db.NewUpdate().Model(member).
		Column("role").
		Where("team_id = ?", team).
		Where("user_id = ?", userID).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, ErrMemberNotFound
	}

	return member, nil
}

func NewUpdateTeamMemberRepository(db bun.IDB) UpdateTeamMemberRepository {
	return &updateTeamMemberRepositoryImpl{
		db: db,
	}
}
