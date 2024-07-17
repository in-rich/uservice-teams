package dao_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
)

var getUserRoleFixture = []interface{}{
	&entities.Team{
		ID:      lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		OwnerID: "user-id-1",
		Name:    "team-1",
	},

	&entities.TeamMember{
		ID:     lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		TeamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		UserID: "user-id-1",
		Role:   entities.MemberRoleAdmin,
	},
	&entities.TeamMember{
		ID:     lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
		TeamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		UserID: "user-id-2",
		Role:   entities.MemberRoleMember,
	},
	&entities.TeamMember{
		ID:     lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
		TeamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		UserID: "user-id-3",
		Role:   entities.MemberRoleAdmin,
	},
}

func TestGetUserByRole(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		teamID uuid.UUID
		userID string

		expect    *entities.Role
		expectErr error
	}{
		{
			name:   "GetOwnerRole",
			teamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			userID: "user-id-1",
			expect: lo.ToPtr(entities.RoleOwner),
		},
		{
			name:   "GetAdminRole",
			teamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			userID: "user-id-3",
			expect: lo.ToPtr(entities.RoleAdmin),
		},
		{
			name:   "GetUserRole",
			teamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			userID: "user-id-2",
			expect: lo.ToPtr(entities.RoleMember),
		},
		{
			name:      "TeamNotFound",
			teamID:    uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			userID:    "user-id-1",
			expectErr: dao.ErrTeamNotFound,
		},
		{
			name:      "MemberNotFound",
			teamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			userID:    "user-id-4",
			expectErr: dao.ErrMemberNotFound,
		},
	}

	stx := BeginTX(db, getUserRoleFixture)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewGetUserRoleRepository(tx)

			role, err := repo.GetUserRole(context.TODO(), tt.teamID, tt.userID)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, role)
		})
	}
}
