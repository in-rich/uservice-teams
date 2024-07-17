package dao

import (
	"context"
	"github.com/in-rich/uservice-teams/pkg/dao/sql"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/uptrace/bun"
)

type ListUserTeamsRepository interface {
	ListUserTeams(ctx context.Context, userID string, limit int, offset int) ([]*entities.Team, error)
}

type listUserTeamsRepositoryImpl struct {
	db bun.IDB
}

func (r *listUserTeamsRepositoryImpl) ListUserTeams(
	ctx context.Context, userID string, limit int, offset int,
) ([]*entities.Team, error) {
	teams := make([]*entities.Team, 0)

	err := r.db.NewRaw(sql.SelectUserTeamsSQL, userID, limit, offset).Scan(ctx, &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func NewListUserTeamsRepository(db bun.IDB) ListUserTeamsRepository {
	return &listUserTeamsRepositoryImpl{
		db: db,
	}
}
