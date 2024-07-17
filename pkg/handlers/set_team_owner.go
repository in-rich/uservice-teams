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

type SetTeamOwnerHandler struct {
	teams_pb.SetTeamOwnerServer
	service services.SetTeamOwnerService
}

func (h *SetTeamOwnerHandler) SetTeamOwner(ctx context.Context, in *teams_pb.SetTeamOwnerRequest) (*teams_pb.Team, error) {
	team, err := h.service.Exec(ctx, &models.SetTeamOwnerRequest{
		TeamID: in.GetTeamId(),
		UserID: in.GetOwnerId(),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid team data: %v", err)
		}
		if errors.Is(err, dao.ErrTeamNotFound) {
			return nil, status.Errorf(codes.NotFound, "team not found: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to set team owner: %v", err)
	}

	return &teams_pb.Team{
		TeamId:  team.TeamID,
		Name:    team.Name,
		OwnerId: team.OwnerID,
	}, nil
}

func NewSetTeamOwnerHandler(service services.SetTeamOwnerService) *SetTeamOwnerHandler {
	return &SetTeamOwnerHandler{
		service: service,
	}
}
