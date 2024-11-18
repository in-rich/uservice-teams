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

var consumeInvitationCodeFixtures = []*entities.InvitationCode{
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
		ExpiresAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
	},
}

func TestConsumeInvitationCode(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		now  time.Time
		data *dao.ConsumeInvitationCodeData

		expect    *entities.InvitationCode
		expectErr error
	}{
		{
			name: "ConsumeInvitationCode",

			now: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			data: &dao.ConsumeInvitationCodeData{
				Code: "code-2",
			},

			expect: &entities.InvitationCode{
				ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
				Code:      "code-2",
				TeamID:    uuid.MustParse("20000000-0000-0000-0000-000000000000"),
				ExpiresAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "ConsumeInvitationCodeExpired",

			now: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			data: &dao.ConsumeInvitationCodeData{
				Code: "code-1",
			},

			expectErr: dao.ErrCodeNotFound,
		},
		{
			name: "ConsumeInvitationCodeNotFound",

			now: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			data: &dao.ConsumeInvitationCodeData{
				Code: "code-3",
			},

			expectErr: dao.ErrCodeNotFound,
		},
	}

	stx := BeginTX(db, consumeInvitationCodeFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewConsumeInvitationCodeRepository(tx)

			invitationCode, err := repo.ConsumeInvitationCode(context.TODO(), tt.now, tt.data)

			if invitationCode != nil {
				invitationCode.ExpiresAt = invitationCode.ExpiresAt.Round(time.Second)
			}

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, invitationCode)

			// Ensure consumed code has been deleted.
			ok, err := db.NewSelect().Model((*entities.InvitationCode)(nil)).Where("code = ?", tt.data.Code).Exists(context.TODO())
			require.NoError(t, err)
			require.False(t, ok)
		})
	}
}
