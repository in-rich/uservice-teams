package handlers_test

import (
	"context"
	"errors"
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

func TestCreateTeamMember(t *testing.T) {
	testData := []struct {
		name string

		in *teams_pb.CreateTeamMemberRequest

		serviceResponse *models.TeamMember
		serviceErr      error

		expect     *teams_pb.TeamMember
		expectCode codes.Code
	}{
		{
			name: "CreateTeamMember",
			in: &teams_pb.CreateTeamMemberRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
				Role:   "member",
			},
			serviceResponse: &models.TeamMember{
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
			name: "AlreadyExists",
			in: &teams_pb.CreateTeamMemberRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
				Role:   "member",
			},
			serviceErr: dao.ErrMemberAlreadyExists,
			expectCode: codes.AlreadyExists,
		},
		{
			name: "InvalidArgument",
			in: &teams_pb.CreateTeamMemberRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
				Role:   "member",
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "InternalError",
			in: &teams_pb.CreateTeamMemberRequest{
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
			service := servicesmocks.NewMockCreateTeamMemberService(t)

			service.On("Exec", context.TODO(), mock.Anything).Return(tt.serviceResponse, tt.serviceErr)

			handler := handlers.NewCreateTeamMemberHandler(service)

			resp, err := handler.CreateTeamMember(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, resp)

			service.AssertExpectations(t)
		})
	}
}
