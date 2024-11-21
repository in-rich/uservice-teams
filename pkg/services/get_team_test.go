package services_test

import (
	"context"
	"github.com/google/uuid"
	daomocks "github.com/in-rich/uservice-teams/pkg/dao/mocks"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetTeam(t *testing.T) {
	testData := []struct {
		name string

		in *models.GetTeamRequest

		shouldCallGetTeam bool
		getTeamResponse   *entities.Team
		getTeamErr        error

		expect    *models.Team
		expectErr error
	}{
		{
			name: "GetTeam",
			in: &models.GetTeamRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetTeam: true,
			getTeamResponse: &entities.Team{
				ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				OwnerID:   "user-id-1",
				Name:      "team-1",
				CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &models.Team{
				TeamID:  "00000000-0000-0000-0000-000000000001",
				OwnerID: "user-id-1",
				Name:    "team-1",
			},
		},
		{
			name: "GetTeamError",
			in: &models.GetTeamRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetTeam: true,
			getTeamErr:        FooErr,
			expectErr:         FooErr,
		},
		{
			name: "InvalidData",
			in: &models.GetTeamRequest{
				TeamID: "",
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			GetTeamRepository := daomocks.NewMockGetTeamRepository(t)

			if tt.shouldCallGetTeam {
				GetTeamRepository.
					On("GetTeam", context.TODO(), uuid.MustParse(tt.in.TeamID)).
					Return(tt.getTeamResponse, tt.getTeamErr)
			}

			service := services.NewGetTeamService(GetTeamRepository)

			teams, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, teams)

			GetTeamRepository.AssertExpectations(t)
		})
	}
}
