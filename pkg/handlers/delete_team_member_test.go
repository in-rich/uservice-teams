package handlers_test

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/handlers"
	"github.com/in-rich/uservice-teams/pkg/services"
	servicesmocks "github.com/in-rich/uservice-teams/pkg/services/mocks"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestDeleteTeamMember(t *testing.T) {
	testData := []struct {
		name string

		in *teams_pb.DeleteTeamMemberRequest

		serviceErr error

		expect     *emptypb.Empty
		expectCode codes.Code
	}{
		{
			name: "DeleteTeamMember",
			in: &teams_pb.DeleteTeamMemberRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
			},
			expect: new(emptypb.Empty),
		},
		{
			name: "InvalidArgument",
			in: &teams_pb.DeleteTeamMemberRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
			},
			serviceErr: services.ErrInvalidData,
			expectCode: codes.InvalidArgument,
		},
		{
			name: "InternalError",
			in: &teams_pb.DeleteTeamMemberRequest{
				TeamId: "team-id-1",
				UserId: "user-id-1",
			},
			serviceErr: errors.New("internal error"),
			expectCode: codes.Internal,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			service := servicesmocks.NewMockDeleteTeamMemberService(t)
			service.On("Exec", context.TODO(), tt.in.GetTeamId(), tt.in.GetUserId()).Return(tt.serviceErr)

			handler := handlers.NewDeleteTeamMemberHandler(service, monitor.NewDummyGRPCLogger())

			resp, err := handler.DeleteTeamMember(context.TODO(), tt.in)

			RequireGRPCCodesEqual(t, err, tt.expectCode)
			require.Equal(t, tt.expect, resp)

			service.AssertExpectations(t)
		})
	}
}
