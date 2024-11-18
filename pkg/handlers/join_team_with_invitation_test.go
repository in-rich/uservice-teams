package handlers_test

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/handlers"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	servicesmocks "github.com/in-rich/uservice-teams/pkg/services/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestJoinTeamWithInvitation(t *testing.T) {
	testData := []struct {
		name string

		in *teams_pb.JoinTeamWIthInvitationRequest

		serviceResponse *models.JoinTeamWithInvitationResponse
		serviceErr      error

		expect     *teams_pb.JoinTeamWIthInvitationResponse
		expectCode codes.Code
	}{
		{
			name: "JoinTeamWithInvitation",

			in: &teams_pb.JoinTeamWIthInvitationRequest{
				Code:   "code-1",
				UserId: "user-id-1",
			},

			serviceResponse: &models.JoinTeamWithInvitationResponse{
				TeamID: "team-id-1",
			},

			expect: &teams_pb.JoinTeamWIthInvitationResponse{
				TeamId: "team-id-1",
			},
		},
		{
			name: "InvalidArgument",

			in: &teams_pb.JoinTeamWIthInvitationRequest{
				Code:   "code-1",
				UserId: "user-id-1",
			},

			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "NotFound",

			in: &teams_pb.JoinTeamWIthInvitationRequest{
				Code:   "code-1",
				UserId: "user-id-1",
			},

			serviceErr: dao.ErrCodeNotFound,
			expectCode: codes.NotFound,
		},
		{
			name: "InternalError",

			in: &teams_pb.JoinTeamWIthInvitationRequest{
				Code:   "code-1",
				UserId: "user-id-1",
			},

			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockJoinTeamWithInvitationService(t)

			service.On("Exec", context.Background(), mock.Anything).Return(tt.serviceResponse, tt.serviceErr)

			handler := handlers.NewJoinTeamWithInvitationHandler(service, monitor.NewDummyGRPCLogger())

			res, err := handler.JoinTeamWIthInvitation(context.Background(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, res)

			service.AssertExpectations(t)
		})
	}
}
