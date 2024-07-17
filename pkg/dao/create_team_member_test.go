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

var createTeamMemberFixtures = []*entities.TeamMember{
	{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		UserID:    "user-id-1",
		Role:      entities.MemberRoleAdmin,
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
}

func TestCreateTeamMember(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		team   uuid.UUID
		userID string
		data   *dao.CreateTeamMemberData

		expect    *entities.TeamMember
		expectErr error
	}{
		{
			name:   "CreateTeamMember",
			team:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			userID: "user-id-2",
			data: &dao.CreateTeamMemberData{
				Role: entities.MemberRoleMember,
			},
			expect: &entities.TeamMember{
				TeamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				UserID: "user-id-2",
				Role:   entities.MemberRoleMember,
			},
		},
		{
			name:   "CreateTeamAdmin",
			team:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			userID: "user-id-2",
			data: &dao.CreateTeamMemberData{
				Role: entities.MemberRoleAdmin,
			},
			expect: &entities.TeamMember{
				TeamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				UserID: "user-id-2",
				Role:   entities.MemberRoleAdmin,
			},
		},
		{
			name:   "MemberAlreadyExists",
			team:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			userID: "user-id-1",
			data: &dao.CreateTeamMemberData{
				Role: entities.MemberRoleMember,
			},
			expectErr: dao.ErrMemberAlreadyExists,
		},
		{
			name:   "InNewTeam",
			team:   uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			userID: "user-id-1",
			data: &dao.CreateTeamMemberData{
				Role: entities.MemberRoleMember,
			},
			expect: &entities.TeamMember{
				TeamID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				UserID: "user-id-1",
				Role:   entities.MemberRoleMember,
			},
		},
	}

	stx := BeginTX(db, createTeamMemberFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewCreateTeamMemberRepository(tx)

			teamMember, err := repo.CreateTeamMember(context.TODO(), tt.team, tt.userID, tt.data)

			if teamMember != nil {
				// ID is randomly generated, so we ignore it.
				teamMember.ID = nil
				// CreatedAt is randomly generated, so we ignore it.
				teamMember.CreatedAt = nil
			}

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, teamMember)
		})
	}
}
