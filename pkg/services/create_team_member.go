package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/models"
)

type CreateTeamMemberService interface {
	Exec(ctx context.Context, in *models.CreateTeamMemberRequest) (*models.TeamMember, error)
}

type createTeamMemberService struct {
	createTeamMemberRepository dao.CreateTeamMemberRepository
	getTeamRepository          dao.GetTeamRepository
}

func (s *createTeamMemberService) Exec(ctx context.Context, in *models.CreateTeamMemberRequest) (*models.TeamMember, error) {
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

	// Ensure team exists.
	if _, err = s.getTeamRepository.GetTeam(ctx, teamID); err != nil {
		return nil, err
	}

	teamMember, err := s.createTeamMemberRepository.CreateTeamMember(ctx, teamID, in.UserID, &dao.CreateTeamMemberData{
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

func NewCreateTeamMemberService(createTeamMemberRepository dao.CreateTeamMemberRepository, getTeamRepository dao.GetTeamRepository) CreateTeamMemberService {
	return &createTeamMemberService{
		createTeamMemberRepository: createTeamMemberRepository,
		getTeamRepository:          getTeamRepository,
	}
}
