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

type GetUserRoleInTeamHandler struct {
	teams_pb.GetUserRoleInTeamServer
	service services.GetUserRoleInTeamService
}

func (h *GetUserRoleInTeamHandler) GetUserRoleInTeam(ctx context.Context, in *teams_pb.GetUserRoleInTeamRequest) (*teams_pb.GetUserRoleInTeamResponse, error) {
	role, err := h.service.Exec(ctx, &models.GetUserRoleInTeamRequest{
		TeamID: in.GetTeamId(),
		UserID: in.GetUserId(),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid selector: %v", err)
		}
		if errors.Is(err, dao.ErrTeamNotFound) {
			return nil, status.Errorf(codes.NotFound, "team not found: %v", err)
		}
		if errors.Is(err, dao.ErrMemberNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to get user role in team: %v", err)
	}

	return &teams_pb.GetUserRoleInTeamResponse{
		Role: role,
	}, nil
}

func NewGetUserRoleInTeamHandler(service services.GetUserRoleInTeamService) *GetUserRoleInTeamHandler {
	return &GetUserRoleInTeamHandler{
		service: service,
	}
}
