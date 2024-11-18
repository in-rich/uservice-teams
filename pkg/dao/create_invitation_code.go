package dao

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"time"
)

type CreateInvitationCodeData struct {
	Code      string
	TeamID    uuid.UUID
	ExpiresAt time.Time
}

type CreateInvitationCodeRepository interface {
	CreateInvitationCode(ctx context.Context, id uuid.UUID, data *CreateInvitationCodeData) (*entities.InvitationCode, error)
}

type createInvitationCodeRepositoryImpl struct {
	db bun.IDB
}

func (r *createInvitationCodeRepositoryImpl) CreateInvitationCode(
	ctx context.Context, id uuid.UUID, data *CreateInvitationCodeData,
) (*entities.InvitationCode, error) {
	invitationCode := &entities.InvitationCode{
		ID:        &id,
		Code:      data.Code,
		TeamID:    data.TeamID,
		ExpiresAt: data.ExpiresAt,
	}

	if _, err := r.db.NewInsert().Model(invitationCode).Returning("*").Exec(ctx); err != nil {
		var pgErr pgdriver.Error
		if errors.As(err, &pgErr) && pgErr.IntegrityViolation() {
			return nil, ErrCodeAlreadyExists
		}

		return nil, err
	}

	return invitationCode, nil
}

func NewCreateInvitationCodeRepository(db bun.IDB) CreateInvitationCodeRepository {
	return &createInvitationCodeRepositoryImpl{
		db: db,
	}
}
