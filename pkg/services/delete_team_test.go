package services_test

import (
	"context"
	daomocks "github.com/in-rich/uservice-teams/pkg/dao/mocks"
	"github.com/in-rich/uservice-teams/pkg/services"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteTeam(t *testing.T) {
	testData := []struct {
		name string

		teamID string

		shouldCallDeleteTeamRepository bool
		deleteTeamRepositoryErr        error

		expectErr error
	}{
		{
			name:                           "DeleteTeam",
			teamID:                         "00000000-0000-0000-0000-000000000001",
			shouldCallDeleteTeamRepository: true,
		},
		{
			name:                           "DeleteTeamError",
			teamID:                         "00000000-0000-0000-0000-000000000001",
			shouldCallDeleteTeamRepository: true,
			deleteTeamRepositoryErr:        FooErr,
			expectErr:                      FooErr,
		},
		{
			name:      "InvalidData",
			teamID:    "invalid-team-id",
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			deleteTeamRepository := daomocks.NewMockDeleteTeamRepository(t)

			if tt.shouldCallDeleteTeamRepository {
				deleteTeamRepository.On("DeleteTeam", context.TODO(), mock.Anything).Return(tt.deleteTeamRepositoryErr)
			}

			service := services.NewDeleteTeamService(deleteTeamRepository)

			err := service.Exec(context.TODO(), tt.teamID)

			require.ErrorIs(t, err, tt.expectErr)

			deleteTeamRepository.AssertExpectations(t)
		})
	}
}
