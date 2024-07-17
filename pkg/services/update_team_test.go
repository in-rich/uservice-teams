package services_test

import (
	"context"
	"github.com/google/uuid"
	daomocks "github.com/in-rich/uservice-teams/pkg/dao/mocks"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUpdateTeam(t *testing.T) {
	testData := []struct {
		name string

		in *models.UpdateTeamRequest

		shouldCallUpdateTeam bool
		updateTeamResponse   *entities.Team
		updateTeamErr        error

		expect    *models.Team
		expectErr error
	}{
		{
			name: "UpdateTeam",
			in: &models.UpdateTeamRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
				Name:   "new-name",
			},
			shouldCallUpdateTeam: true,
			updateTeamResponse: &entities.Team{
				ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				OwnerID:   "owner-id-1",
				Name:      "new-name",
				CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &models.Team{
				TeamID:  "00000000-0000-0000-0000-000000000001",
				OwnerID: "owner-id-1",
				Name:    "new-name",
			},
		},
		{
			name: "UpdateTeamError",
			in: &models.UpdateTeamRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
				Name:   "new-name",
			},
			shouldCallUpdateTeam: true,
			updateTeamErr:        FooErr,
			expectErr:            FooErr,
		},
		{
			name: "InvalidData",
			in: &models.UpdateTeamRequest{
				TeamID: "invalid-team-id",
				Name:   "new-name",
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			updateTeamRepository := daomocks.NewMockUpdateTeamRepository(t)

			if tt.shouldCallUpdateTeam {
				updateTeamRepository.
					On("UpdateTeam", context.TODO(), mock.Anything, mock.Anything).
					Return(tt.updateTeamResponse, tt.updateTeamErr)
			}

			service := services.NewUpdateTeamService(updateTeamRepository)

			team, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, team)

			updateTeamRepository.AssertExpectations(t)
		})
	}
}
