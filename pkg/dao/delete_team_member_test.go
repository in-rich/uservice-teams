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

var deleteTeamMemberFixtures = []*entities.TeamMember{
	{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		UserID:    "user-id-1",
		Role:      entities.MemberRoleMember,
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
	{
		ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
		TeamID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		UserID:    "user-id-2",
		Role:      entities.MemberRoleAdmin,
		CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
}

func TestDeleteTeamMember(t *testing.T) {
	db := OpenDB()
	defer CloseDB(db)

	testData := []struct {
		name string

		team   uuid.UUID
		userID string

		expectErr error
	}{
		{
			name:   "DeleteTeamMember",
			team:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			userID: "user-id-1",
		},
		{
			name:   "DeleteTeamAdmin",
			team:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			userID: "user-id-2",
		},
		{
			name:   "DeleteTeamMemberNotFound",
			team:   uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			userID: "user-id-1",
		},
	}

	stx := BeginTX(db, deleteTeamMemberFixtures)
	defer RollbackTX(stx)

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tx := BeginTX[interface{}](stx, nil)
			defer RollbackTX(tx)

			repo := dao.NewDeleteTeamMemberRepository(tx)

			err := repo.DeleteTeamMember(context.TODO(), tt.team, tt.userID)

			require.ErrorIs(t, err, tt.expectErr)
		})
	}
}
