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

var listUserTeamsFixtures = []interface{}{
	&entities.Team{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		OwnerID:   "owner-id-1",
		Name:      "team-1",
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
	&entities.Team{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
		OwnerID:   "owner-id-1",
		Name:      "team-2",
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
	},
	&entities.Team{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
		OwnerID:   "owner-id-1",
		Name:      "team-3",
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
	},
	// Team 4 is not part of the user teams.
	&entities.Team{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000004")),
		OwnerID:   "owner-id-1",
		Name:      "team-4",
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)),
	},

	&entities.TeamMember{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		UserID:    "user-id-2",
		Role:      entities.MemberRoleAdmin,
		CreatedAt: lo.ToPtr(time.Date(2021, 2, 4, 0, 0, 0, 0, time.UTC)),
	},
	&entities.TeamMember{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
		TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000002"),
		UserID:    "user-id-2",
		Role:      entities.MemberRoleMember,
		CreatedAt: lo.ToPtr(time.Date(2021, 2, 3, 0, 0, 0, 0, time.UTC)),
	},
	&entities.TeamMember{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
		TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000003"),
		UserID:    "user-id-2",
		Role:      entities.MemberRoleMember,
		CreatedAt: lo.ToPtr(time.Date(2021, 2, 2, 0, 0, 0, 0, time.UTC)),
	},

	// Other users.
	&entities.TeamMember{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000004")),
		TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		UserID:    "user-id-3",
		Role:      entities.MemberRoleMember,
		CreatedAt: lo.ToPtr(time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)),
	},
	&entities.TeamMember{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000005")),
		TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000004"),
		UserID:    "user-id-3",
		Role:      entities.MemberRoleMember,
		CreatedAt: lo.ToPtr(time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)),
	},
}

func TestListUserTeams(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		userID string
		limit  int
		offset int

		expect    []*entities.Team
		expectErr error
	}{
		{
			name:   "ListUserTeams",
			userID: "user-id-2",
			limit:  50,
			expect: []*entities.Team{
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
					OwnerID:   "owner-id-1",
					Name:      "team-1",
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
					OwnerID:   "owner-id-1",
					Name:      "team-2",
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
					OwnerID:   "owner-id-1",
					Name:      "team-3",
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name:   "ListUserTeams/Limit",
			userID: "user-id-2",
			limit:  2,
			expect: []*entities.Team{
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
					OwnerID:   "owner-id-1",
					Name:      "team-1",
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
					OwnerID:   "owner-id-1",
					Name:      "team-2",
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name:   "ListUserTeams/Offset",
			userID: "user-id-2",
			offset: 1,
			limit:  50,
			expect: []*entities.Team{
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
					OwnerID:   "owner-id-1",
					Name:      "team-2",
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
					OwnerID:   "owner-id-1",
					Name:      "team-3",
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name:   "ListUserTeams/NoTeams",
			userID: "user-id-5",
			expect: []*entities.Team{},
		},
	}

	stx := BeginTX(db, listUserTeamsFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewListUserTeamsRepository(tx)

			teams, err := repo.ListUserTeams(context.TODO(), tt.userID, tt.limit, tt.offset)

			require.Equal(t, tt.expectErr, err)
			require.Equal(t, tt.expect, teams)
		})
	}
}
