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

var deleteTeamFixture = []interface{}{
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

	// External team user.
	&entities.TeamMember{
		ID:     lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000004")),
		TeamID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
		UserID: "user-id-1",
		Role:   entities.MemberRoleAdmin,
	},
}

func TestDeleteTeam(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		teamID uuid.UUID

		expectErr error
	}{
		{
			name:   "DeleteTeam",
			teamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		},
		{
			name:   "DeleteNonExistingTeam",
			teamID: uuid.MustParse("00000000-0000-0000-0000-000000000004"),
		},
	}

	stx := BeginTX(db, deleteTeamFixture)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewDeleteTeamRepository(tx)

			err := repo.DeleteTeam(context.TODO(), tt.teamID)

			require.ErrorIs(t, err, tt.expectErr)

			if err == nil {
				// Check if the team was deleted.
				members, err := dao.NewListTeamMembersRepository(tx).ListTeamMembers(context.TODO(), tt.teamID, 10, 0)

				require.NoError(t, err)
				require.Empty(t, members)
			}
		})
	}
}
