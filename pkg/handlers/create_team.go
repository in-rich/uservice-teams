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

type CreateTeamHandler struct {
	teams_pb.CreateTeamServer
	service services.CreateTeamService
	logger  monitor.GRPCLogger
}

func (h *CreateTeamHandler) createTeam(ctx context.Context, in *teams_pb.CreateTeamRequest) (*teams_pb.Team, error) {
	team, err := h.service.Exec(ctx, &models.CreateTeamRequest{
		Name:      in.GetName(),
		CreatorID: in.GetCreatorId(),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid team data: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to create team: %v", err)
	}

	return &teams_pb.Team{
		TeamId:  team.TeamID,
		OwnerId: team.OwnerID,
		Name:    team.Name,
	}, nil
}

func (h *CreateTeamHandler) CreateTeam(ctx context.Context, in *teams_pb.CreateTeamRequest) (*teams_pb.Team, error) {
	res, err := h.createTeam(ctx, in)
	h.logger.Report(ctx, "CreateTeam", err)
	return res, err
}

func NewCreateTeamHandler(service services.CreateTeamService, logger monitor.GRPCLogger) *CreateTeamHandler {
	return &CreateTeamHandler{
		service: service,
		logger:  logger,
	}
}
