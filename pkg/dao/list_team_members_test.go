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

var listTeamMembersFixture = []*entities.TeamMember{
	{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		UserID:    "user-id-1",
		Role:      entities.MemberRoleAdmin,
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
	{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
		TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		UserID:    "user-id-2",
		Role:      entities.MemberRoleMember,
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
	},
	{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
		TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		UserID:    "user-id-3",
		Role:      entities.MemberRoleAdmin,
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
	},
}

func TestListTeamMembers(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		teamID uuid.UUID
		limit  int
		offset int

		expect    []*entities.TeamMember
		expectErr error
	}{
		{
			name:   "GetAllMembers",
			teamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			limit:  10,
			offset: 0,
			expect: []*entities.TeamMember{
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
					TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					UserID:    "user-id-2",
					Role:      entities.MemberRoleMember,
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
					TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					UserID:    "user-id-3",
					Role:      entities.MemberRoleAdmin,
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
					TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					UserID:    "user-id-1",
					Role:      entities.MemberRoleAdmin,
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name:   "GetMembersWithLimit",
			teamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			limit:  1,
			offset: 0,
			expect: []*entities.TeamMember{
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
					TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					UserID:    "user-id-2",
					Role:      entities.MemberRoleMember,
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name:   "GetMembersWithOffset",
			teamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			limit:  10,
			offset: 1,
			expect: []*entities.TeamMember{
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
					TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					UserID:    "user-id-3",
					Role:      entities.MemberRoleAdmin,
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
					TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					UserID:    "user-id-1",
					Role:      entities.MemberRoleAdmin,
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
	}

	stx := BeginTX(db, listTeamMembersFixture)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewListTeamMembersRepository(tx)

			role, err := repo.ListTeamMembers(context.TODO(), tt.teamID, tt.limit, tt.offset)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, role)
		})
	}
}
