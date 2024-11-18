package dao_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var createInvitationCodeFixtures = []*entities.InvitationCode{
	// Expired.
	{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		Code:      "code-1",
		TeamID:    uuid.MustParse("10000000-0000-0000-0000-000000000000"),
		ExpiresAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	// Not expired.
	{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
		Code:      "code-2",
		TeamID:    uuid.MustParse("20000000-0000-0000-0000-000000000000"),
		ExpiresAt: time.Now().In(time.UTC).Add(time.Hour),
	},
}

func TestCreateInvitationCode(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		id   uuid.UUID
		data *dao.CreateInvitationCodeData

		expect    *entities.InvitationCode
		expectErr error
	}{
		{
			name: "CreateInvitationCode",

			id: uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			data: &dao.CreateInvitationCodeData{
				Code:      "code-3",
				TeamID:    uuid.MustParse("30000000-0000-0000-0000-000000000000"),
				ExpiresAt: time.Now().In(time.UTC).Add(time.Hour).Round(time.Minute),
			},

			expect: &entities.InvitationCode{
				ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
				Code:      "code-3",
				TeamID:    uuid.MustParse("30000000-0000-0000-0000-000000000000"),
				ExpiresAt: time.Now().In(time.UTC).Add(time.Hour).Round(time.Minute),
			},
		},
		{
			name: "CreateInvitationCode/SameTeam/DifferentCode",

			id: uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			data: &dao.CreateInvitationCodeData{
				Code:      "code-3",
				TeamID:    uuid.MustParse("20000000-0000-0000-0000-000000000000"),
				ExpiresAt: time.Now().In(time.UTC).Add(time.Hour).Round(time.Minute),
			},

			expect: &entities.InvitationCode{
				ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
				Code:      "code-3",
				TeamID:    uuid.MustParse("20000000-0000-0000-0000-000000000000"),
				ExpiresAt: time.Now().In(time.UTC).Add(time.Hour).Round(time.Minute),
			},
		},
		{
			name: "CreateInvitationCode/SameTeam/SameCode",

			id: uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			data: &dao.CreateInvitationCodeData{
				Code:      "code-2",
				TeamID:    uuid.MustParse("20000000-0000-0000-0000-000000000000"),
				ExpiresAt: time.Now().In(time.UTC).Add(time.Hour).Round(time.Minute),
			},

			expectErr: dao.ErrCodeAlreadyExists,
		},
		{
			name: "CreateInvitationCode/SameTeam/SameCode/OldCodeExpired",

			id: uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			data: &dao.CreateInvitationCodeData{
				Code:      "code-1",
				TeamID:    uuid.MustParse("10000000-0000-0000-0000-000000000000"),
				ExpiresAt: time.Now().In(time.UTC).Add(time.Hour).Round(time.Minute),
			},

			expect: &entities.InvitationCode{
				ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
				Code:      "code-1",
				TeamID:    uuid.MustParse("10000000-0000-0000-0000-000000000000"),
				ExpiresAt: time.Now().In(time.UTC).Add(time.Hour).Round(time.Minute),
			},
		},
	}

	stx := BeginTX(db, createInvitationCodeFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewCreateInvitationCodeRepository(tx)

			invitationCode, err := repo.CreateInvitationCode(context.TODO(), tt.id, tt.data)

			if invitationCode != nil {
				invitationCode.ExpiresAt = invitationCode.ExpiresAt.Round(time.Minute)
			}

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, invitationCode)
		})
	}
}
