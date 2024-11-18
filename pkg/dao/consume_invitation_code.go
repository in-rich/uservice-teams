package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/uptrace/bun"
	"time"
)

type ConsumeInvitationCodeData struct {
	Code string
}

type ConsumeInvitationCodeRepository interface {
	ConsumeInvitationCode(ctx context.Context, now time.Time, data *ConsumeInvitationCodeData) (*entities.InvitationCode, error)
}

type consumeInvitationCodeRepositoryImpl struct {
	db bun.IDB
}

func (r *consumeInvitationCodeRepositoryImpl) ConsumeInvitationCode(
	ctx context.Context, now time.Time, data *ConsumeInvitationCodeData,
) (*entities.InvitationCode, error) {
	invitationCode := new(entities.InvitationCode)

	if err := r.db.NewSelect().Model(invitationCode).Where("code = ?", data.Code).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrCodeNotFound
		}

		return nil, fmt.Errorf("consume invitation code: %w", err)
	}

	// Code has been consumed, delete it.
	if _, err := r.db.NewDelete().Model(invitationCode).Where("code = ?", data.Code).Exec(ctx); err != nil {
		return nil, fmt.Errorf("consume invitation code: %w", err)
	}

	// Check if code is expired. If so, return a not found error.
	if invitationCode.ExpiresAt.Before(now) {
		return nil, ErrCodeNotFound
	}

	return invitationCode, nil
}

func NewConsumeInvitationCodeRepository(db bun.IDB) ConsumeInvitationCodeRepository {
	return &consumeInvitationCodeRepositoryImpl{
		db: db,
	}
}
