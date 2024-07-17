package entities

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type TeamMember struct {
	bun.BaseModel `bun:"table:team_members"`

	ID *uuid.UUID `bun:"id,pk,type:uuid"`

	TeamID uuid.UUID `bun:"team_id,notnull"`
	UserID string    `bun:"user_id,notnull"`

	Role MemberRole `bun:"role,notnull"`

	CreatedAt *time.Time `bun:"created_at,notnull"`
}
