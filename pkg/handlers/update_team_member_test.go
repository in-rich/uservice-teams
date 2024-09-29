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

func TestUpdateTeamMember(t *testing.T) {
	testData := []struct {
		name string

		in *teams_pb.UpdateTeamMemberRequest

		serviceResp *models.TeamMember
		serviceErr  error

		expect     *teams_pb.TeamMember
		expectCode codes.Code
	}{
		{
			name: "UpdateTeamMember",
			in: &teams_pb.UpdateTeamMemberRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
				Role:   "member",
			},
			serviceResp: &models.TeamMember{
				TeamID: "team-id-1",
				UserID: "user-id-1",
				Role:   "member",
			},
			expect: &teams_pb.TeamMember{
				TeamId: "team-id-1",
				UserId: "user-id-1",
				Role:   "member",
			},
		},
		{
			name: "InvalidArgument",
			in: &teams_pb.UpdateTeamMemberRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
				Role:   "member",
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "MemberNotFound",
			in: &teams_pb.UpdateTeamMemberRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
				Role:   "member",
			},
			serviceErr: dao.ErrMemberNotFound,
			expectCode: codes.NotFound,
		},
		{
			name: "InternalError",
			in: &teams_pb.UpdateTeamMemberRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
				Role:   "member",
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockUpdateTeamMemberService(t)
			service.On("Exec", context.TODO(), mock.Anything).Return(tt.serviceResp, tt.serviceErr)

			handler := handlers.NewUpdateTeamMemberHandler(service, monitor.NewDummyGRPCLogger())

			resp, err := handler.UpdateTeamMember(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, resp)

			service.AssertExpectations(t)
		})
	}
}
