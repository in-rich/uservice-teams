package services_test

import (
	"context"
	daomocks "github.com/in-rich/uservice-teams/pkg/dao/mocks"
	"github.com/in-rich/uservice-teams/pkg/services"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteTeamMember(t *testing.T) {
	testData := []struct {
		name string

		teamID string
		userID string

		shouldCallDeleteTeamMemberRepository bool
		deleteTeamMemberRepositoryErr        error

		expectErr error
	}{
		{
			name:                                 "DeleteTeamMember",
			teamID:                               "00000000-0000-0000-0000-000000000001",
			userID:                               "user-id-1",
			shouldCallDeleteTeamMemberRepository: true,
		},
		{
			name:                                 "DeleteTeamMemberError",
			teamID:                               "00000000-0000-0000-0000-000000000001",
			userID:                               "user-id-1",
			shouldCallDeleteTeamMemberRepository: true,
			deleteTeamMemberRepositoryErr:        FooErr,
			expectErr:                            FooErr,
		},
		{
			name:      "InvalidData",
			teamID:    "invalid-team-id",
			userID:    "user-id-1",
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			deleteTeamMemberRepository := daomocks.NewMockDeleteTeamMemberRepository(t)

			if tt.shouldCallDeleteTeamMemberRepository {
				deleteTeamMemberRepository.
					On("DeleteTeamMember", context.TODO(), mock.Anything, tt.userID).
					Return(tt.deleteTeamMemberRepositoryErr)
			}

			service := services.NewDeleteTeamMemberService(deleteTeamMemberRepository)

			err := service.Exec(context.TODO(), tt.teamID, tt.userID)

			require.ErrorIs(t, err, tt.expectErr)

			deleteTeamMemberRepository.AssertExpectations(t)
		})
	}
}
