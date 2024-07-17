package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/models"
)

type UpdateTeamMemberService interface {
	Exec(ctx context.Context, in *models.UpdateTeamMemberRequest) (*models.TeamMember, error)
}

type updateTeamMemberService struct {
	updateTeamMemberRepository dao.UpdateTeamMemberRepository
}

func (s *updateTeamMemberService) Exec(ctx context.Context, in *models.UpdateTeamMemberRequest) (*models.TeamMember, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(in); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	teamID, err := uuid.Parse(in.TeamID)
	if err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	userRole, err := StringRoleToEntityRole(in.Role)
	if err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	teamMember, err := s.updateTeamMemberRepository.UpdateTeamMember(ctx, teamID, in.UserID, &dao.UpdateTeamMemberData{
		Role: *userRole,
	})
	if err != nil {
		return nil, err
	}

	return &models.TeamMember{
		TeamID: teamMember.TeamID.String(),
		UserID: teamMember.UserID,
		Role:   string(teamMember.Role),
	}, nil
}

func NewUpdateTeamMemberService(updateTeamMemberRepository dao.UpdateTeamMemberRepository) UpdateTeamMemberService {
	return &updateTeamMemberService{
		updateTeamMemberRepository: updateTeamMemberRepository,
	}
}
