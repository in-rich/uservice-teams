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

var setTeamOwnerFixtures = []*entities.Team{
	{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		OwnerID:   "owner-id-1",
		Name:      "team-1",
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
}

func TestSetTeamOwner(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		teamID  uuid.UUID
		ownerID string

		expect    *entities.Team
		expectErr error
	}{
		{
			name:    "SetTeamOwner",
			teamID:  uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			ownerID: "owner-id-2",
			expect: &entities.Team{
				ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				OwnerID:   "owner-id-2",
				Name:      "team-1",
				CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name:      "TeamNotFound",
			teamID:    uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			ownerID:   "owner-id-2",
			expectErr: dao.ErrTeamNotFound,
		},
	}

	stx := BeginTX(db, setTeamOwnerFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewSetTeamOwnerRepository(tx)

			team, err := repo.SetTeamOwner(context.TODO(), tt.teamID, tt.ownerID)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, team)
		})
	}
}
