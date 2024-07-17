package dao

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/uptrace/bun"
)

type DeleteTeamRepository interface {
	DeleteTeam(ctx context.Context, team uuid.UUID) error
}

type deleteTeamRepositoryImpl struct {
	db bun.IDB
}

func (r *deleteTeamRepositoryImpl) DeleteTeam(
	ctx context.Context, team uuid.UUID,
) error {
	// Run all operations in a transaction, all or nothing.
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = r.db.NewDelete().Model((*entities.Team)(nil)).Where("id = ?", team).Exec(ctx)
	if err != nil {
		return errors.Join(err, tx.Rollback())
	}

	// Delete all members.
	_, err = r.db.NewDelete().Model((*entities.TeamMember)(nil)).Where("team_id = ?", team).Exec(ctx)
	if err != nil {
		return errors.Join(err, tx.Rollback())
	}

	return tx.Commit()
}

func NewDeleteTeamRepository(db bun.IDB) DeleteTeamRepository {
	return &deleteTeamRepositoryImpl{
		db: db,
	}
}
