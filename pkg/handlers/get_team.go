package handlers

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetTeamHandler struct {
	teams_pb.GetTeamServer
	service services.GetTeamService
	logger  monitor.GRPCLogger
}

func (h *GetTeamHandler) getTeam(ctx context.Context, in *teams_pb.GetTeamRequest) (*teams_pb.Team, error) {
	team, err := h.service.Exec(ctx, &models.GetTeamRequest{
		TeamID: in.GetTeamId(),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid team data: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to get team: %v", err)
	}

	return &teams_pb.Team{
		TeamId:  team.TeamID,
		OwnerId: team.OwnerID,
		Name:    team.Name,
	}, nil
}

func (h *GetTeamHandler) GetTeam(ctx context.Context, in *teams_pb.GetTeamRequest) (*teams_pb.Team, error) {
	res, err := h.getTeam(ctx, in)
	h.logger.Report(ctx, "GetTeam", err)
	return res, err
}

func NewGetTeamHandler(service services.GetTeamService, logger monitor.GRPCLogger) *GetTeamHandler {
	return &GetTeamHandler{
		service: service,
		logger:  logger,
	}
}
