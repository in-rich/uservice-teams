package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/uptrace/bun"
)

type SetTeamOwnerRepository interface {
	SetTeamOwner(ctx context.Context, teamID uuid.UUID, ownerID string) (*entities.Team, error)
}

type setTeamOwnerRepositoryImpl struct {
	db bun.IDB
}

func (r *setTeamOwnerRepositoryImpl) SetTeamOwner(
	ctx context.Context, teamID uuid.UUID, ownerID string,
) (*entities.Team, error) {
	team := &entities.Team{
		OwnerID: ownerID,
	}

	res, err := r.db.NewUpdate().
		Model(team).
		Column("owner_id").
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

func NewSetTeamOwnerRepository(db bun.IDB) SetTeamOwnerRepository {
	return &setTeamOwnerRepositoryImpl{
		db: db,
	}
}
