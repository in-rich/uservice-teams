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

type JoinTeamWithInvitationHandler struct {
	teams_pb.JoinTeamWIthInvitationServer
	service services.JoinTeamWithInvitationService
	logger  monitor.GRPCLogger
}

func (h *JoinTeamWithInvitationHandler) joinTeamWithInvitation(ctx context.Context, in *teams_pb.JoinTeamWIthInvitationRequest) (*teams_pb.JoinTeamWIthInvitationResponse, error) {
	invitation, err := h.service.Exec(ctx, &models.JoinTeamWithInvitationRequest{
		Code:   in.GetCode(),
		UserID: in.GetUserId(),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid invitation code data: %v", err)
		}
		if errors.Is(err, dao.ErrCodeNotFound) {
			return nil, status.Errorf(codes.NotFound, "invalid invitation code data: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to join team with invitation: %v", err)
	}

	return &teams_pb.JoinTeamWIthInvitationResponse{
		TeamId: invitation.TeamID,
	}, nil
}

func (h *JoinTeamWithInvitationHandler) JoinTeamWIthInvitation(ctx context.Context, in *teams_pb.JoinTeamWIthInvitationRequest) (*teams_pb.JoinTeamWIthInvitationResponse, error) {
	res, err := h.joinTeamWithInvitation(ctx, in)
	h.logger.Report(ctx, "JoinTeamWIthInvitation", err)
	return res, err
}

func NewJoinTeamWithInvitationHandler(service services.JoinTeamWithInvitationService, logger monitor.GRPCLogger) *JoinTeamWithInvitationHandler {
	return &JoinTeamWithInvitationHandler{
		service: service,
		logger:  logger,
	}
}
