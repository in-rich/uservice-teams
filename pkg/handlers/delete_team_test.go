package handlers_test

import (
	"context"
	"errors"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/handlers"
	"github.com/in-rich/uservice-teams/pkg/services"
	servicesmocks "github.com/in-rich/uservice-teams/pkg/services/mocks"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestDeleteTeam(t *testing.T) {
	testData := []struct {
		name string

		in *teams_pb.DeleteTeamRequest

		serviceErr error

		expect     *emptypb.Empty
		expectCode codes.Code
	}{
		{
			name: "DeleteTeam",
			in: &teams_pb.DeleteTeamRequest{
				TeamId: "team-id-1",
			},
			expect: new(emptypb.Empty),
		},
		{
			name: "InvalidArgument",
			in: &teams_pb.DeleteTeamRequest{
				TeamId: "team-id-1",
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "InternalError",
			in: &teams_pb.DeleteTeamRequest{
				TeamId: "team-id-1",
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockDeleteTeamService(t)
			service.On("Exec", context.TODO(), tt.in.GetTeamId()).Return(tt.serviceErr)

			handler := handlers.NewDeleteTeamHandler(service)

			resp, err := handler.DeleteTeam(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, resp)

			service.AssertExpectations(t)
		})
	}
}
