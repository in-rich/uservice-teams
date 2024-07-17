package services_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
	daomocks "github.com/in-rich/uservice-teams/pkg/dao/mocks"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateTeam(t *testing.T) {
	testData := []struct {
		name string

		in *models.CreateTeamRequest

		shouldCallCreateTeam bool
		createTeamResponse   *entities.Team
		createTeamErr        error

		expect    *models.Team
		expectErr error
	}{
		{
			name: "CreateTeam",
			in: &models.CreateTeamRequest{
				Name:      "team-name-1",
				CreatorID: "creator-id-1",
			},
			shouldCallCreateTeam: true,
			createTeamResponse: &entities.Team{
				ID:      lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				OwnerID: "creator-id-1",
				Name:    "team-name-1",
			},
			expect: &models.Team{
				TeamID:  "00000000-0000-0000-0000-000000000001",
				OwnerID: "creator-id-1",
				Name:    "team-name-1",
			},
		},
		{
			name: "CreateTeamDAOError",
			in: &models.CreateTeamRequest{
				Name:      "team-name-1",
				CreatorID: "creator-id-1",
			},
			shouldCallCreateTeam: true,
			createTeamErr:        FooErr,
			expectErr:            FooErr,
		},
		{
			name: "InvalidData",
			in: &models.CreateTeamRequest{
				Name:      "",
				CreatorID: "creator-id-1",
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			createTeamRepository := daomocks.NewMockCreateTeamRepository(t)

			if tt.shouldCallCreateTeam {
				createTeamRepository.
					On("CreateTeam", context.TODO(), &dao.CreateTeamData{
						Name:    tt.in.Name,
						OwnerID: tt.in.CreatorID,
					}).
					Return(tt.createTeamResponse, tt.createTeamErr)
			}

			service := services.NewCreateTeamService(createTeamRepository)

			resp, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, resp)

			createTeamRepository.AssertExpectations(t)
		})
	}
}
