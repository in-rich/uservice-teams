package services_test

import (
	"context"
	"github.com/google/uuid"
	daomocks "github.com/in-rich/uservice-teams/pkg/dao/mocks"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUpdateTeamMember(t *testing.T) {
	testData := []struct {
		name string

		in *models.UpdateTeamMemberRequest

		shouldCallUpdateTeamMemberRepository bool
		updateTeamMemberRepositoryResponse   *entities.TeamMember
		updateTeamMemberRepositoryErr        error

		expect    *models.TeamMember
		expectErr error
	}{
		{
			name: "UpdateTeamMember",
			in: &models.UpdateTeamMemberRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
				UserID: "user-id-1",
				Role:   "admin",
			},
			shouldCallUpdateTeamMemberRepository: true,
			updateTeamMemberRepositoryResponse: &entities.TeamMember{
				TeamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				UserID: "user-id-1",
				Role:   entities.MemberRoleAdmin,
			},
			expect: &models.TeamMember{
				TeamID: "00000000-0000-0000-0000-000000000001",
				UserID: "user-id-1",
				Role:   "admin",
			},
		},
		{
			name: "UpdateTeamMemberError",
			in: &models.UpdateTeamMemberRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
				UserID: "user-id-1",
				Role:   "admin",
			},
			shouldCallUpdateTeamMemberRepository: true,
			updateTeamMemberRepositoryErr:        FooErr,
			expectErr:                            FooErr,
		},
		{
			name: "InvalidData",
			in: &models.UpdateTeamMemberRequest{
				TeamID: "invalid-team-id",
				UserID: "user-id-1",
				Role:   "admin",
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			updateTeamMemberRepository := daomocks.NewMockUpdateTeamMemberRepository(t)

			if tt.shouldCallUpdateTeamMemberRepository {
				updateTeamMemberRepository.
					On("UpdateTeamMember", context.TODO(), mock.Anything, tt.in.UserID, mock.Anything).
					Return(tt.updateTeamMemberRepositoryResponse, tt.updateTeamMemberRepositoryErr)
			}

			service := services.NewUpdateTeamMemberService(updateTeamMemberRepository)

			teamMember, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, teamMember)

			updateTeamMemberRepository.AssertExpectations(t)
		})
	}
}
