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

func TestJoinTeamWithInvitation(t *testing.T) {
	testData := []struct {
		name string

		in *models.JoinTeamWithInvitationRequest

		shouldCallConsumeInvitationCode bool
		consumeInvitationCodeResponse   *entities.InvitationCode
		consumeInvitationCodeErr        error

		shouldCallCreateTeamMember bool
		createTeamMemberResponse   *entities.TeamMember
		createTeamMemberErr        error

		expect    *models.JoinTeamWithInvitationResponse
		expectErr error
	}{
		{
			name: "JoinTeamWithInvitation",

			in: &models.JoinTeamWithInvitationRequest{
				Code:   "code-1",
				UserID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallConsumeInvitationCode: true,
			consumeInvitationCodeResponse: &entities.InvitationCode{
				TeamID: uuid.MustParse("10000000-0000-0000-0000-000000000000"),
			},

			shouldCallCreateTeamMember: true,
			createTeamMemberResponse: &entities.TeamMember{
				TeamID: uuid.MustParse("10000000-0000-0000-0000-000000000000"),
			},

			expect: &models.JoinTeamWithInvitationResponse{
				TeamID: "10000000-0000-0000-0000-000000000000",
			},
		},
		{
			name: "ConsumeInvitationCodeError",

			in: &models.JoinTeamWithInvitationRequest{
				Code:   "code-1",
				UserID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallConsumeInvitationCode: true,
			consumeInvitationCodeErr:        FooErr,

			expectErr: FooErr,
		},
		{
			name: "CreateTeamMemberError",

			in: &models.JoinTeamWithInvitationRequest{
				Code:   "code-1",
				UserID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallConsumeInvitationCode: true,
			consumeInvitationCodeResponse: &entities.InvitationCode{
				TeamID: uuid.MustParse("10000000-0000-0000-0000-000000000000"),
			},

			shouldCallCreateTeamMember: true,
			createTeamMemberErr:        FooErr,

			expectErr: FooErr,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			consumeInvitationCodeDAO := daomocks.NewMockConsumeInvitationCodeRepository(t)
			createTeamMemberDAO := daomocks.NewMockCreateTeamMemberRepository(t)

			if tt.shouldCallConsumeInvitationCode {
				consumeInvitationCodeDAO.
					On("ConsumeInvitationCode", context.Background(), mock.Anything, mock.Anything).
					Return(tt.consumeInvitationCodeResponse, tt.consumeInvitationCodeErr)
			}

			if tt.shouldCallCreateTeamMember {
				createTeamMemberDAO.
					On("CreateTeamMember", context.Background(), mock.Anything, mock.Anything, mock.Anything).
					Return(tt.createTeamMemberResponse, tt.createTeamMemberErr)
			}

			service := services.NewJoinTeamWithInvitationService(consumeInvitationCodeDAO, createTeamMemberDAO)

			res, err := service.Exec(context.Background(), tt.in)
			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, res)

			consumeInvitationCodeDAO.AssertExpectations(t)
			createTeamMemberDAO.AssertExpectations(t)
		})
	}
}
