package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/uptrace/bun"
)

type GetTeamRepository interface {
	GetTeam(ctx context.Context, teamID uuid.UUID) (*entities.Team, error)
}

type getTeamRepositoryImpl struct {
	db bun.IDB
}

func (r *getTeamRepositoryImpl) GetTeam(ctx context.Context, teamID uuid.UUID) (*entities.Team, error) {
	team := new(entities.Team)

	err := r.db.NewSelect().
		Model(team).
		Where("id = ?", teamID).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrTeamNotFound
		}

		return nil, err
	}

	return team, nil
}

func NewGetTeamRepository(db bun.IDB) GetTeamRepository {
	return &getTeamRepositoryImpl{
		db: db,
	}
}
