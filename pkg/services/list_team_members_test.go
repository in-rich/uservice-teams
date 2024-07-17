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

func TestListTeamMembers(t *testing.T) {
	testData := []struct {
		name string

		in *models.ListTeamMembersRequest

		shouldCallListTeamMembers bool
		listTeamMembersResponse   []*entities.TeamMember
		listTeamMembersErr        error

		expect    []*models.TeamMember
		expectErr error
	}{
		{
			name: "ListTeamMembers",
			in: &models.ListTeamMembersRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
				Limit:  10,
				Offset: 0,
			},
			shouldCallListTeamMembers: true,
			listTeamMembersResponse: []*entities.TeamMember{
				{
					TeamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					UserID: "user-id-1",
					Role:   entities.MemberRoleAdmin,
				},
				{
					TeamID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					UserID: "user-id-2",
					Role:   entities.MemberRoleMember,
				},
			},
			expect: []*models.TeamMember{
				{
					TeamID: "00000000-0000-0000-0000-000000000001",
					UserID: "user-id-1",
					Role:   "admin",
				},
				{
					TeamID: "00000000-0000-0000-0000-000000000001",
					UserID: "user-id-2",
					Role:   "member",
				},
			},
		},
		{
			name: "ListTeamMembersError",
			in: &models.ListTeamMembersRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
				Limit:  10,
				Offset: 0,
			},
			shouldCallListTeamMembers: true,
			listTeamMembersErr:        FooErr,
			expectErr:                 FooErr,
		},
		{
			name: "InvalidData",
			in: &models.ListTeamMembersRequest{
				TeamID: "invalid-team-id",
				Limit:  10,
				Offset: 0,
			},
			expectErr: services.ErrInvalidData,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			listTeamMembersRepository := daomocks.NewMockListTeamMembersRepository(t)

			if tt.shouldCallListTeamMembers {
				listTeamMembersRepository.
					On("ListTeamMembers", context.TODO(), mock.Anything, tt.in.Limit, tt.in.Offset).
					Return(tt.listTeamMembersResponse, tt.listTeamMembersErr)
			}

			service := services.NewListTeamMembersService(listTeamMembersRepository)

			response, err := service.Exec(context.TODO(), tt.in)

			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, response)

			listTeamMembersRepository.AssertExpectations(t)
		})
	}
}
