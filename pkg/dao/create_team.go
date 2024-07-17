package dao

import (
	"context"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/uptrace/bun"
)

type CreateTeamData struct {
	Name    string
	OwnerID string
}

type CreateTeamRepository interface {
	CreateTeam(ctx context.Context, data *CreateTeamData) (*entities.Team, error)
}

type createTeamRepositoryImpl struct {
	db bun.IDB
}

func (r *createTeamRepositoryImpl) CreateTeam(
	ctx context.Context, data *CreateTeamData,
) (*entities.Team, error) {
	team := &entities.Team{
		Name:    data.Name,
		OwnerID: data.OwnerID,
	}

	if _, err := r.db.NewInsert().Model(team).Returning("*").Exec(ctx); err != nil {
		return nil, err
	}

	return team, nil
}

func NewCreateTeamRepository(db bun.IDB) CreateTeamRepository {
	return &createTeamRepositoryImpl{
		db: db,
	}
}
