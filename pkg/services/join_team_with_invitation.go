package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/in-rich/uservice-teams/pkg/models"
	"time"
)

type JoinTeamWithInvitationService interface {
	Exec(ctx context.Context, in *models.JoinTeamWithInvitationRequest) (*models.JoinTeamWithInvitationResponse, error)
}

type joinTeamWithInvitationServiceImpl struct {
	consumeInvitationCodeDAO dao.ConsumeInvitationCodeRepository
	createTeamMemberDAO      dao.CreateTeamMemberRepository
}

func (s *joinTeamWithInvitationServiceImpl) Exec(ctx context.Context, in *models.JoinTeamWithInvitationRequest) (*models.JoinTeamWithInvitationResponse, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(in); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	// Consume the invitation code.
	code, err := s.consumeInvitationCodeDAO.ConsumeInvitationCode(ctx, time.Now(), &dao.ConsumeInvitationCodeData{Code: in.Code})
	if err != nil {
		return nil, fmt.Errorf("consume invitation code: %w", err)
	}

	// Create the team member.
	team, err := s.createTeamMemberDAO.CreateTeamMember(
		ctx,
		code.TeamID,
		in.UserID,
		&dao.CreateTeamMemberData{Role: entities.MemberRoleMember},
	)
	if err != nil {
		return nil, fmt.Errorf("create team member: %w", err)
	}

	return &models.JoinTeamWithInvitationResponse{
		TeamID: team.TeamID.String(),
	}, nil
}

func NewJoinTeamWithInvitationService(
	consumeInvitationCodeDAO dao.ConsumeInvitationCodeRepository,
	createTeamMemberDAO dao.CreateTeamMemberRepository,
) JoinTeamWithInvitationService {
	return &joinTeamWithInvitationServiceImpl{
		consumeInvitationCodeDAO: consumeInvitationCodeDAO,
		createTeamMemberDAO:      createTeamMemberDAO,
	}
}
