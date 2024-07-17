package services_test

import (
	"context"
	daomocks "github.com/in-rich/uservice-teams/pkg/dao/mocks"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetUserRoleInTeam(t *testing.T) {
	testData := []struct {
		name string

		in *models.GetUserRoleInTeamRequest

		shouldCallGetUserRoleInTeamRepository bool
		getUserRoleInTeamRepositoryResponse   *entities.Role
		getUserRoleInTeamRepositoryErr        error

		expect    string
		expectErr error
	}{
		{
			name: "GetUserRoleInTeam",
			in: &models.GetUserRoleInTeamRequest{
				UserID: "user-id-1",
				TeamID: "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetUserRoleInTeamRepository: true,
			getUserRoleInTeamRepositoryResponse:   lo.ToPtr(entities.RoleAdmin),
			expect:                                "admin",
		},
		{
			name: "GetUserRoleInTeamError",
			in: &models.GetUserRoleInTeamRequest{
				UserID: "user-id-1",
				TeamID: "00000000-0000-0000-0000-000000000001",
			},
			shouldCallGetUserRoleInTeamRepository: true,
			getUserRoleInTeamRepositoryErr:        FooErr,
			expectErr:                             FooErr,
		},
		{
			name: "InvalidData",
			in: &models.GetUserRoleInTeamRequest{
				UserID: "user-id-1",
				TeamID: "invalid-team-id",
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			getUserRoleInTeamRepository := daomocks.NewMockGetUserRoleRepository(t)

			if tt.shouldCallGetUserRoleInTeamRepository {
				getUserRoleInTeamRepository.
					On("GetUserRole", context.TODO(), mock.Anything, tt.in.UserID).
					Return(tt.getUserRoleInTeamRepositoryResponse, tt.getUserRoleInTeamRepositoryErr)
			}

			service := services.NewGetUserRoleInTeamService(getUserRoleInTeamRepository)

			role, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, role)
		})
	}
}
