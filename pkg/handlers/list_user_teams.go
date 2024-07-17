package handlers

import (
	"context"
	"errors"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ListUserTeamsHandler struct {
	teams_pb.ListUserTeamsServer
	service services.ListUserTeamsService
}

func (h *ListUserTeamsHandler) ListUserTeams(ctx context.Context, in *teams_pb.ListUserTeamsRequest) (*teams_pb.ListUserTeamsResponse, error) {
	teams, err := h.service.Exec(ctx, &models.ListUserTeamsRequest{
		UserID: in.GetUserId(),
		Limit:  int(in.GetLimit()),
		Offset: int(in.GetOffset()),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid selector: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to get user teams: %v", err)
	}

	teamsResp := make([]*teams_pb.Team, 0, len(teams))
	for _, team := range teams {
		teamsResp = append(teamsResp, &teams_pb.Team{
			TeamId:  team.TeamID,
			Name:    team.Name,
			OwnerId: team.OwnerID,
		})
	}

	return &teams_pb.ListUserTeamsResponse{
		Teams: teamsResp,
	}, nil
}

func NewListUserTeamsHandler(service services.ListUserTeamsService) *ListUserTeamsHandler {
	return &ListUserTeamsHandler{
		service: service,
	}
}
