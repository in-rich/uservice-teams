package handlers_test

import (
	"context"
	"errors"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/handlers"
	"github.com/in-rich/uservice-teams/pkg/services"
	servicesmocks "github.com/in-rich/uservice-teams/pkg/services/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestGetUserRoleInTeam(t *testing.T) {
	testData := []struct {
		name string

		in *teams_pb.GetUserRoleInTeamRequest

		serviceResp string
		serviceErr  error

		expect     *teams_pb.GetUserRoleInTeamResponse
		expectCode codes.Code
	}{
		{
			name: "GetUserRoleInTeam",
			in: &teams_pb.GetUserRoleInTeamRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
			},
			serviceResp: "member",
			expect: &teams_pb.GetUserRoleInTeamResponse{
				Role: "member",
			},
		},
		{
			name: "InvalidArgument",
			in: &teams_pb.GetUserRoleInTeamRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "TeamNotFound",
			in: &teams_pb.GetUserRoleInTeamRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
			},
			serviceErr: dao.ErrTeamNotFound,
			expectCode: codes.NotFound,
		},
		{
			name: "UserNotFound",
			in: &teams_pb.GetUserRoleInTeamRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
			},
			serviceErr: dao.ErrMemberNotFound,
			expectCode: codes.NotFound,
		},
		{
			name: "InternalError",
			in: &teams_pb.GetUserRoleInTeamRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockGetUserRoleInTeamService(t)
			service.On("Exec", context.TODO(), mock.Anything).Return(tt.serviceResp, tt.serviceErr)

			handler := handlers.NewGetUserRoleInTeamHandler(service)

			resp, err := handler.GetUserRoleInTeam(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, resp)

			service.AssertExpectations(t)
		})
	}
}
