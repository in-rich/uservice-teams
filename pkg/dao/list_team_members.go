package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/uptrace/bun"
)

type ListTeamMembersRepository interface {
	ListTeamMembers(ctx context.Context, teamID uuid.UUID, limit int, offset int) ([]*entities.TeamMember, error)
}

type listTeamMembersRepositoryImpl struct {
	db bun.IDB
}

func (r *listTeamMembersRepositoryImpl) ListTeamMembers(ctx context.Context, teamID uuid.UUID, limit int, offset int) ([]*entities.TeamMember, error) {
	var members []*entities.TeamMember

	err := r.db.NewSelect().
		Model(&members).
		Where("team_id = ?", teamID).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Scan(ctx)

	return members, err
}

func NewListTeamMembersRepository(db bun.IDB) ListTeamMembersRepository {
	return &listTeamMembersRepositoryImpl{
		db: db,
	}
}
