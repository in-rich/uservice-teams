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

func TestUpdateTeam(t *testing.T) {
	testData := []struct {
		name string

		in *teams_pb.UpdateTeamRequest

		serviceResp *models.Team
		serviceErr  error

		expect     *teams_pb.Team
		expectCode codes.Code
	}{
		{
			name: "UpdateTeam",
			in: &teams_pb.UpdateTeamRequest{
				TeamId: "team-id-1",
				Name:   "team-1",
			},
			serviceResp: &models.Team{
				TeamID:  "team-id-1",
				OwnerID: "user-id-1",
				Name:    "team-1",
			},
			expect: &teams_pb.Team{
				TeamId:  "team-id-1",
				OwnerId: "user-id-1",
				Name:    "team-1",
			},
		},
		{
			name: "InvalidArgument",
			in: &teams_pb.UpdateTeamRequest{
				TeamId: "team-id-1",
				Name:   "team-1",
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "TeamNotFound",
			in: &teams_pb.UpdateTeamRequest{
				TeamId: "team-id-1",
				Name:   "team-1",
			},
			serviceErr: dao.ErrTeamNotFound,
			expectCode: codes.NotFound,
		},
		{
			name: "InternalError",
			in: &teams_pb.UpdateTeamRequest{
				TeamId: "team-id-1",
				Name:   "team-1",
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockUpdateTeamService(t)
			service.On("Exec", context.TODO(), mock.Anything).Return(tt.serviceResp, tt.serviceErr)

			handler := handlers.NewUpdateTeamHandler(service)

			resp, err := handler.UpdateTeam(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, resp)

			service.AssertExpectations(t)
		})
	}
}
