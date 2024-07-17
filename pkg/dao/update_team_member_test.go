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

var updateTeamMemberFixtures = []*entities.TeamMember{
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
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
}

func TestUpdateTeamMember(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		team   uuid.UUID
		userID string
		data   *dao.UpdateTeamMemberData

		expect    *entities.TeamMember
		expectErr error
	}{
		{
			name:   "UpdateTeamAdmin",
			team:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			userID: "user-id-1",
			data: &dao.UpdateTeamMemberData{
				Role: entities.MemberRoleMember,
			},
			expect: &entities.TeamMember{
				ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				UserID:    "user-id-1",
				Role:      entities.MemberRoleMember,
				CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name:   "UpdateTeamMember",
			team:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			userID: "user-id-2",
			data: &dao.UpdateTeamMemberData{
				Role: entities.MemberRoleAdmin,
			},
			expect: &entities.TeamMember{
				ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
				TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				UserID:    "user-id-2",
				Role:      entities.MemberRoleAdmin,
				CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name:   "UpdateTeamAdmin",
			team:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			userID: "user-id-3",
			data: &dao.UpdateTeamMemberData{
				Role: entities.MemberRoleMember,
			},
			expectErr: dao.ErrMemberNotFound,
		},
	}

	stx := BeginTX(db, updateTeamMemberFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewUpdateTeamMemberRepository(tx)

			teamMember, err := repo.UpdateTeamMember(context.TODO(), tt.team, tt.userID, tt.data)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, teamMember)
		})
	}
}
