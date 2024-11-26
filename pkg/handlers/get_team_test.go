package handlers_test

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/handlers"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	servicesmocks "github.com/in-rich/uservice-teams/pkg/services/mocks"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestGetTeam(t *testing.T) {
	testData := []struct {
		name string

		in *teams_pb.GetTeamRequest

		serviceResponse *models.Team
		serviceErr      error

		expect     *teams_pb.Team
		expectCode codes.Code
	}{
		{
			name: "GetTeam",
			in: &teams_pb.GetTeamRequest{
				TeamId: "team-id-1",
			},
			serviceResponse: &models.Team{
				TeamID:  "team-id-1",
				OwnerID: "creator-id-1",
				Name:    "team-name-1",
			},
			expect: &teams_pb.Team{
				TeamId:  "team-id-1",
				OwnerId: "creator-id-1",
				Name:    "team-name-1",
			},
		},
		{
			name: "InvalidArgument",
			in: &teams_pb.GetTeamRequest{
				TeamId: "team-id-1",
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "InternalError",
			in: &teams_pb.GetTeamRequest{
				TeamId: "team-id-1",
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockGetTeamService(t)

			service.
				On("Exec", context.TODO(), &models.GetTeamRequest{
					TeamID: tt.in.GetTeamId(),
				}).
				Return(tt.serviceResponse, tt.serviceErr)

			handler := handlers.NewGetTeamHandler(service, monitor.NewDummyGRPCLogger())

			team, err := handler.GetTeam(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, team)

			service.AssertExpectations(t)
		})
	}
}
