package handlers

import (
	"context"
	"errors"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateTeamHandler struct {
	teams_pb.UpdateTeamServer
	service services.UpdateTeamService
}

func (h *UpdateTeamHandler) UpdateTeam(ctx context.Context, in *teams_pb.UpdateTeamRequest) (*teams_pb.Team, error) {
	team, err := h.service.Exec(ctx, &models.UpdateTeamRequest{
		TeamID: in.GetTeamId(),
		Name:   in.GetName(),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid team data: %v", err)
		}
		if errors.Is(err, dao.ErrTeamNotFound) {
			return nil, status.Errorf(codes.NotFound, "team not found: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to update team: %v", err)
	}

	return &teams_pb.Team{
		TeamId:  team.TeamID,
		Name:    team.Name,
		OwnerId: team.OwnerID,
	}, nil
}

func NewUpdateTeamHandler(service services.UpdateTeamService) *UpdateTeamHandler {
	return &UpdateTeamHandler{
		service: service,
	}
}