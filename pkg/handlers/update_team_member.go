package handlers

import (
	"context"
	"errors"
	"github.com/in-rich/lib-go/monitor"
	teams_pb "github.com/in-rich/proto/proto-go/teams"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/models"
	"github.com/in-rich/uservice-teams/pkg/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateTeamMemberHandler struct {
	teams_pb.UpdateTeamMemberServer
	service services.UpdateTeamMemberService
	logger  monitor.GRPCLogger
}

func (h *UpdateTeamMemberHandler) updateTeamMember(ctx context.Context, in *teams_pb.UpdateTeamMemberRequest) (*teams_pb.TeamMember, error) {
	teamMember, err := h.service.Exec(ctx, &models.UpdateTeamMemberRequest{
		TeamID: in.GetTeamId(),
		UserID: in.GetUserId(),
		Role:   in.GetRole(),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid team member data: %v", err)
		}
		if errors.Is(err, dao.ErrMemberNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to update team member: %v", err)
	}

	return &teams_pb.TeamMember{
		TeamId: teamMember.TeamID,
		UserId: teamMember.UserID,
		Role:   teamMember.Role,
	}, nil
}

func (h *UpdateTeamMemberHandler) UpdateTeamMember(ctx context.Context, in *teams_pb.UpdateTeamMemberRequest) (*teams_pb.TeamMember, error) {
	res, err := h.updateTeamMember(ctx, in)
	h.logger.Report(ctx, "UpdateTeamMember", err)
	return res, err
}

func NewUpdateTeamMemberHandler(service services.UpdateTeamMemberService, logger monitor.GRPCLogger) *UpdateTeamMemberHandler {
	return &UpdateTeamMemberHandler{
		service: service,
		logger:  logger,
	}
}
