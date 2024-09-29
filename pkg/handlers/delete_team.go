package handlers

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DeleteTeamHandler struct {
	teams_pb.DeleteTeamServer
	service services.DeleteTeamService
	logger  monitor.GRPCLogger
}

func (h *DeleteTeamHandler) deleteTeam(ctx context.Context, in *teams_pb.DeleteTeamRequest) (*emptypb.Empty, error) {
	if err := h.service.Exec(ctx, in.GetTeamId()); err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid team data: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to delete team: %v", err)
	}

	return new(emptypb.Empty), nil
}

func (h *DeleteTeamHandler) DeleteTeam(ctx context.Context, in *teams_pb.DeleteTeamRequest) (*emptypb.Empty, error) {
	res, err := h.deleteTeam(ctx, in)
	h.logger.Report(ctx, "DeleteTeam", err)
	return res, err
}

func NewDeleteTeamHandler(service services.DeleteTeamService, logger monitor.GRPCLogger) *DeleteTeamHandler {
	return &DeleteTeamHandler{
		service: service,
		logger:  logger,
	}
}
