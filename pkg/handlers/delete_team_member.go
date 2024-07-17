package handlers

import (
	"context"
	"errors"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DeleteTeamMemberHandler struct {
	teams_pb.DeleteTeamMemberServer
	service services.DeleteTeamMemberService
}

func (h *DeleteTeamMemberHandler) DeleteTeamMember(ctx context.Context, in *teams_pb.DeleteTeamMemberRequest) (*emptypb.Empty, error) {
	if err := h.service.Exec(ctx, in.GetTeamId(), in.GetUserId()); err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid team member data: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to delete team: %v", err)
	}

	return new(emptypb.Empty), nil
}

func NewDeleteTeamMemberHandler(service services.DeleteTeamMemberService) *DeleteTeamMemberHandler {
	return &DeleteTeamMemberHandler{
		service: service,
	}
}
