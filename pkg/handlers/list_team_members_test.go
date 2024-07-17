package handlers_test

import (
	"context"
	"errors"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/handlers"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	servicesmocks "github.com/in-rich/uservice-teams/pkg/services/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestListTeamMembers(t *testing.T) {
	testData := []struct {
		name string

		in *teams_pb.ListTeamMembersRequest

		serviceResponse []*models.TeamMember
		serviceErr      error

		expect     *teams_pb.ListTeamMembersResponse
		expectCode codes.Code
	}{
		{
			name: "ListTeamMembers",
			in: &teams_pb.ListTeamMembersRequest{
				TeamId: "team-id-1",
				Limit:  10,
			},
			serviceResponse: []*models.TeamMember{
				{
					TeamID: "team-id-1",
					UserID: "user-id-1",
					Role:   "member",
				},
				{
					TeamID: "team-id-1",
					UserID: "user-id-2",
					Role:   "member",
				},
			},
			expect: &teams_pb.ListTeamMembersResponse{
				Members: []*teams_pb.TeamMember{
					{
						TeamId: "team-id-1",
						UserId: "user-id-1",
						Role:   "member",
					},
					{
						TeamId: "team-id-1",
						UserId: "user-id-2",
						Role:   "member",
					},
				},
			},
		},
		{
			name: "InvalidArgument",
			in: &teams_pb.ListTeamMembersRequest{
				TeamId: "team-id-1",
				Limit:  10,
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "InternalError",
			in: &teams_pb.ListTeamMembersRequest{
				TeamId: "team-id-1",
				Limit:  10,
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockListTeamMembersService(t)
			service.On("Exec", context.TODO(), mock.Anything).Return(tt.serviceResponse, tt.serviceErr)

			handler := handlers.NewListTeamMembersHandler(service)

			resp, err := handler.ListTeamMembers(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, resp)

			service.AssertExpectations(t)
		})
	}
}
