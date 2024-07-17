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

func TestSetTeamOwner(t *testing.T) {
	testData := []struct {
		name string

		in *models.SetTeamOwnerRequest

		shouldCallSetTeamOwnerRepository bool
		setTeamOwnerRepositoryResponse   *entities.Team
		setTeamOwnerRepositoryErr        error

		expect    *models.Team
		expectErr error
	}{
		{
			name: "SetTeamOwner",
			in: &models.SetTeamOwnerRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
				UserID: "user-id-1",
			},
			shouldCallSetTeamOwnerRepository: true,
			setTeamOwnerRepositoryResponse: &entities.Team{
				ID:        lo.ToPtr(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
				Name:      "team-name-1",
				OwnerID:   "user-id-1",
				CreatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expect: &models.Team{
				TeamID:  "00000000-0000-0000-0000-000000000001",
				Name:    "team-name-1",
				OwnerID: "user-id-1",
			},
		},
		{
			name: "SetTeamOwnerError",
			in: &models.SetTeamOwnerRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
				UserID: "user-id-1",
			},
			shouldCallSetTeamOwnerRepository: true,
			setTeamOwnerRepositoryErr:        FooErr,
			expectErr:                        FooErr,
		},
		{
			name: "InvalidData",
			in: &models.SetTeamOwnerRequest{
				TeamID: "invalid-team-id",
				UserID: "user-id-1",
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			setTeamOwnerRepository := daomocks.NewMockSetTeamOwnerRepository(t)

			if tt.shouldCallSetTeamOwnerRepository {
				setTeamOwnerRepository.
					On("SetTeamOwner", context.TODO(), mock.Anything, tt.in.UserID).
					Return(tt.setTeamOwnerRepositoryResponse, tt.setTeamOwnerRepositoryErr)
			}

			service := services.NewSetTeamOwnerService(setTeamOwnerRepository)

			got, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, got)

			setTeamOwnerRepository.AssertExpectations(t)
		})
	}
}
