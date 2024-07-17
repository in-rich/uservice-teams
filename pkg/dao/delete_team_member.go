package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/uptrace/bun"
)

type DeleteTeamMemberRepository interface {
	DeleteTeamMember(ctx context.Context, team uuid.UUID, userID string) error
}

type deleteTeamMemberRepositoryImpl struct {
	db bun.IDB
}

func (r *deleteTeamMemberRepositoryImpl) DeleteTeamMember(
	ctx context.Context, team uuid.UUID, userID string,
) error {
	_, err := r.db.NewDelete().Model((*entities.TeamMember)(nil)).
		Where("team_id = ?", team).
		Where("user_id = ?", userID).
		Exec(ctx)

	return err
}

func NewDeleteTeamMemberRepository(db bun.IDB) DeleteTeamMemberRepository {
	return &deleteTeamMemberRepositoryImpl{
		db: db,
	}
}
