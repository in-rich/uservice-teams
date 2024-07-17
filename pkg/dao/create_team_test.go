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

var createTeamFixtures = []*entities.Team{
	{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		Name:      "team-1",
		OwnerID:   "user-id-1",
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
}

func TestCreateTeam(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		data *dao.CreateTeamData

		expect    *entities.Team
		expectErr error
	}{
		{
			name: "CreateTeam",
			data: &dao.CreateTeamData{
				Name:    "team-2",
				OwnerID: "user-id-2",
			},
			expect: &entities.Team{
				Name:    "team-2",
				OwnerID: "user-id-2",
			},
		},
		{
			name: "CreateTeam/SameOwner",
			data: &dao.CreateTeamData{
				Name:    "team-2",
				OwnerID: "user-id-1",
			},
			expect: &entities.Team{
				Name:    "team-2",
				OwnerID: "user-id-1",
			},
		},
		{
			name: "CreateTeam/SameOwner/SameName",
			data: &dao.CreateTeamData{
				Name:    "team-1",
				OwnerID: "user-id-1",
			},
			expect: &entities.Team{
				Name:    "team-1",
				OwnerID: "user-id-1",
			},
		},
	}

	stx := BeginTX(db, createTeamFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewCreateTeamRepository(tx)

			team, err := repo.CreateTeam(context.TODO(), tt.data)

			if team != nil {
				// ID is randomly generated, so we ignore it.
				team.ID = nil
				// CreatedAt is randomly generated, so we ignore it.
				team.CreatedAt = nil
			}

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, team)
		})
	}
}
