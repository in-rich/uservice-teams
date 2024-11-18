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

type CreateInvitationCodeHandler struct {
	teams_pb.CreateInvitationCodeServer
	service services.CreateInvitationCodeService
	logger  monitor.GRPCLogger
}

func (h *CreateInvitationCodeHandler) createInvitationCode(ctx context.Context, in *teams_pb.CreateInvitationCodeRequest) (*teams_pb.CreateInvitationCodeResponse, error) {
	invitationCode, err := h.service.Exec(ctx, &models.CreateInvitationCodeRequest{
		TeamID: in.GetTeamId(),
	})
	if err != nil {
		if errors.Is(err, services.ErrInvalidData) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid invitation code data: %v", err)
		}
		if errors.Is(err, dao.ErrCodeAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "failed to create unique code: %v", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to create invitation code: %v", err)
	}

	return &teams_pb.CreateInvitationCodeResponse{
		Code: invitationCode.Code,
	}, nil
}

func (h *CreateInvitationCodeHandler) CreateInvitationCode(ctx context.Context, in *teams_pb.CreateInvitationCodeRequest) (*teams_pb.CreateInvitationCodeResponse, error) {
	res, err := h.createInvitationCode(ctx, in)
	h.logger.Report(ctx, "CreateInvitationCode", err)
	return res, err
}

func NewCreateInvitationCodeHandler(service services.CreateInvitationCodeService, logger monitor.GRPCLogger) *CreateInvitationCodeHandler {
	return &CreateInvitationCodeHandler{
		service: service,
		logger:  logger,
	}
}
