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

type DeleteTeamHandler struct {
	teams_pb.DeleteTeamServer
	service services.DeleteTeamService
}

func (h *DeleteTeamHandler) DeleteTeam(ctx context.Context, in *teams_pb.DeleteTeamRequest) (*emptypb.Empty, error) {
	if err := h.service.Exec(ctx, in.GetTeamId()); err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid team data: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to delete team: %v", err)
	}

	return new(emptypb.Empty), nil
}

func NewDeleteTeamHandler(service services.DeleteTeamService) *DeleteTeamHandler {
	return &DeleteTeamHandler{
		service: service,
	}
}
