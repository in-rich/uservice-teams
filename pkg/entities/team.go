package entities

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Team struct {
	bun.BaseModel `bun:"table:teams"`

	ID *uuid.UUID `bun:"id,pk,type:uuid"`

	OwnerID string `bun:"owner_id,notnull"`
	Name    string `bun:"name,notnull"`

	CreatedAt *time.Time `bun:"created_at,notnull"`
}
