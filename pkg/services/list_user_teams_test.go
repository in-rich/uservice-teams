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

func TestListUserTeams(t *testing.T) {
	testData := []struct {
		name string

		in *models.ListUserTeamsRequest

		shouldCallListUserTeams bool
		listUserTeamsResponse   []*entities.Team
		listUserTeamsErr        error

		expect    []*models.Team
		expectErr error
	}{
		{
			name: "ListUserTeams",
			in: &models.ListUserTeamsRequest{
				UserID: "user-id-1",
				Limit:  10,
				Offset: 0,
			},
			shouldCallListUserTeams: true,
			listUserTeamsResponse: []*entities.Team{
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
					OwnerID:   "user-id-1",
					Name:      "team-1",
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
					OwnerID:   "user-id-1",
					Name:      "team-2",
					CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
			},
			expect: []*models.Team{
				{
					TeamID:  "00000000-0000-0000-0000-000000000001",
					OwnerID: "user-id-1",
					Name:    "team-1",
				},
				{
					TeamID:  "00000000-0000-0000-0000-000000000002",
					OwnerID: "user-id-1",
					Name:    "team-2",
				},
			},
		},
		{
			name: "ListUserTeamsError",
			in: &models.ListUserTeamsRequest{
				UserID: "user-id-1",
				Limit:  10,
				Offset: 0,
			},
			shouldCallListUserTeams: true,
			listUserTeamsErr:        FooErr,
			expectErr:               FooErr,
		},
		{
			name: "InvalidData",
			in: &models.ListUserTeamsRequest{
				UserID: "",
				Limit:  10,
				Offset: 0,
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			listUserTeamsRepository := daomocks.NewMockListUserTeamsRepository(t)

			if tt.shouldCallListUserTeams {
				listUserTeamsRepository.
					On("ListUserTeams", context.TODO(), tt.in.UserID, tt.in.Limit, tt.in.Offset).
					Return(tt.listUserTeamsResponse, tt.listUserTeamsErr)
			}

			service := services.NewListUserTeamsService(listUserTeamsRepository)

			teams, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, teams)

			listUserTeamsRepository.AssertExpectations(t)
		})
	}
}
