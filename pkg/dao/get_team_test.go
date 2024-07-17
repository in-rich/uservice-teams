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

var getTeamFixtures = []*entities.Team{
	{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		OwnerID:   "owner-id-1",
		Name:      "team-name-1",
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
}

func TestGetTeam(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		teamID uuid.UUID

		expect    *entities.Team
		expectErr error
	}{
		{
			name:   "GetTeam",
			teamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			expect: &entities.Team{
				ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				OwnerID:   "owner-id-1",
				Name:      "team-name-1",
				CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name:      "TeamNotFound",
			teamID:    uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			expectErr: dao.ErrTeamNotFound,
		},
	}

	stx := BeginTX(db, getTeamFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewGetTeamRepository(tx)

			team, err := repo.GetTeam(context.TODO(), tt.teamID)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, team)
		})
	}
}
