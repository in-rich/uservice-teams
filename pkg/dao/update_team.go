package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/uptrace/bun"
)

type UpdateTeamData struct {
	Name string
}

type UpdateTeamRepository interface {
	UpdateTeam(ctx context.Context, teamID uuid.UUID, data *UpdateTeamData) (*entities.Team, error)
}

type updateTeamRepositoryImpl struct {
	db bun.IDB
}

func (r *updateTeamRepositoryImpl) UpdateTeam(
	ctx context.Context, teamID uuid.UUID, data *UpdateTeamData,
) (*entities.Team, error) {
	team := &entities.Team{
		Name: data.Name,
	}

	res, err := r.db.NewUpdate().
		Model(team).
		Column("name").
		Where("id = ?", teamID).
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
		return nil, ErrTeamNotFound
	}

	return team, nil
}

func NewUpdateTeamRepository(db bun.IDB) UpdateTeamRepository {
	return &updateTeamRepositoryImpl{
		db: db,
	}
}
