package services_test

import (
	"context"
	"github.com/in-rich/uservice-teams/pkg/dao"
	daomocks "github.com/in-rich/uservice-teams/pkg/dao/mocks"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateInvitationCode(t *testing.T) {
	testCases := []struct {
		name string

		in *models.CreateInvitationCodeRequest

		shouldCallCreateInvitationCodeTimes int
		invitationCodeDAOResponses          []*entities.InvitationCode
		invitationCodeDAOErrs               []error

		expect    *models.CreateInvitationCodeResponse
		expectErr error
	}{
		{
			name: "CreateInvitationCode",

			in: &models.CreateInvitationCodeRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallCreateInvitationCodeTimes: 1,
			invitationCodeDAOResponses: []*entities.InvitationCode{
				{Code: "code-1"},
			},
			invitationCodeDAOErrs: []error{
				nil,
			},

			expect: &models.CreateInvitationCodeResponse{
				Code: "code-1",
			},
		},
		{
			name: "FirstCodeAlreadyExists",

			in: &models.CreateInvitationCodeRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallCreateInvitationCodeTimes: 2,

			invitationCodeDAOResponses: []*entities.InvitationCode{
				nil,
				{Code: "code-1"},
			},
			invitationCodeDAOErrs: []error{
				dao.ErrCodeAlreadyExists,
				nil,
			},

			expect: &models.CreateInvitationCodeResponse{
				Code: "code-1",
			},
		},
		{
			name: "AllCodesAlreadyExists",

			in: &models.CreateInvitationCodeRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallCreateInvitationCodeTimes: 3,

			invitationCodeDAOResponses: []*entities.InvitationCode{
				nil,
				nil,
				nil,
			},
			invitationCodeDAOErrs: []error{
				dao.ErrCodeAlreadyExists,
				dao.ErrCodeAlreadyExists,
				dao.ErrCodeAlreadyExists,
			},

			expectErr: services.ErrGenerateCode,
		},
		{
			name: "DAOFailure",

			in: &models.CreateInvitationCodeRequest{
				TeamID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallCreateInvitationCodeTimes: 1,

			invitationCodeDAOResponses: []*entities.InvitationCode{
				nil,
			},
			invitationCodeDAOErrs: []error{
				FooErr,
			},

			expectErr: services.ErrGenerateCode,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			invitationCodeDAO := daomocks.NewMockCreateInvitationCodeRepository(t)

			for i := 0; i < tt.shouldCallCreateInvitationCodeTimes; i++ {
				invitationCodeDAO.
					On(
						"CreateInvitationCode",
						context.Background(),
						mock.Anything,
						mock.Anything,
					).
					Return(tt.invitationCodeDAOResponses[i], tt.invitationCodeDAOErrs[i]).
					Once()
			}

			service := services.NewCreateInvitationCodeService(invitationCodeDAO)

			res, err := service.Exec(context.Background(), tt.in)
			require.ErrorIs(t, err, tt.expectErr)
			require.Equal(t, tt.expect, res)

			invitationCodeDAO.AssertExpectations(t)
		})
	}
}
