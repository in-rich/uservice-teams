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

type CreateTeamMemberHandler struct {
	teams_pb.CreateTeamMemberServer
	service services.CreateTeamMemberService
	logger  monitor.GRPCLogger
}

func (h *CreateTeamMemberHandler) createTeamMember(ctx context.Context, in *teams_pb.CreateTeamMemberRequest) (*teams_pb.TeamMember, error) {
	teamMember, err := h.service.Exec(ctx, &models.CreateTeamMemberRequest{
		TeamID: in.GetTeamId(),
		UserID: in.GetUserId(),
		Role:   in.GetRole(),
	})
	if err != nil {
		if errors.Is(err, dao.ErrMemberAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "team member already exists")
		}
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid team member data: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to create team member: %v", err)
	}

	return &teams_pb.TeamMember{
		TeamId: teamMember.TeamID,
		UserId: teamMember.UserID,
		Role:   teamMember.Role,
	}, nil
}

func (h *CreateTeamMemberHandler) CreateTeamMember(ctx context.Context, in *teams_pb.CreateTeamMemberRequest) (*teams_pb.TeamMember, error) {
	res, err := h.createTeamMember(ctx, in)
	h.logger.Report(ctx, "CreateTeamMember", err)
	return res, err
}

func NewCreateTeamMemberHandler(service services.CreateTeamMemberService, logger monitor.GRPCLogger) *CreateTeamMemberHandler {
	return &CreateTeamMemberHandler{
		service: service,
		logger:  logger,
	}
}
