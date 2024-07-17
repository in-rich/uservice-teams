package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/models"
)

type GetUserRoleInTeamService interface {
	Exec(ctx context.Context, in *models.GetUserRoleInTeamRequest) (string, error)
}

type getUserRoleInTeamService struct {
	getUserRoleInTeamRepository dao.GetUserRoleRepository
}

func (s *getUserRoleInTeamService) Exec(ctx context.Context, in *models.GetUserRoleInTeamRequest) (string, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(in); err != nil {
		return "", errors.Join(ErrInvalidData, err)
	}

	teamID, err := uuid.Parse(in.TeamID)
	if err != nil {
		return "", errors.Join(ErrInvalidData, err)
	}

	role, err := s.getUserRoleInTeamRepository.GetUserRole(ctx, teamID, in.UserID)
	if err != nil {
		return "", err
	}

	return string(*role), nil
}

func NewGetUserRoleInTeamService(getUserRoleInTeamRepository dao.GetUserRoleRepository) GetUserRoleInTeamService {
	return &getUserRoleInTeamService{
		getUserRoleInTeamRepository: getUserRoleInTeamRepository,
	}
}
