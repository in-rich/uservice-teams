package dao

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

type CreateTeamMemberData struct {
	Role entities.MemberRole
}

type CreateTeamMemberRepository interface {
	CreateTeamMember(ctx context.Context, team uuid.UUID, userID string, data *CreateTeamMemberData) (*entities.TeamMember, error)
}

type createTeamMemberRepositoryImpl struct {
	db bun.IDB
}

func (r *createTeamMemberRepositoryImpl) CreateTeamMember(
	ctx context.Context, team uuid.UUID, userID string, data *CreateTeamMemberData,
) (*entities.TeamMember, error) {
	teamMember := &entities.TeamMember{
		TeamID: team,
		UserID: userID,
		Role:   data.Role,
	}

	if _, err := r.db.NewInsert().Model(teamMember).Returning("*").Exec(ctx); err != nil {
		var pgErr pgdriver.Error
		if errors.As(err, &pgErr) && pgErr.IntegrityViolation() {
			return nil, ErrMemberAlreadyExists
		}

		return nil, err
	}

	return teamMember, nil
}

func NewCreateTeamMemberRepository(db bun.IDB) CreateTeamMemberRepository {
	return &createTeamMemberRepositoryImpl{
		db: db,
	}
}
