package entities

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type InvitationCode struct {
	bun.BaseModel `bun:"table:invitation_codes"`

	ID *uuid.UUID `bun:"id,pk,type:uuid"`

	Code   string    `bun:"code"`
	TeamID uuid.UUID `bun:"team_id"`

	ExpiresAt time.Time `bun:"expires_at"`
}
