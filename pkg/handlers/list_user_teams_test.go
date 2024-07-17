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

func TestListUserTeams(t *testing.T) {
	testData := []struct {
		name string

		in *teams_pb.ListUserTeamsRequest

		serviceResp []*models.Team
		serviceErr  error

		expect     *teams_pb.ListUserTeamsResponse
		expectCode codes.Code
	}{
		{
			name: "ListUserTeams",
			in: &teams_pb.ListUserTeamsRequest{
				UserId: "user-id-1",
				Limit:  10,
			},
			serviceResp: []*models.Team{
				{
					TeamID:  "team-id-1",
					OwnerID: "user-id-1",
					Name:    "team-1",
				},
				{
					TeamID:  "team-id-2",
					OwnerID: "user-id-2",
					Name:    "team-2",
				},
			},
			expect: &teams_pb.ListUserTeamsResponse{
				Teams: []*teams_pb.Team{
					{
						TeamId:  "team-id-1",
						OwnerId: "user-id-1",
						Name:    "team-1",
					},
					{
						TeamId:  "team-id-2",
						OwnerId: "user-id-2",
						Name:    "team-2",
					},
				},
			},
		},
		{
			name: "InvalidArgument",
			in: &teams_pb.ListUserTeamsRequest{
				UserId: "user-id-1",
				Limit:  10,
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "InternalError",
			in: &teams_pb.ListUserTeamsRequest{
				UserId: "user-id-1",
				Limit:  10,
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockListUserTeamsService(t)
			service.On("Exec", context.TODO(), mock.Anything).Return(tt.serviceResp, tt.serviceErr)

			handler := handlers.NewListUserTeamsHandler(service)

			resp, err := handler.ListUserTeams(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, resp)

			service.AssertExpectations(t)
		})
	}
}
